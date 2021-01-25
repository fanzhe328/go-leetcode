package main

import "fmt"

func main() {
	fmt.Println(replaceSpace("We are happy."))
}

func replaceSpace(s string) string {
	res := make([]byte, 0, 3*len(s))
	for i, j := 0, 0; i < len(s); i++ {
		fmt.Println(i)
		if s[i] != ' ' {
			res[j] = s[i]
			j++
			continue
		}
		res[j] = '%'
		res[j+1] = '2'
		res[j+2] = '0'
		j = j + 2
	}
	return string(res)
}
