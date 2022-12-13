package validator_test

import (
	"strings"
	"testing"

	"github.com/go-ticketing/pkgs/validator"
	"github.com/google/go-cmp/cmp"
)

func TestNew_ReturnsACorrectlyConfiguredValidator(t *testing.T) {
	t.Parallel()

	want := &validator.Validator{
		Errors: make(map[string]string),
	}

	got := validator.New()
	if !cmp.Equal(got, want) {
		t.Errorf("Invalid validator: %s", cmp.Diff(want, got))
	}
}

func TestValid_IsTrueForEmptyErrorsMap(t *testing.T) {
	t.Parallel()

	vld := validator.New()

	want := true
	got := vld.Valid()
	if got != want {
		t.Errorf("Valid should return false as the errors map is empty, but it returned: %t", got)
	}
}

func TestValid_IsFalseForNonEmptyErrorsMap(t *testing.T) {
	t.Parallel()

	vld := validator.New()
	vld.AddError("error", "error is present")

	want := false
	got := vld.Valid()
	if got != want {
		t.Errorf("Valid should return true as the errors map is not empty, but it returned: %t", got)
	}
}

func TestAddError_CanAddAKeyWhenItIsNotAlreadyPresentInTheErrorsMap(t *testing.T) {
	t.Parallel()
	key := "test"

	vld := validator.New()
	vld.AddError(key, "value")

	_, ok := vld.Errors[key]
	if !ok {
		t.Errorf("The key \"%s\" should be present in the errors map", key)
	}
}

func TestAddError_CannotAddAKeyWhenItIsAlreadyPresentInTheErrorsMap(t *testing.T) {
	t.Parallel()
	key := "test"

	want := "value 1"
	vld := validator.New()
	vld.AddError(key, want)
	vld.AddError(key, "value 2")

	got, ok := vld.Errors[key]
	if !ok {
		t.Fatalf("The key \"%s\" should be present", key)
	}

	if got != want {
		t.Errorf("The key \"%s\" should still have the value \"%s\"", key, want)
	}
}

func TestCheck_CanAddAKeyForFalseValidationCheck(t *testing.T) {
	t.Parallel()
	key := "test"
	validationCheck := strings.Contains(key, "abc")

	vld := validator.New()
	vld.Check(validationCheck, key, "value")

	_, ok := vld.Errors[key]
	if !ok {
		t.Errorf("The key \"%s\" should be present in the errors map", key)
	}
}

func TestCheck_CannotAddAKeyForTrueValidationCheck(t *testing.T) {
	t.Parallel()
	key := "test"
	validationCheck := strings.Contains(key, "t")

	vld := validator.New()
	vld.Check(validationCheck, key, "value")

	_, ok := vld.Errors[key]
	if ok {
		t.Errorf("The key \"%s\" should not be present in the errors map", key)
	}
}
