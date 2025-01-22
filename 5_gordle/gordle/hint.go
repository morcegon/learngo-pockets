package gordle

import "strings"

type (
	hint     byte
	feedback []hint
)

const (
	absentCharacter hint = iota
	wrongPosition
	correctPostion
)

func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "⬜️" // grey square
	case wrongPosition:
		return "🟡"
	case correctPostion:
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

func (fb feedback) Equal(other feedback) bool {
	if len(fb) != len(other) {
		return false
	}

	for index, value := range fb {
		if value != other[index] {
			return false
		}
	}

	return true
}
