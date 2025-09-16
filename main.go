package main

import (
	"fmt"
	"os"

	"github.com/kszarek/goreleaser-test/version"
)

func main() {
	// Check if version flag is requested
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Printf("goreleaser-test %s\n", version.GetFullVersion())
		return
	}

	fmt.Println("Hello, World!")
	fmt.Printf("Version: %s\n", version.GetVersion())
	fmt.Printf("Commit: %s\n", version.GetCommit())
	fmt.Printf("Built: %s\n", version.GetDate())
}
