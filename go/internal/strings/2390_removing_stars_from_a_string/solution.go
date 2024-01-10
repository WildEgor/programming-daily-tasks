package removing_stars_from_a_string

/**
You are given a string s, which contains stars *.
In one operation, you can:
Choose a star in s.
Remove the closest non-star character to its left, as well as remove the star itself.
Return the string after all stars have been removed.
Note:
The input will be generated such that the operation is always possible.
It can be shown that the resulting string will always be unique.
*/

// Create stack for efficient removal of characters
func removeStarsAlt(s string) string {
	//stack := make([]byte, 0, len(s)) // bad idea, because it will allocate memory for the whole string
	var stack []byte

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}

func removeStars(s string) string {
	for i := 0; i < len(s); i++ {

		if s[i] == '*' {
			if i == 0 {
				s = s[i+1:]
				continue
			}

			// remove asterisk and the closest non-star character to its left
			s = s[:i-1] + s[i+1:]
			i = -1
		}
	}

	return s
}

var Solution = removeStarsAlt
