func twoSum(nums []int, target int) []int {
	nm := make(map[int]int)
	for i, v := range nums {
		nm[v] = i
	}
	for i, v := range nums {
		if j, ok:= nm[target - v]; ok && i != j {
			if i < j {
				return []int{i, j}
			}
			return []int{j, i}
		}
	}
	return nil
}
