func longestConsecutive(nums []int) int {
	minmap := make(map[int]int)
	maxmap := make(map[int]int)
	max := 0
	for _, v := range nums {
		if _, ok := minmap[v-1]; !ok {
			minmap[v] = v
		} else {
			minmap[v] = minmap[v-1]
		}

		if _, ok := maxmap[v+1]; !ok {
			maxmap[v] = v
		} else {
			maxmap[v] = maxmap[v+1]
			minmap[v + 1] = minmap[v]
			minmap[maxmap[v+1]] = minmap[v]
		}
		if _, ok := minmap[v-1]; ok {
			maxmap[v - 1] = maxmap[v]
			maxmap[minmap[v-1]] = maxmap[v]
		}
		
		if max < maxmap[v] - minmap[v] + 1 {
			max = maxmap[v] - minmap[v] + 1
		}
	}
	return max
}
