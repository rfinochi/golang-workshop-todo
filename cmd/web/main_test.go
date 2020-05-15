package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Setenv("PORT", "82")
	os.Setenv("TODO_REPOSITORY_TYPE", "Memory")

	go main()
}
