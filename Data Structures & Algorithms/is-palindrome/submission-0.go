func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var sb strings.Builder
	for i := range len(s) {
		if (s[i] >= '0' && s[i] <= '9') || (s[i] >= 'a' && s[i] <= 'z') {
			sb.WriteString(string(s[i]))
		}
	}
	s = sb.String()
	// fmt.Println(s)

	for i := range len(s) {
		if (len(s) % 2 == 1 && i == len(s) / 2) || (i >= len(s) / 2) {
			break
		}
		if s[i] != s[len(s) -1 -i] {
			return false
		}
	}
	return true
}
