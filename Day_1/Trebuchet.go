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
		re := regexp.MustCompile("(zerone|zero|oneight|one|twone|two|threeight|three|four|fiveight|five|six|sevenine|seven|eighthree|eight|nineight|nine)|[0-9]")
		//zerone, oneight, twone, threeight, fiveight, nineight,
		// re2 := regexp.MustCompile(re.ReplaceAllLiteralString())
		nums := re.FindAllString(line, -1)
		arrlen := len(nums)
		// fmt.Printf("String #%v: Text: %v Slice: %v Length: %v\n", ctr, line, nums, arrlen)
		// ctr++

		//if nums[0] == [0-9] jump to strconv.Atoi
		//if nums[0] == "one" replace it with "1" and so on for 0-9
		digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		digitstr := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		//fmt.Println(digits[0])
		// fmt.Printf("String #%v: Text: %v Slice: %v Length: %v\n", (ctr), line, nums, arrlen)

		//NumZero
		//fmt.Printf("Beginning string: %v\n",nums[0])
		switch {
		case nums[0] == "zerone":
			nums[0] = digits[0]
		case nums[0] == "twone":
			nums[0] = digits[2]
		case nums[0] == "eighthree":
			nums[0] = digits[8]
		case nums[0] == "oneight":
			nums[0] = digits[1]
		case nums[0] == "threeight":
			nums[0] = digits[3]
		case nums[0] == "fiveight":
			nums[0] = digits[5]
		case nums[0] == "nineight":
			nums[0] = digits[9]
		case nums[0] == "sevenine":
			nums[0] = digits[7]
		case nums[0] == digitstr[0]:
			nums[0] = digits[0]
			// fmt.Printf("Beginning string: %v\n", nums[0])
		case nums[0] == digitstr[1]:
			nums[0] = digits[1]
			// fmt.Printf("Beginning string: %v\n", nums[0])
		case nums[0] == digitstr[2]:
			nums[0] = digits[2]
			// fmt.Printf("Beginning string: %v\n", nums[0])
		case nums[0] == digitstr[3]:
			nums[0] = digits[3]
			// fmt.Printf("Beginning string: %v\n", nums[0])
		case nums[0] == digitstr[4]:
			nums[0] = digits[4]
			// fmt.Printf("Beginning string: %v\n", nums[0])
		case nums[0] == digitstr[5]:
			nums[0] = digits[5]
			// fmt.Printf("Beginning string: %v\n", nums[0])
		case nums[0] == digitstr[6]:
			nums[0] = digits[6]
			// fmt.Printf("Beginning string: %v\n", nums[0])
		case nums[0] == digitstr[7]:
			nums[0] = digits[7]
			// fmt.Printf("Beginning string: %v\n", nums[0])
		case nums[0] == digitstr[8]:
			nums[0] = digits[8]
			// fmt.Printf("Beginning string: %v\n", nums[0])
		case nums[0] == digitstr[9]:
			nums[0] = digits[9]
			// fmt.Printf("Beginning string: %v\n", nums[0])
		default:
			// fmt.Printf("Beginning string: %v\n", nums[0])
		}
		//NumOne
		switch {
		case nums[arrlen-1] == "zerone":
			nums[arrlen-1] = digits[1]
		case nums[arrlen-1] == "twone":
			nums[arrlen-1] = digits[1]
		case nums[arrlen-1] == "eighthree":
			nums[arrlen-1] = digits[3]
		case nums[arrlen-1] == "oneight":
			nums[arrlen-1] = digits[8]
		case nums[arrlen-1] == "threeight":
			nums[arrlen-1] = digits[8]
		case nums[arrlen-1] == "fiveight":
			nums[arrlen-1] = digits[8]
		case nums[arrlen-1] == "nineight":
			nums[arrlen-1] = digits[8]
		case nums[arrlen-1] == "sevenine":
			nums[arrlen-1] = digits[9]
		case nums[arrlen-1] == digitstr[0]:
			nums[arrlen-1] = digits[0]
			// fmt.Printf("Ending string: %v\n", nums[arrlen-1])
		case nums[arrlen-1] == digitstr[1]:
			nums[arrlen-1] = digits[1]
			// fmt.Printf("Ending string: %v\n", nums[arrlen-1])
		case nums[arrlen-1] == digitstr[2]:
			nums[arrlen-1] = digits[2]
			// fmt.Printf("Ending string: %v\n", nums[arrlen-1])
		case nums[arrlen-1] == digitstr[3]:
			nums[arrlen-1] = digits[3]
			// fmt.Printf("Ending string: %v\n", nums[arrlen-1])
		case nums[arrlen-1] == digitstr[4]:
			nums[arrlen-1] = digits[4]
			// fmt.Printf("Ending string: %v\n", nums[arrlen-1])
		case nums[arrlen-1] == digitstr[5]:
			nums[arrlen-1] = digits[5]
			// fmt.Printf("Ending string: %v\n", nums[arrlen-1])
		case nums[arrlen-1] == digitstr[6]:
			nums[arrlen-1] = digits[6]
			// fmt.Printf("Ending string: %v\n", nums[arrlen-1])
		case nums[arrlen-1] == digitstr[7]:
			nums[arrlen-1] = digits[7]
			// fmt.Printf("Ending string: %v\n", nums[arrlen-1])
		case nums[arrlen-1] == digitstr[8]:
			nums[arrlen-1] = digits[8]
			// fmt.Printf("Ending string: %v\n", nums[arrlen-1])
		case nums[arrlen-1] == digitstr[9]:
			nums[arrlen-1] = digits[9]
			// fmt.Printf("Ending string: %v\n", nums[arrlen-1])
		//Special case for twone to result in getting '1' instead of '2ne'

		default:
			// fmt.Printf("Ending string: %v\n", nums[arrlen-1])
		}

		fmt.Printf("String #%v: Text: %v Slice: %v Length: %v\n", (ctr), line, nums, arrlen)
		ctr++
		//NumZero Tens place, NumOne Ones Place
		for i := 0; i <= (arrlen - 1); i++ {
			// fmt.Printf("Nums List for str to int %v\n", nums[i])
		}
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
		// fmt.Println(sum)
	}
	fmt.Printf("Value %v", sum)
}
