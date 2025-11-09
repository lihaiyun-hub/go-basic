package main

func main() {

}

func PlusOne(digits []int) []int {
	n := len(digits)

	// 从最低位开始处理进位
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++ // 直接加 1，结束
			return digits
		}
		digits[i] = 0 // 9 -> 0，继续进位
	}

	// 走到这里说明全部是 9，例如 [9,9,9]
	// 需要在最前面补 1
	return append([]int{1}, digits...)
}
