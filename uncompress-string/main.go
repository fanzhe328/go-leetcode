package main

import (
	"fmt"
)

func main() {
	s := "HG[3|B[2|CA]]F"
	fmt.Printf("%s\n", uncompress(s))
}

func uncompress(s string) string {
	type Point struct {
		Times int
		Idx   int
	}

	str := ""
	points := make([]*Point, 0, 10)
	times := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			str += string(s[i])
		} else if s[i] == '[' {
			times = 0
		} else if s[i] == '|' {
			point := &Point{times, i}
			points = append(points, point)
		} else if s[i] >= '0' && s[i] <= '9' {
			times = times*10 + int(s[i]) - '0'
		} else if s[i] == ']' {
			if len(points) > 0 {
				point := points[len(points)-1]
				point.Times--
				if point.Times == 0 {
					points = points[:len(points)-1]
					continue
				}
				i = point.Idx
			}
		}
	}
	return str
}
