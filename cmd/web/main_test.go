package main

import (
	"os"
	"testing"

	"github.com/rfinochi/golang-workshop-todo/pkg/common"
)

func TestMain(t *testing.T) {
	os.Setenv(common.RepositoryEnvVarName, common.RepositoryMemory)
	os.Setenv(common.PortEnvVarName, "")

	go main()
}
