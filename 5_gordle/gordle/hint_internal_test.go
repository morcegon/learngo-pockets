package gordle

import "testing"

func TestFeedbackString(t *testing.T) {
	tt := map[string]struct {
		feedback feedback
		want     string
	}{
		"one hint": {
			feedback: feedback{0},
			want:     "⬜️",
		},
		"all hints": {
			feedback: feedback{0, 1, 2},
			want:     "⬜️🟡💚",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tc.feedback.String()
			if tc.want != got {
				t.Fatalf("Want %q, got %q", tc.want, got)
			}
		})
	}
}
