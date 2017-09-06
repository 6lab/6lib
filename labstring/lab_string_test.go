package labstring

import (
	"testing"
)

func TestStringBuild(t *testing.T) {
	s := StringBuild("Hello %1, welcome in '%2' today !", "World", "ExpoZ")

	if s != "Hello World, welcome in 'ExpoZ' today !" {
		t.Error("Error in StringBuild")
	}
}

func TestEscapeQuote(t *testing.T) {
	// Reference
	s := EscapeQuote("This app is named 'ExpoZ' ")

	if s != "This app is named ''ExpoZ'' " {
		t.Error("Error in StringBuild")
	}
}
