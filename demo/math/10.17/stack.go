package stack

// isValid 20.有效的括号
func isValid(s string) bool {
	if len(s) % 2 != 0 {
		return false
	}
	ruleMap := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var stack []byte
	strArray := []byte(s)
	for _, str := range strArray {
		if right, ok := ruleMap[str]; ok {
			stack = append(stack, right)
		} else if len(stack) == 0 || str != stack[len(stack)-1] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
