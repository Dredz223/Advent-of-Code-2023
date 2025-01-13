package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, ferr := os.Open("einput.txt")
	if ferr != nil {
		panic(ferr)
	}
	arry := make([][]int, 0, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num := ""
		fmt.Println(num)
		line := scanner.Text()
		fmt.Printf("Text file: %v\n", line)

		re := regexp.MustCompile("[0-9]+")
		nums := re.FindAllString(line, -1)
		fmt.Printf("Nums %v\n", nums)
	}
	fmt.Println(arry)

}
