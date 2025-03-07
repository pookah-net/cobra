package cobra

import (
	"reflect"
	"testing"
)

func TestArgvRailsStyle(t *testing.T) {
	cmdLine := []string{"db:seed:load"}

	rootCmd := &Command{
		Use:       "root",
		Delimiter: RailsDelimiter,
		Run:       func(_ *Command, args []string) { cmdLine = args },
	}
	expected := []string{"db", "seed", "load"}
	received := rootCmd.argv(cmdLine)

	if !reflect.DeepEqual(expected, received) {
		t.Errorf("expected command line (%s) is not the same as the received (%s)", expected, received)
	}
}

func TestArgvStandardStyle(t *testing.T) {
	cmdLine := []string{"db", "seed", "load"}

	rootCmd := &Command{
		Use: "root",
		// Delimiter: RailsDelimiter,
		Run: func(_ *Command, args []string) { cmdLine = args },
	}
	expected := []string{"db", "seed", "load"}
	received := rootCmd.argv(cmdLine)

	if !reflect.DeepEqual(expected, received) {
		t.Errorf("expected command line (%s) is not the same as the received (%s)", expected, received)
	}
}

func TestArgvRailsPassesFlags(t *testing.T) {
	cmdLine := []string{"db:seed:load", "-f"}

	rootCmd := &Command{
		Use:       "root",
		Delimiter: RailsDelimiter,
		Run:       func(_ *Command, args []string) { cmdLine = args },
	}
	expected := []string{"db", "seed", "load", "-f"}
	received := rootCmd.argv(cmdLine)

	if !reflect.DeepEqual(expected, received) {
		t.Errorf("expected command line (%s) is not the same as the received (%s)", expected, received)
	}
}
