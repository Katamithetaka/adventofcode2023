package day1_part1


import (
	_ "embed"
	// "fmt"
	"strings"
	"strconv"
)

//go:embed input.txt
var input string

func Run(input string) string {

	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	var result int = 0; 
	for _, element := range lines {
		var first byte = 'A';
		var last byte = 'A';
		
		// Forward Loop
		for index := range element {
			character := element[index]
			if character >= '0' && character <= '9' {		
				first = character - '0'
				break;		
			}
		}

		// Reverse loop
		for index := range element {
			character := element[len(element) - 1 - index];

			if character >= '0' && character <= '9' {		
				last = character - '0'
				break;
			}
		}

		if first != 'A' {
			// fmt.Println(int(first) * 10 + int(last));
			result += int(first) * 10 + int(last);
		}
	}


	return strconv.Itoa(result)
}

func RunWithInput() string {
	return Run(input)
}


