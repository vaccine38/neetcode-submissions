func groupAnagrams(strs []string) [][]string {

	strMap := make(map[string][]string)
	for _, str := range strs {
		chars := make([]byte, 26)
		for j := range len(str) {
			chars[str[j] - 'a'] = chars[str[j] - 'a'] + 1
		}
		var key strings.Builder
		for i, v := range chars {
			key.WriteString(fmt.Sprintf("%v%v,", i, v))
		}
		strMap[key.String()] = append(strMap[key.String()], str)
	}

	ans := make([][]string, 0)
	for _, v := range strMap {
		ans = append(ans, v)
	}
	return ans
}
