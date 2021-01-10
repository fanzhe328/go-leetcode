package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "We are happy."
	fmt.Println(replaceSpace(s))
}

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

// func replaceSpace(s string) string {
// 	var buffer bytes.Buffer
// 	n := len(s)
// 	for i := 0; i < n; i++ {
// 		if s[i] == ' ' {
// 			buffer.WriteString("%20")
// 		} else {
// 			buffer.WriteByte(s[i])
// 		}
// 	}
// 	return buffer.String()
// }
