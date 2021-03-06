package leetcode

func contains(nums []int, v int) bool {
	for _, num := range nums {
		if num == v {
			return true
		}
	}
	return false
}

func mutate(result *[][]int, cur, nums []int) {
	if len(cur) == len(nums) {
		*result = append(*result, append([]int{}, cur...))
	} else {
		for _, n := range nums {
			if contains(cur, n) {
				continue
			}
			cur = append(cur, n)
			mutate(result, cur, nums)
			cur = cur[:len(cur)-1]
		}
	}
}

func permute(nums []int) [][]int {
	res := [][]int{}
	mutate(&res, []int{}, nums)
	return res
}
