package main

func main() {

}

func IsValid(s string) bool {
	stack := []rune{}
    pairs := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for _, r := range s {
        if r == '(' || r == '[' || r == '{' {
            stack = append(stack, r)
        } else if closer, ok := pairs[r]; ok {
            if len(stack) == 0 || stack[len(stack)-1] != closer {
                return false
            }
            stack = stack[:len(stack)-1]
        } else {
            return false // 非法字符
        }
    }
    return len(stack) == 0
}
