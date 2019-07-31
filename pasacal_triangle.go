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

	fmt.Println("read a number of rows for a pascal triangle")
	reader := bufio.NewReader(os.Stdin)
	num, err := reader.ReadString('\n')

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
	//fmt.Println(2 / 1)

	for i := 0; i < n; i++ {
		k := 1
		for k := 1; k < (n - i); k++ {
			fmt.Print("\t")
		}
		for j := 0; j <= i; j++ {
			if i == 0 || j == 0 {
				fmt.Print(k)
			} else {
				c := k * (i - j + 1)
				k = c / j
				fmt.Print(k)
			}

			fmt.Print("\t \t")
		}
		fmt.Println()

	}

}
