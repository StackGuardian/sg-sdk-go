// Command sg is the StackGuardian command-line interface.
package main

import (
	"os"

	"github.com/StackGuardian/sg-sdk-go/cli"
)

func main() {
	os.Exit(cli.Execute())
}
