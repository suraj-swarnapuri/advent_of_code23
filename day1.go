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
	numList := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	
	wordMap := make(map[string]int,0)

	for _,number := range(numList){
		index := strings.Index(input, number)
		if index != -1{
			wordMap[number] = index
		}
	}

	min,max := findMinMax(wordMap)


	return min*10+max
}

func findMinMax(m map[string]int) (int,int) {
	// Initialize min and max with the first element's value
	firstValue,lastValue := "",""
	min,max := 1000000,-1


	// Iterate over the map to find min and max
	for i, v := range m {
		if v < min {
			min = v
			firstValue = i
		}
		if v > max {
			max = v
			lastValue = i
		}
	}

	return numMap[firstValue], numMap[lastValue]
}
