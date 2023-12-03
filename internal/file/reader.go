package file
import(
	"fmt"
	"os"
	"bufio"

)


func LoadFile(path string)([]string,error){

	var outputList []string
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate through each line
	for scanner.Scan() {
		line := scanner.Text()
		outputList = append(outputList,line)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, scanner.Err()
	}

	return outputList, nil
}