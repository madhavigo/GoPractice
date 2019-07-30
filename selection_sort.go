package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	in1, out := make([]int, 10)
	out := make([]int, 10)
	for k, i := range os.Args[1:] {
		j, err := strconv.Atoi(i)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		in1[k] = j
	}
	out = selection_sort(in1)
	fmt.Println("the sorted list is")
	fmt.Println(out)
}

func selection_sort(in []int) []int {
	length := len(in)
	for i := range in {
		min_idx := i
		for j := i + 1; j < length; j++ {
			if in[min_idx] > in[j] {
				min_idx = j
			}
		}
		in[i], in[min_idx] = in[min_idx], in[i]
	}
	return in
}
