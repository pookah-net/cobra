package cobra

import (
	"reflect"
	"testing"
)

func TestArgv(t *testing.T) {
	cmdLine := []string{"db:seed:load"}

	rootCmd := &Command{
		Use: "root",
		Run: func(_ *Command, args []string) { cmdLine = args },
	}
	expected := []string{"db", "seed", "load"}
	received := rootCmd.argv(cmdLine)

	if !reflect.DeepEqual(expected, received) {
		t.Errorf("expected command line is not the same as the received")
	}
}
