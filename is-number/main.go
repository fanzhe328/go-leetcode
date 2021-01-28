package main

import "fmt"

func main() {
	fmt.Println(isNumber("423.3e3.2"))
}

func isNumber(s string) bool {
	// 定义有限状态机，有8个状态，分别通过eE和.进行切分
	states := map[int]map[rune]int{
		0: {' ': 0, 's': 1, '.': 4, 'd': 2}, // 初始状态，后可跟空格、正负号、小数点或数字
		1: {'.': 4, 'd': 2},                 // e 之前的正负号，后可跟小数点或数字
		2: {'d': 2, ' ': 8, '.': 3, 'e': 5}, //小数点之前的数字，后可跟数字、空格、小数点或幂符号
		3: {'d': 3, 'e': 5, ' ': 8},         // 小数点后的数字，后可跟数字、幂符号或空格
		4: {'d': 3},                         // 当小数点前为空时，小数点后当数字，后只能跟数字
		5: {'s': 6, 'd': 7},                 // 幂符号，后可跟正负号或数字
		6: {'d': 7},                         // 幂符号后的符号，后可跟数字
		7: {'d': 7, ' ': 8},                 // 幂符号后的数字，后可跟数字或空格
		8: {' ': 8},                         // 尾部的空格
	}
	step := 0
	for _, v := range s {
		// fmt.Printf("%c\n", v)
		c := charType(v)
		if next, ok := states[step][c]; ok {
			step = next
			continue
		}
		return false
	}
	if step == 2 || step == 3 || step == 7 || step == 8 {
		return true
	}
	return false
}

func charType(s rune) rune {
	if s >= '0' && s <= '9' {
		return 'd' // digit
	}
	if s == ' ' || s == '.' {
		return s
	}
	if s == 'e' || s == 'E' {
		return 'e'
	}
	if s == '-' || s == '+' {
		return 's' // sign
	}
	return s
}
