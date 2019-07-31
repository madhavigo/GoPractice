package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Please enter the triangle length")
	reader := bufio.NewReader(os.Stdin)
	num, err := reader.ReadString('\n')
	//fmt.Println(num, s, err)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	num = strings.Replace(num, "\n", "", -1)
	n, error := strconv.Atoi(num)
	if error != nil {
		log.Fatal(error)
		os.Exit(1)
	}
	floyd_triangle(n)
}

func floyd_triangle(n int) {
	k := 1
	//fmt.Println(int(num))
	for i := 0; i < n; i++ {
		fmt.Println()
		for j := 0; j <= i; j++ {
			if j == 0 {
				fmt.Printf("%d", k)
			} else {
				fmt.Printf("%4d", k)
			}
			k = k + 1
		}
	}

}
