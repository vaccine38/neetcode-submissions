func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if s == t {
		return true
	}
	sm := make(map[byte]int)
	tm := make(map[byte]int)
	for i := range len(s) {
		sm[s[i]] = sm[s[i]] + 1
		tm[t[i]] = tm[t[i]] + 1
	}

	if len(sm) != len(tm) {
		return false
	}
	for k, v := range sm {
		if v != tm[k] {
			return false
		}
	}
	return true
}
