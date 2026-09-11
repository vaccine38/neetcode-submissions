func hasDuplicate(nums []int) bool {
    nm := make(map[int]bool)
    for _, n := range nums {
        if _, ok := nm[n]; ok {
            return true
        }
        nm[n] = true
    }
    return false
}
