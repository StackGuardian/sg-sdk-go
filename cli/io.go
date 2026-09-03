package cli

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// addBodyFlags registers --body / --body-file on commands that send a JSON request body.
func addBodyFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("body", "b", "", "request body as a JSON string")
	cmd.Flags().StringP("body-file", "f", "", "path to a JSON file holding the request body (\"-\" reads stdin)")
	cmd.MarkFlagsMutuallyExclusive("body", "body-file")
}

// readBody decodes the JSON request body from --body, --body-file or stdin into dst.
func readBody(cmd *cobra.Command, dst interface{}) error {
	inline, _ := cmd.Flags().GetString("body")
	file, _ := cmd.Flags().GetString("body-file")
	var (
		data []byte
		err  error
	)
	switch {
	case inline != "":
		data = []byte(inline)
	case file == "-":
		data, err = io.ReadAll(cmd.InOrStdin())
	case file != "":
		data, err = os.ReadFile(file)
	default:
		return errors.New("a request body is required: use --body or --body-file")
	}
	if err != nil {
		return fmt.Errorf("read request body: %w", err)
	}
	if err := json.Unmarshal(data, dst); err != nil {
		return fmt.Errorf("invalid JSON request body: %w", err)
	}
	applyExplicitNulls(data, dst)
	return checkUnknownFields(data, dst)
}

const corePkgPath = "github.com/StackGuardian/sg-sdk-go/core"

// applyExplicitNulls restores null-versus-omitted semantics for top-level
// *core.Optional[T] fields. encoding/json sets a pointer field to nil for a JSON
// null without calling the SDK's UnmarshalJSON, which would silently turn
// "clear this field" into "leave it unchanged".
func applyExplicitNulls(data []byte, dst interface{}) {
	var raw map[string]json.RawMessage
	if json.Unmarshal(data, &raw) != nil {
		return
	}
	nulls := map[string]bool{}
	for k, v := range raw {
		if string(bytes.TrimSpace(v)) == "null" {
			nulls[strings.ToLower(k)] = true
		}
	}
	if len(nulls) == 0 {
		return
	}
	v := reflect.ValueOf(dst)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return
	}
	v = v.Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !isOptionalPointer(f.Type) || !v.Field(i).CanSet() {
			continue
		}
		name := strings.Split(f.Tag.Get("json"), ",")[0]
		if name == "" {
			name = f.Name
		}
		if name == "-" || !nulls[strings.ToLower(name)] {
			continue
		}
		opt := reflect.New(f.Type.Elem())
		opt.Elem().FieldByName("Null").SetBool(true)
		v.Field(i).Set(opt)
	}
}

func isOptionalPointer(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct &&
		t.Elem().PkgPath() == corePkgPath && strings.HasPrefix(t.Elem().Name(), "Optional[")
}

// checkUnknownFields rejects top-level keys that dst's type does not declare.
// The SDK silently drops such keys on encode, so a typo or a field the SDK
// lacks would otherwise vanish from the request without any diagnostic.
func checkUnknownFields(data []byte, dst interface{}) error {
	known := knownJSONKeys(reflect.TypeOf(dst))
	if known == nil {
		return nil // untagged wrapper (union types): nothing to compare against
	}
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil // not a JSON object: the typed decode above already accepted it
	}
	var unknown []string
	for k := range raw {
		if !known[strings.ToLower(k)] {
			unknown = append(unknown, k)
		}
	}
	if len(unknown) == 0 {
		return nil
	}
	sort.Strings(unknown)
	return fmt.Errorf("unknown field(s) in request body: %s", strings.Join(unknown, ", "))
}

// knownJSONKeys returns the lower-cased JSON keys a struct type encodes (json
// matches names case-insensitively), or nil when the type declares none.
func knownJSONKeys(t reflect.Type) map[string]bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil
	}
	keys := map[string]bool{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Anonymous {
			for k := range knownJSONKeys(f.Type) {
				keys[k] = true
			}
			continue
		}
		tag, ok := f.Tag.Lookup("json")
		if !ok {
			continue
		}
		name := strings.Split(tag, ",")[0]
		switch name {
		case "-":
			continue
		case "":
			name = f.Name
		}
		keys[strings.ToLower(name)] = true
	}
	if len(keys) == 0 {
		return nil
	}
	return keys
}

// printResponse writes the captured API response body to stdout as JSON.
// Empty bodies print nothing. A non-JSON body is written as-is, which can only
// happen for endpoints whose SDK method does not decode a typed response.
func (a *app) printResponse(cmd *cobra.Command) error {
	body := a.capture.lastBody()
	if len(bytes.TrimSpace(body)) == 0 {
		return nil
	}
	var buf bytes.Buffer
	var err error
	if a.compact {
		err = json.Compact(&buf, body)
	} else {
		err = json.Indent(&buf, body, "", "  ")
	}
	if err != nil {
		buf.Reset()
		buf.Write(body)
	}
	buf.WriteByte('\n')
	_, err = cmd.OutOrStdout().Write(buf.Bytes())
	return err
}

// hasBody reports whether the caller supplied a request body flag.
func hasBody(cmd *cobra.Command) bool {
	return cmd.Flags().Changed("body") || cmd.Flags().Changed("body-file")
}

// readOptionalBody is readBody for commands where the body may be omitted,
// in which case dst is left at its zero value.
func readOptionalBody(cmd *cobra.Command, dst interface{}) error {
	if !hasBody(cmd) {
		return nil
	}
	return readBody(cmd, dst)
}
