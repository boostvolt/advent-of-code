package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	list1 := make([]int, 0, 1000)
	list2 := make([]int, 0, 1000)
	freq2 := make(map[int]int)

	for _, ln := range bytes.Fields(bytes.ReplaceAll(data, []byte{'\n'}, []byte{' '})) {
		n := 0
		for _, d := range ln {
			n = n*10 + int(d-'0')
		}
		if len(list1) > len(list2) {
			list2 = append(list2, n)
			freq2[n]++
		} else {
			list1 = append(list1, n)
		}
	}

	sort.Ints(list1)
	sort.Ints(list2)

	p1, p2 := 0, 0
	for i, v := range list1 {
		p1 += abs(list2[i] - v)
		p2 += v * freq2[v]
	}
	fmt.Println(p1, p2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
