package main
import(
	"fmt"
	"strconv"
	"strings"
	"github.com/suraj-swarnapuri/advent_of_code23/internal/file"
)

var numMap = map[string]int{
	"1":1,
	"2":2,
	"3":3,
	"4":4,
	"5":5,
	"6":6,
	"7":7,
	"8":8,
	"9":9,
	"one":1,
	"two":2,
	"three":3,
	"four":4,
	"five":5,
	"six":6,
	"seven":7,
	"eight":8,
	"nine":9,
}

func main(){
	
	day_two()
}

func day_one(){
	inputList,err := file.LoadFile("data/day1.txt")
	if err != nil{
		panic(err)
	}
	total := 0
	for _, input := range(inputList){
		total = total + findCalibrationValue(input)
	}
	fmt.Println()
	fmt.Printf("The Total is: %d\n",total)
}

func day_two(){
	inputList,err := file.LoadFile("data/day1.txt")
	if err != nil{
		panic(err)
	}
	total := 0
	for _, input := range(inputList){
		total = total + findCalibrationValueDayTwo(input)
	}
	fmt.Println()
	fmt.Printf("The Total is: %d\n",total)
}


func findCalibrationValue(input string)int {
	firstIndex,lastIndex := 0,len(input)-1
	firstFound,lastFound := false, false
	firstValue,lastValue := 0,0
	var err error
	for true {
		if firstFound && lastFound{
			break
		}
		firstValue, err = strconv.Atoi(string(input[firstIndex]))
		if err != nil{
			firstIndex = firstIndex+1
		} else{
			firstFound = true
		}

		lastValue, err = strconv.Atoi(string(input[lastIndex]))
		if err != nil{
			lastIndex = lastIndex-1
		} else{
			lastFound = true
		}

	}
	calibrationValue := firstValue*10+lastValue
	fmt.Printf("The Calibration Value is : %d\n",calibrationValue)
	return calibrationValue
}

func findCalibrationValueDayTwo(input string) int{
	arg := extractDigitsAndSpeltStrings(input)
	return arg[0]*10+arg[len(arg)-1]
}

func extractDigitsAndSpeltStrings(str string) []int {
	spellingMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	digits := []int{}
	for i, char := range str {
		if digit, err := strconv.Atoi(string(char)); err == nil {
			digits = append(digits, digit)
		} else {
			for spelling, number := range spellingMap {
				if strings.HasPrefix(str[i:], spelling) {
					digits = append(digits, number)
					break
				}
			}
		}
	}

	return digits
}
