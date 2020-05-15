package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Setenv("TODO_REPOSITORY_TYPE", "Memory")

	go main()
}
