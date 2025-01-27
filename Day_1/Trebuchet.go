package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	ctr := 1
	file, ferr := os.Open("input.txt")
	if ferr != nil {
		panic(ferr)
	}
	var sum int
	// var numZero int
	// var numOne int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Printf("Text%v ", line)
		re := regexp.MustCompile("([0-9]|eightwo|twone|oneight|nineight|one|two|three|four|five|six|seven|eight|nine)")
		//zerone, oneight, twone, threeight, fiveight, nineight,
		// re2 := regexp.MustCompile(re.ReplaceAllLiteralString())
		nums := re.FindAllString(line, -1)
		arrlen := len(nums)
		fstr := nums[0]
		lstr := nums[arrlen-1]
		// fmt.Printf("String #%v: Text: %v Slice: %v Length: %v\n", ctr, line, nums, arrlen)
		// ctr++

		//if fstr== [0-9] jump to strconv.Atoi
		//if fstr== "one" replace it with "1" and so on for 0-9
		digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		digitstr := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		//fmt.Println(digits[0])
		// fmt.Printf("String #%v: Text: %v Slice: %v Length: %v\n", (ctr), line, nums, arrlen)

		//NumZero
		//fmt.Printf("Beginning string: %v\n",nums[0])
		if fstr == "oneight" {
			fstr = "1"
		} else if fstr == "twone" {
			fstr = "2"
		} else if fstr == "eightwo" {
			fstr = "8"
		} else if fstr == "nineight" {
			fstr = "9"
		} else {
			switch fstr {
			case digitstr[0]:
				fstr = digits[0]
			case digitstr[1]:
				fstr = digits[1]
			case digitstr[2]:
				fstr = digits[2]
			case digitstr[3]:
				fstr = digits[3]
			case digitstr[4]:
				fstr = digits[4]
			case digitstr[5]:
				fstr = digits[5]
			case digitstr[6]:
				fstr = digits[6]
			case digitstr[7]:
				fstr = digits[7]
			case digitstr[8]:
				fstr = digits[8]
			case digitstr[9]:
				fstr = digits[9]
			default:
			}
		}
		//NumOne
		if lstr == "oneight" {
			lstr = "8"
		} else if lstr == "twone" {
			lstr = "1"
		} else if lstr == "eightwo" {
			lstr = "2"
		} else if lstr == "nineight" {
			lstr = "8"
		} else {
			switch lstr {
			case digitstr[0]:
				lstr = digits[0]
			case digitstr[1]:
				lstr = digits[1]
			case digitstr[2]:
				lstr = digits[2]
			case digitstr[3]:
				lstr = digits[3]
			case digitstr[4]:
				lstr = digits[4]
			case digitstr[5]:
				lstr = digits[5]
			case digitstr[6]:
				lstr = digits[6]
			case digitstr[7]:
				lstr = digits[7]
			case digitstr[8]:
				lstr = digits[8]
			case digitstr[9]:
				lstr = digits[9]
			default:
			}
		}
		fmt.Printf("String #%v: Text: %v Slice: %v Length: %v\n", (ctr), line, nums, arrlen)
		ctr++
		//NumZero Tens place, NumOne Ones Place
		// for i := 0; i <= (arrlen - 1); i++ {
		// 	// fmt.Printf("Nums List for str to int %v\n", nums[i])
		// }nums
		numZero, err := strconv.Atoi(fstr)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			// fmt.Printf("NumZero: %v\n", numZero)
		}
		numOne, err := strconv.Atoi(lstr)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			// fmt.Printf("NumOne: %v\n", numOne)
		}

		temp := numZero*10 + numOne
		fmt.Printf("Value before addition: %v\n", temp)
		sum += temp
		// fmt.Println(sum)
	}
	fmt.Printf("Value %v", sum)
}
