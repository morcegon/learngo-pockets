package main

import "testing"

func TestGreet_English(t *testing.T) {
	lang := language("en")
	want := "Hello world"

	got := greet(lang)

	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet_French(t *testing.T) {
	lang := language("fr")
	want := "Bonjour le monde"

	got := greet(lang)

	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet_Akkadian(t *testing.T) {
	lang := language("akk")
	want := ""

	got := greet(lang)

	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}
