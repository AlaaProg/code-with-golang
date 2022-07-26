package thealogrithms

func TestSearchOfTarget(nums []int, target int) int {

	return FindBytBinarySearch(nums, target, false)
}

func TestSearchOfTargetFromRight(nums []int, target int) int {

	return FindBytBinarySearch(nums, target, true)
}

func TestSearchOfRangeTarget(nums []int, target int) []int {

	return []int{
		FindBytBinarySearch(nums, target, false),
		FindBytBinarySearch(nums, target, true),
	}

}

func RunTest() {

	// fmt.Println(TestSearchOfTargetFromRight([]int{10, 12, 16, 59, 100, 66, 11}, 2))
	// fmt.Println(TestSearchOfRangeTarget([]int{5, 6, 7, 7, 7, 8, 8, 8, 8, 10}, 8))
}
