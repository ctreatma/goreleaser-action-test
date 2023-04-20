package main

import (
	"fmt"

	"example.com/goreleaser-action-test/version"
)

func main() {
	fmt.Printf("provider version: %s 👍\n", version.ProviderVersion)
}
