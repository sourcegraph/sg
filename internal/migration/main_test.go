package migration

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	cw, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("failed to get working directory: %s", err))
	}
	testRepoRoot = cw

	os.Exit(m.Run())
}
