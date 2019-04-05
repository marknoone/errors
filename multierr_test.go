package errors

import (
	"fmt"
	"testing"
)

func TestNoError(t *testing.T) {
	var errors MultiError
	err := errors.Error()
	if err != "" {
		t.Error("Empty errors not returning nil")

	}
}

func TestErrors(t *testing.T) {
	var errors MultiError
	errors.Append(fmt.Errorf("The First Error"))
	got := errors.Error()
	if got == "" {
		t.Fatal("Non-Empty mutlierror returning nil")
	}

	expected := "1 error(s) occurred:\n* The First Error"
	if got != expected {
		t.Fatalf("Expected %s, Got: %s", expected, got)
	}

	if len(errors) != 1 {
		t.Fatal("Length mismatch")
	}

	errors.Append(fmt.Errorf("The Second Error"))
	got = errors.Error()
	expected = "2 error(s) occurred:\n* The First Error\n* The Second Error"
	if got != expected {
		t.Fatalf("Expected %s, Got: %s", expected, got)
	}

	if len(errors) != 2 {
		t.Fatal("Length mismatch")
	}
}
