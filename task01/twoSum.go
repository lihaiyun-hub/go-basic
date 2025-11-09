package main

func main() {

}
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int) // value -> index

	for i, num := range nums {
		complement := target - num
		if j, ok := seen[complement]; ok {
			return []int{j, i} // 找到答案
		}
		seen[num] = i // 记录当前数
	}

	// 题目保证有解，这里不会执行
	return nil
}
