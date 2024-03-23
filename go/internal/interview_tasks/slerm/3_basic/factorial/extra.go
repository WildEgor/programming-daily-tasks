package factorial

import "fmt"

type S struct {
	a, b int
}

// String implements the fmt.Stringer interface
func (s *S) String() string {
	return fmt.Sprintf("%s", s) // Sprintf will call s.String()
}
