package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	mul = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	do  = regexp.MustCompile(`(?s)(?:^|do\(\)).*?(?:don't\(\)|$)`)
)

func main() {
	b, _ := os.ReadFile("input.txt")
	s := string(b)
	fmt.Println(sum(s))
	fmt.Println(sum(strings.Join(do.FindAllString(s, -1), "")))
}

func sum(s string) (r int) {
	for _, m := range mul.FindAllStringSubmatch(s, -1) {
		x, _ := strconv.Atoi(m[1])
		y, _ := strconv.Atoi(m[2])
		r += x * y
	}
	return
}
