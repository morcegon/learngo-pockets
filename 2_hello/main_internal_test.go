package main

import "testing"

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	tests := map[string]testCase{
		"English": {
			lang: "en",
			want: "Hello world",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"Akkadian": {
			lang: "akk",
			want: `unsupported language: "akk"`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)

			if got != tc.want {
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}
}
