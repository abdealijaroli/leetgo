package contains_duplicate

func containsDuplicate(nums []int) bool {
    set := map[int]struct{}{}

	for _, n := range nums {
		if _, ok := set[n]; ok {
			return true
		}
		set[n] = struct{}{}
	}
    
	return false  
}