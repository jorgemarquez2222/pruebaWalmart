package main

import (
	"fmt"
	"math"
)

func countRigth(s string, c string, j int) int {
	count := 0
	for i := j; i < len(s); i++ {
		if string(s[i]) == c {
			return count
		}
		count++
	}
	return 0
}

func countLeft(s string, c string, j int) int {
	count := 0
	for i := j; i > 0; i-- {
		if string(s[i]) == string(c) {
			return count
		}
		count++
	}
	return 0

}

func addValue(r, l int, output []int) []int {
	if r == 0 {
		output = append(output, l)
	}
	if l == 0 {
		output = append(output, r)
	}
	if r != 0 && l != 0 {
		output = append(output, int(math.Min(float64(r), float64(l))))
	}
	return output
}

func main() {
	s := "walmarttech"
	c := "a"
	var output = []int{}
	fmt.Println(s)
	for i, v := range s {
		if string(v) == c {
			output = append(output, 0)
			continue
		}
		r := countRigth(s, c, i)
		l := countLeft(s, c, i)
		output = addValue(r, l, output)
	}
	fmt.Println(output)
}
