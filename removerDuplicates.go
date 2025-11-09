package main

func main() {

}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 1 // 第一个元素始终保留，从第 2 个位置开始写入
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
