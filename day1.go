package aoc2023

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func calibrate(filePath string) int {

    result := 0
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        
		line := scanner.Text()
        re := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
        repLine := re.ReplaceAllStringFunc(line, replace)
		if numbers := re.FindAllString(repLine, -1); len(numbers) > 1 {
            result += 10 * convertToInt(numbers[0]) + convertToInt(numbers[len(numbers) - 1])
        } else {
            result += 11 * convertToInt(numbers[0])
        }
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return result
}

func replace(match string) string {
    switch match {
    case "one":
        return "o1e"
    case "two":
        return "t2o"
    case "three":
        return "t3e"
    case "four":
        return "f4r"
    case "five":
        return "f5e"
    case "six":
        return "s6x"
    case "seven":
        return "s7n"
    case "eight":
        return "e8t"
    case "nine":
        return "n9e"
    default:
        return match 
    }
}

func convertToInt(match string) int {
    switch match {
    case "one":
        return 1
    case "two":
        return 2
    case "three":
        return 3
    case "four":
        return 4
    case "five":
        return 5
    case "six":
        return 6
    case "seven":
        return 7
    case "eight":
        return 8
    case "nine":
        return 9
    default:
        int, _ := strconv.Atoi(match)
        return int 
    }
}
