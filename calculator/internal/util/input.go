package util

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ReadLine(r *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	line, _ := r.ReadString('\n')
	return strings.TrimSpace(line)
}

func ReadInt(r *bufio.Reader, prompt string) int {
	for {
		s := ReadLine(r, prompt)
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			fmt.Println("Please input a number.")
			continue
		}
		return n
	}
}
