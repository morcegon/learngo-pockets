package gordle

import "strings"

type (
	hint     byte
	feedback []hint
)

const (
	absentCharacter hint = iota
	wrongPosition
	correctPostiion
)

func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "⬜️" // grey square
	case wrongPosition:
		return "🟡"
	case correctPostiion:
		return "💚"
	default:
		// this should never happen
		return "💔"
	}
}

// String implements the Stringer interace for a slice of hint
func (fb feedback) String() string {
	sb := strings.Builder{}
	for _, h := range fb {
		sb.WriteString(h.String())
	}
	return sb.String()
}
