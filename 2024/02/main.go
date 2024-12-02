package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	fmt.Println(solve(string(input)))
}

func solve(input string) (int, int) {
	p1, p2 := 0, 0
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		var nums [16]int
		n := 0
		for _, s := range strings.Fields(line) {
			nums[n], _ = strconv.Atoi(s)
			n++
		}

		slice := nums[:n]
		if isValid(slice) {
			p1++
			p2++
			continue
		}

		for i := 0; i < n; i++ {
			temp := make([]int, 0, n-1)
			temp = append(temp, slice[:i]...)
			temp = append(temp, slice[i+1:]...)
			if isValid(temp) {
				p2++
				break
			}
		}
	}
	return p1, p2
}

func isValid(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	d := nums[1] - nums[0]
	if abs(d) < 1 || abs(d) > 3 {
		return false
	}

	for i := 2; i < len(nums); i++ {
		curr := nums[i] - nums[i-1]
		if curr*d <= 0 || abs(curr) < 1 || abs(curr) > 3 {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
