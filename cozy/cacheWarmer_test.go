package cozy

import "testing"

func TestCheckForSubstringInSlice(t *testing.T) {
	find := []string{"/leistungen/*/**"}
	search := "https://example.com/leistungen/some/path"

	result := checkForSubstringInSlice(find, search)

	if result != true {
		t.Errorf("Expected true, but got %v", result)
	}

	search = "/not/matching/path"

	result = checkForSubstringInSlice(find, search)

	if result != false {
		t.Errorf("Expected false, but got %v", result)
	}
}
