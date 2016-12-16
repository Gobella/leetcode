package basic





func findKthSortedArray(nums1 []int, nums2 []int,k int) float64 {

	if len(nums1) > len(nums2) {
		return findKthSortedArray(nums2,nums1,k)
	}
	if nums1 == nil || len(nums1) == 0{
		return float64(nums2[k-1])
	}
	if k == 1 {
		if nums1[0] < nums2[0]{
			return float64(nums1[0])
		}
		return float64(nums2[0])

	}
	var pa int
	if k/2<len(nums1){
		pa = k/2
	}else {
		pa = len(nums1)
	}

	pb := k -pa
	if nums1[pa - 1] < nums2[pb-1] {
		nums1Temp := nums1[pa:]
		return findKthSortedArray(nums1Temp,nums2,k-pa)
	}else if nums1[pa - 1] > nums2[pb-1]{
		nums2Temp := nums2[pb:]
		return findKthSortedArray(nums1,nums2Temp,k-pb)
	}else{
		return float64(nums1[pa-1])
	}

}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total % 2  == 1{
		return findKthSortedArray(nums1,nums2,(total + 1)/2)
	}
	return (findKthSortedArray(nums1,nums2,total/2) + findKthSortedArray(nums1,nums2,total/2 + 1))/2
}