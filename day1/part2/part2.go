package day1_part2

import (
	_ "embed"
	"fmt"
	// "fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var numbers = map[string]byte{
	"zero": '0',
	"one": '1',
	"two": '2',
	"three": '3',
	"four": '4',
	"five": '5',
	"six": '6',
	"seven": '7',
	"eight": '8',
	"nine": '9',
}

func findNumberFromKey(index int, element string, key string) byte {
	if len(key) > len(element) - index {
		return 'A'
	}

	for i := 0; i < len(key); i += 1 {
		if element[i + index] != key[i] {
			return 'A'
		}
	}

	return numbers[key]
}

func getNumber(index int, element string) byte {
	if element[index] >= '0' && element[index] <= '9' {
		return element[index];
	}

	for key := range numbers {
		if findNumberFromKey(index, element, key) == numbers[key] {
			return numbers[key]
		}
	}

	return 'A';
}

func Run(input string) string {

	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")



	var result int = 0; 
	for _, element := range lines {
		var first byte = 'A';
		var last byte = 'A';
		
		// Forward Loop
		for index := range element {
			character := getNumber(index, element);
			if character >= '0' && character <= '9' {
				if first == 'A' {
					first = character - '0'
				}		
				last = character - '0'
			}
		}



		if first != 'A' {
			fmt.Println(int(first) * 10 + int(last));
			result += int(first) * 10 + int(last);
		}
	}


	return strconv.Itoa(result)
}

func RunWithInput() string {
	return Run(input)
}

