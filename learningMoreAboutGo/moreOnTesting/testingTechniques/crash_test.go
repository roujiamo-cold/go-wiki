package testingTechniques

import (
	"os"
	"os/exec"
	"testing"
)

func TestCrasher(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		Crasher()
		return
	}
	t.Log(os.Args[0])
	cmd := exec.Command(os.Args[0], "-test.run=TestCrasher")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestCmd(t *testing.T) {
	cmd := exec.Command("cat", "crash.go")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}
}
