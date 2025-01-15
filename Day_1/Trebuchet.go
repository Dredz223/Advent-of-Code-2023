package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, ferr := os.Open("input.txt")
	if ferr != nil {
		panic(ferr)
	}
	ctr := 1
	var sum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile("[0-9]")
		nums := re.FindAllString(line, -1)
		ctr++
		arrlen := len(nums)

		//NumZero Tens place, NumOne Ones Place
		numZero, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println("Error", err)
		} else {
			// fmt.Printf("NumZero: %v\n", numZero)
		}
		numOne, err := strconv.Atoi(nums[arrlen-1])
		if err != nil {
			fmt.Println("Error", err)
		} else {
			// fmt.Printf("NumOne: %v\n", numOne)
		}
		temp := numZero*10 + numOne
		sum += temp
	}
	fmt.Printf("Value %v", sum)
}
