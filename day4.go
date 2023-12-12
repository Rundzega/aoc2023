package aoc2023

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func scratchcards(fileInput string) int {
	file, err := os.Open(fileInput)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		card := strings.Split(scanner.Text(), ":")[1]
		winningNums := findNums(strings.Split(card, "|")[0])
		scratchedNums := findNums(strings.Split(card, "|")[1])
		result += computeCard(winningNums, scratchedNums)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return result
}

func madness(fileInput string) int {
	file, err := os.Open(fileInput)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	currentLine := 0
	scanner := bufio.NewScanner(file)
	repeatLines := make(map[int]int)
	totalScratched := 0
	for scanner.Scan() {
		currentLine++
		repeatLines[currentLine]++
		cardContents := strings.Split(scanner.Text(), ":")[1]
		winningNums := findNums(strings.Split(cardContents, "|")[0])
		scratchedNums := findNums(strings.Split(cardContents, "|")[1])
		bonus := 0
		bonus += computebonus(winningNums, scratchedNums) //* repeatLines[currentLine]
		totalScratched += repeatLines[currentLine]
		for i := 0; i < bonus; i++ {
			repeatLines[currentLine+1+i] += repeatLines[currentLine]
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return totalScratched
}

func findCardId(line string) int {
	numStr := regexp.MustCompile(`Card (\d+)`).FindStringSubmatch(line)[1]
	num, _ := strconv.Atoi(numStr)
	return num
}

func findNums(sequence string) map[int]bool {
	nums := make(map[int]bool)
	for _, numStr := range regexp.MustCompile(`\d+`).FindAllString(sequence, -1) {
		num, _ := strconv.Atoi(numStr)
		nums[num] = true
	}
	return nums
}

func computeCard(winNums, scratchedNums map[int]bool) int {
	result := 0
	count := 0
	for wNum := range winNums {
		if scratchedNums[wNum] {
			count++
		}
	}
	if count != 0 {
		result = int(math.Pow(2.0, float64(count-1)))
	}

	return result
}

func computebonus(winNums, scratchedNums map[int]bool) int {
	count := 0
	for wNum := range winNums {
		if scratchedNums[wNum] {
			count++
		}
	}

	return count
}
