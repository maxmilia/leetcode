package validparentheses

func isValid(s string) bool {
	m := map[byte]byte{
		'[': ']',
		'(': ')',
		'{': '}',
	}
	stack := make([]byte, 0)

	if len(s)%2 != 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && m[stack[len(stack)-1]] == s[i] {
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
