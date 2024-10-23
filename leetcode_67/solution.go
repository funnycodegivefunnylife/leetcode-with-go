package leetcode_67

import "strings"

func addBinary(a string, b string) string {
	// add 0 to the front of the shorter string
	if len(a) < len(b) {
		a = addZero(a, len(b)-len(a))
	} else {
		b = addZero(b, len(a)-len(b))
	}

	carry := 0
	result := ""
	for i := len(a) - 1; i >= 0; i-- {
		sum := int(a[i]-'0') + int(b[i]-'0') + carry
		result = string(rune(sum%2+'0')) + result
		carry = sum / 2
	}

	if carry == 1 {
		result = "1" + result
	}

	return result
}

func addZero(s string, n int) string {
	return strings.Repeat("0", n) + s
}
