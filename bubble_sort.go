package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the length of the input slice")
	n, err := reader.ReadString('\n')
	fmt.Println(n)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	//size, _ := strconv.Atoi(n)
	text := strings.Replace(n, "\n", "", -1)
	size, _ := strconv.Atoi(text)
	fmt.Println(size)
	in := generate_slice(size)
	fmt.Println("entering into Bubble sort")
	out := bubble_sort(in)
	fmt.Println("sorted slice is")
	fmt.Println(out)
	//m := rand.Intn(10)
	//fmt.Println(m)
}

func generate_slice(s int) []int {
	l := make([]int, s)
	//fmt.Println(s)
	k := 0
	for i := 0; i < s; i++ {
		k = rand.Intn(10)
		l[i] = k
	}
	fmt.Println("input slice")
	fmt.Println(l)
	return l
}

func bubble_sort(s []int) []int {
	//length := len(s)
	fmt.Println(s)
	//out := make([]int, length)
	for i := range s {
		for j := 0; j < i; j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s
}
