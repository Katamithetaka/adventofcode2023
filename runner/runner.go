package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"takathedinosaur.tech/day1_part1"
	"takathedinosaur.tech/day1_part2"
)

func main() {
	fmt.Println("Select which day (from 1 - 1)")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", ""), "\n", "");
	day, _ := strconv.Atoi(text)

	fmt.Println("Select which part (from 1 - 2)")
	text, _ = reader.ReadString('\n')
	text = strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", ""), "\n", "");

	part, _ := strconv.Atoi(text)

	fmt.Printf("Day %d Part %d", day, part)
	var result string;
	switch day {
	case 0:
	case 1:
		switch part {
		case 0:
		case 1:
			result = day1_part1.RunWithInput()
			break
		case 2:
			result = day1_part2.RunWithInput()
			break
		default:
			_ = fmt.Errorf("Unimplemented.")
			break	
		}
		break
	default:
		_ = fmt.Errorf("Unimplemented.")
		return

	}

	println(result)

}
