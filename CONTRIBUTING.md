# Contributing to StackGuardian Go SDK

Thank you for your interest in contributing to the StackGuardian Go SDK! This guide will help you understand how to make changes to the SDK.

## Table of Contents

- [Development Setup](#development-setup)
- [Project Structure](#project-structure)
- [Making Changes](#making-changes)
  - [Adding a New API Endpoint](#adding-a-new-api-endpoint)
  - [Modifying an Existing Endpoint](#modifying-an-existing-endpoint)
  - [Adding a New Field to a Type](#adding-a-new-field-to-a-type)
  - [Removing a Field](#removing-a-field)
  - [Changing a Field Type](#changing-a-field-type)
- [Testing](#testing)
- [Code Style](#code-style)
- [Submitting Changes](#submitting-changes)

## Development Setup

### Prerequisites

- Go 1.19 or later
- Git
- (Optional) golangci-lint for linting
- (Optional) goimports for formatting

### Install Development Tools

```bash
make install-dev-tools
```

This installs:
- `goimports` - for import management and formatting
- `golangci-lint` - for comprehensive linting

### Clone and Build

```bash
git clone https://github.com/StackGuardian/sg-sdk-go.git
cd sg-sdk-go
go mod download
go build ./...
```

## Project Structure

```
sg-sdk-go/
├── *.go                          # API type definitions (generated from OpenAPI)
├── stackguardian/                # Main SDK package
│   ├── client.go                 # Client implementation
│   ├── config.go                 # Configuration
│   └── doc.go                    # Package documentation
├── services/                     # Service implementations
│   ├── organizations/
│   │   └── organizations.go      # Organizations service
│   ├── stacks/
│   │   └── stacks.go             # Stacks service
│   └── ... (18 total services)
├── internal/                     # Internal utilities (not public API)
│   ├── httpclient.go             # HTTP client with retry logic
│   ├── urlutil.go                # URL building utilities
│   └── query.go                  # Query parameter encoding
├── examples/                     # Usage examples
│   ├── basic/
│   └── workflow_run/
└── tests/                        # Integration tests
```

### Key Packages

- **Root package (`api`)**: Contains all type definitions from the OpenAPI spec. These are auto-generated but can be manually updated.
- **`stackguardian/`**: Main entry point with Client, Config, and utilities.
- **`services/`**: Each service has its own package with methods for API endpoints.
- **`internal/`**: Shared utilities not exposed as public API.

## Making Changes

### Adding a New API Endpoint

When a new endpoint is added to the StackGuardian API:

1. **Update the types** (if needed):

If the endpoint uses new request/response types, add them to the appropriate file in the root package (e.g., `workflows.go`, `stacks.go`):

```go
// In workflows.go
type NewWorkflowActionRequest struct {
    Action string `json:"action" url:"-"`
    Params map[string]interface{} `json:"params,omitempty" url:"-"`
}

type NewWorkflowActionResponse struct {
    Success bool   `json:"success"`
    Message string `json:"message"`

    extraProperties map[string]interface{}
    rawJSON         json.RawMessage
}

// Add getter, UnmarshalJSON, String methods (follow existing patterns)
```

2. **Add the service method**:

In the appropriate service file (e.g., `services/workflows/workflows.go`):

```go
// PerformNewWorkflowAction performs a new action on a workflow.
func (s *Service) PerformNewWorkflowAction(
    ctx context.Context,
    org, workflow, wfGrp string,
    request *api.NewWorkflowActionRequest,
) (*api.NewWorkflowActionResponse, error) {
    // Build the URL path
    path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/workflows/%v/action", org, wfGrp, workflow)
    if err != nil {
        return nil, err
    }

    // Make the request
    var response api.NewWorkflowActionResponse
    err = s.client.Do(ctx, &internal.RequestOptions{
        Method:   http.MethodPost,
        Path:     path,
        Body:     request,
        Response: &response,
    })
    if err != nil {
        return nil, err
    }

    return &response, nil
}
```

**Key Points:**
- Use `internal.BuildURL()` for path construction (handles URL encoding and nested groups)
- Always accept `context.Context` as the first parameter
- Use pointer receivers and return pointers for response types
- Include a doc comment describing what the method does
- Handle errors properly - return them, don't log them

3. **Add tests**:

Create or update the test file (e.g., `services/workflows/workflows_test.go`):

```go
package workflows

import (
    "context"
    "testing"

    api "github.com/StackGuardian/sg-sdk-go"
    "github.com/StackGuardian/sg-sdk-go/internal"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestService_PerformNewWorkflowAction(t *testing.T) {
    // This is an example - actual tests require API credentials
    t.Skip("Integration test - requires API credentials")

    // Setup would go here if we had credentials
    // service := NewServiceWithHTTPClient(mockClient)
    // ...
}
```

### Modifying an Existing Endpoint

When an API endpoint changes:

1. **Update the service method signature** if parameters changed
2. **Update the path** if the URL changed
3. **Update request/response types** as needed
4. **Update tests** to reflect the changes
5. **Update examples** that use the endpoint

Example - adding a query parameter:

```go
// Before
func (s *Service) ListWorkflows(ctx context.Context, org, wfGrp string) (*api.WorkflowListResponse, error) {
    path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/workflows/", org, wfGrp)
    // ...
}

// After - now accepts filter parameter
func (s *Service) ListWorkflows(
    ctx context.Context,
    org, wfGrp string,
    request *api.ListWorkflowsRequest,  // New request type with query params
) (*api.WorkflowListResponse, error) {
    path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/workflows/", org, wfGrp)
    if err != nil {
        return nil, err
    }

    // Add query parameters
    queryParams, err := internal.QueryValues(request)
    if err != nil {
        return nil, err
    }

    var response api.WorkflowListResponse
    err = s.client.Do(ctx, &internal.RequestOptions{
        Method:      http.MethodGet,
        Path:        path,
        QueryParams: queryParams,  // Add this
        Response:    &response,
    })
    // ...
}
```

### Adding a New Field to a Type

When a new field is added to an API request/response:

1. **Add the field to the type definition**:

```go
type Stack struct {
    ResourceName *string `json:"ResourceName,omitempty" url:"-"`
    Description  *string `json:"Description,omitempty" url:"-"`
    Tags         *string `json:"Tags,omitempty" url:"-"`
    // NEW FIELD
    Environment  *string `json:"Environment,omitempty" url:"-"`

    extraProperties map[string]interface{}
    rawJSON         json.RawMessage
}
```

2. **Add getter method** (follow existing patterns):

```go
func (s *Stack) GetEnvironment() *string {
    if s == nil {
        return nil
    }
    return s.Environment
}
```

3. **Update documentation** if the field significantly changes behavior

**Important Notes:**
- Use **pointer types** (`*string`, `*int`, etc.) for optional fields
- Use **value types** for required fields
- JSON tags use `omitempty` for optional fields
- URL tags use `url:"-"` for fields that shouldn't be in query strings
- Always follow the existing pattern in the codebase

### Removing a Field

When a field is removed from the API:

1. **Mark as deprecated** first (if possible):

```go
type Stack struct {
    // Deprecated: OldField is deprecated and will be removed in v3.0.0
    OldField *string `json:"OldField,omitempty" url:"-"`

    NewField *string `json:"NewField,omitempty" url:"-"`
}
```

2. **In a major version**, remove the field entirely
3. **Update all examples** and tests that used the field

### Changing a Field Type

When a field type changes (e.g., `string` → `int`):

This is a **breaking change**. Handle it carefully:

1. **For major versions**: Change the type directly

```go
// Before
type RunnerConstraints struct {
    Cpu *string `json:"cpu,omitempty"`
}

// After (breaking change)
type RunnerConstraints struct {
    Cpu *int `json:"cpu,omitempty"`
}
```

2. **For minor versions**: Add a new field and deprecate the old

```go
type RunnerConstraints struct {
    // Deprecated: CpuString is deprecated, use Cpu instead
    CpuString *string `json:"cpu,omitempty"`

    Cpu *int `json:"cpuInt,omitempty"`
}
```

## Testing

### Writing Tests

1. **Unit tests** for utilities and internal packages:

```go
// internal/urlutil_test.go
func TestBuildURL(t *testing.T) {
    tests := []struct {
        name     string
        template string
        params   []interface{}
        want     string
        wantErr  bool
    }{
        {
            name:     "simple path",
            template: "/api/v1/orgs/%v/",
            params:   []interface{}{"my-org"},
            want:     "/api/v1/orgs/my-org/",
            wantErr:  false,
        },
        {
            name:     "nested workflow group",
            template: "/api/v1/orgs/%v/wfgrps/%v/",
            params:   []interface{}{"my-org", "parent/child"},
            want:     "/api/v1/orgs/my-org/wfgrps/parent/child/",
            wantErr:  false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := BuildURL("", tt.template, tt.params...)
            if tt.wantErr {
                require.Error(t, err)
            } else {
                require.NoError(t, err)
                assert.Equal(t, tt.want, got)
            }
        })
    }
}
```

2. **Integration tests** (require API credentials):

```go
// tests/integration_test.go
func TestOrganizations_ReadOrganization(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration test")
    }

    apiKey := os.Getenv("SG_API_TOKEN")
    if apiKey == "" {
        t.Skip("SG_API_TOKEN not set")
    }

    config := stackguardian.DefaultConfig()
    config.APIKey = "apikey " + apiKey

    client, err := stackguardian.NewClient(config)
    require.NoError(t, err)

    ctx := context.Background()
    org, err := client.Organizations.ReadOrganization(ctx, "test-org")

    require.NoError(t, err)
    assert.NotNil(t, org)
}
```

3. **Example tests** (included in documentation):

```go
func ExampleClient_ReadOrganization() {
    config := stackguardian.DefaultConfig()
    config.APIKey = "apikey your-token"

    client, _ := stackguardian.NewClient(config)

    ctx := context.Background()
    org, err := client.Organizations.ReadOrganization(ctx, "my-org")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Organization: %v\n", org)
}
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run only unit tests (skip integration tests)
go test -short ./...

# Run tests with race detector
go test -race ./...

# Run specific package tests
go test ./internal/...

# Run integration tests (requires SG_API_TOKEN)
export SG_API_TOKEN="your-token"
go test ./tests/...
```

### Test Guidelines

- ✅ **DO** write tests for all new functionality
- ✅ **DO** write table-driven tests for multiple scenarios
- ✅ **DO** use `t.Run()` for subtests
- ✅ **DO** use `testify` assertions for readability
- ✅ **DO** skip integration tests when credentials aren't available
- ❌ **DON'T** commit API credentials to the repository
- ❌ **DON'T** write tests that depend on external state
- ❌ **DON'T** write tests that are flaky or timing-dependent

## Code Style

### Go Standards

Follow standard Go conventions:

- Use `gofmt` and `goimports` for formatting
- Follow [Effective Go](https://go.dev/doc/effective_go)
- Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

### SDK-Specific Patterns

1. **Error Handling**:

```go
// ✅ GOOD - Return errors, don't log them
func (s *Service) DoSomething(ctx context.Context) error {
    result, err := s.client.Do(ctx, opts)
    if err != nil {
        return err  // Let caller handle logging
    }
    return nil
}

// ❌ BAD - Don't log in library code
func (s *Service) DoSomething(ctx context.Context) error {
    result, err := s.client.Do(ctx, opts)
    if err != nil {
        log.Printf("error: %v", err)  // Don't do this
        return err
    }
    return nil
}
```

2. **Context Usage**:

```go
// ✅ GOOD - Context is first parameter
func (s *Service) DoSomething(ctx context.Context, id string) error {
    // ...
}

// ❌ BAD - Context is not first
func (s *Service) DoSomething(id string, ctx context.Context) error {
    // ...
}
```

3. **Pointer vs Value**:

```go
// ✅ GOOD - Use pointers for optional fields
type Config struct {
    Required  string   // Value type for required
    Optional  *string  // Pointer for optional
}

// ✅ GOOD - Return pointers for response types
func (s *Service) Get(ctx context.Context) (*Response, error) {
    var response Response
    // ...
    return &response, nil
}
```

4. **URL Construction**:

```go
// ✅ GOOD - Use internal.BuildURL
path, err := internal.BuildURL("", "/api/v1/orgs/%v/wfgrps/%v/", org, wfGrp)

// ❌ BAD - Don't use fmt.Sprintf (doesn't handle URL encoding)
path := fmt.Sprintf("/api/v1/orgs/%s/wfgrps/%s/", org, wfGrp)
```

5. **Query Parameters**:

```go
// ✅ GOOD - Use internal.QueryValues
queryParams, err := internal.QueryValues(request)
if err != nil {
    return nil, err
}

// ❌ BAD - Don't access fields directly
queryParams := internal.EncodeQueryParams(map[string]interface{}{
    "filter": request.Filter,  // May not exist as direct field
})
```

### Running Linters

```bash
# Format code
make fmt

# Run linters
make lint

# Run all checks
make all
```

## Submitting Changes

### Before Submitting

1. **Format your code**:
```bash
make fmt
```

2. **Run tests**:
```bash
go test ./...
```

3. **Run linters**:
```bash
make lint
```

4. **Update documentation** if needed:
   - Update README.md for user-facing changes
   - Update RESTRUCTURE_NOTES.md for architectural changes
   - Update examples if API changed

### Commit Message Format

Use clear, descriptive commit messages:

```
Add support for workflow template operations

- Add CreateWorkflowTemplate method to workflows service
- Add new WorkflowTemplate type with required fields
- Add tests for template creation and validation
- Update examples to show template usage

Fixes #123
```

Format:
- First line: Brief summary (50 chars or less)
- Blank line
- Detailed description with bullet points
- Reference issues/PRs if applicable

### Pull Request Process

1. **Create a feature branch**:
```bash
git checkout -b feature/add-workflow-templates
```

2. **Make your changes** following the guidelines above

3. **Commit your changes**:
```bash
git add .
git commit -m "Add support for workflow template operations"
```

4. **Push to your fork**:
```bash
git push origin feature/add-workflow-templates
```

5. **Open a Pull Request** on GitHub with:
   - Clear description of changes
   - Link to related issues
   - Examples of usage (if applicable)
   - Test results

6. **Respond to feedback** from reviewers

### PR Review Criteria

Your PR will be reviewed for:

- ✅ Code follows Go conventions and SDK patterns
- ✅ Tests are included and passing
- ✅ Documentation is updated
- ✅ No breaking changes (or properly documented)
- ✅ Commit messages are clear
- ✅ Code is properly formatted

## Questions?

If you have questions:

- Open an issue on GitHub
- Check existing documentation
- Look at similar implementations in the codebase

Thank you for contributing! 🎉
