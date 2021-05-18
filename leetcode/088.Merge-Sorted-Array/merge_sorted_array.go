package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	m, n, l := m-1, n-1, m+n-1
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[l] = nums1[m]
			m--
		} else {
			nums1[l] = nums2[n]
			n--
		}
		l--
	}
	for n >= 0 {
		nums1[l] = nums2[n]
		n--
		l--
	}
}
