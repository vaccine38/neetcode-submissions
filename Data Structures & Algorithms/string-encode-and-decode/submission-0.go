type Solution struct{
	pos []int
}


func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		s.pos = append(s.pos, len(sb.String()))
		sb.WriteString(str)
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	ans := make([]string, 0)
	for i, v := range s.pos {
		if i == len(s.pos) - 1 {
			ans = append(ans, encoded[v:])
		} else {
			ans = append(ans, encoded[v:s.pos[i+1]])
		}
	}
	return ans
}
