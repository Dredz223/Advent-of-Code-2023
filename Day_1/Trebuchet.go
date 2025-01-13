package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// nums := make([]int, 0)
	file, ferr := os.Open("einput.txt")
	if ferr != nil {
		panic(ferr)
	}
	arry := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		brry := make([]int, 0)
		num := ""
		// var numr int
		fmt.Println(num)
		line := scanner.Text()
		fmt.Printf("Text file: %v\n", line)

		re := regexp.MustCompile("[0-9]+")
		nums := re.FindAllString(line, -1)
		// numr, _ = strconv.Atoi(nums)
		// for _, char := range line {
		// 	if char >= '0' && char <= '9' {
		// 		num += string(char)
		// 		fmt.Println(num)

		// 	} else if num != "" {
		// 		n, _ := strconv.Atoi(num)
		// 		//nums = append(nums, n)
		// 		brry = append(brry, n)
		// 		num = ""
		// 	}
		// }
		fmt.Printf("Nums %v\n", nums)
		brry = append(brry, 0)
		arry = append(arry, brry)
	}
	// fmt.Println(nums)
	fmt.Println(arry)

}
