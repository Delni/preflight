package command

import "testing"

func TestArgsValidWithOne(t *testing.T) {
	cmd := PreflightCommand()
	cmd.Flag("checklists").Value.Set("")

	err := ValidateArgs(cmd, []string{"test"})

	if err != nil {
		t.Fatalf("Should not raise error %v", err)
	}
}

func TestArgsValidWithChecklist(t *testing.T) {
	cmd := PreflightCommand()
	cmd.Flag("checklists").Value.Set("test")

	err := ValidateArgs(cmd, []string{})

	if err != nil {
		t.Fatalf("Should not raise error but got  '%v'", err)
	}
}

func TestArgsInvalidValidWithoutFileNotList(t *testing.T) {
	cmd := PreflightCommand()

	err := ValidateArgs(cmd, []string{})

	if err == nil {
		t.Fatal("Should raise error")
	}
}
