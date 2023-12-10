package aoc2023

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func fixEngine(fileInput string) int {
	file, err := os.Open(fileInput)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
    result := 0
    var upperLine, currentLine string
    var matchedNumUpperIdx, unmatchedNumUpperIdx [][]int
    var matchedNumCurIdx, unmatchedNumCurIdx [][]int
    var auxMatched [][]int
    
    // Read First line
    scanner.Scan()
    currentLine = scanner.Text()
    // Find numbers first line
    unmatchedNumCurIdx = findNumbers(currentLine)
    // Get matched and unmatched numbers first line vs first line
    matchedNumCurIdx, unmatchedNumCurIdx = matchAdjacent(unmatchedNumCurIdx, currentLine)

	for scanner.Scan() {
        // Copy current line to upper line
        upperLine = currentLine
        unmatchedNumUpperIdx = unmatchedNumCurIdx
        // Read next line
		currentLine = scanner.Text()
        // Compare upper line to new line (diagonal)
        matchedNumUpperIdx, unmatchedNumUpperIdx = matchAdjacent(unmatchedNumUpperIdx, currentLine)
        matchedNumUpperIdx = append(matchedNumUpperIdx, matchedNumCurIdx...)
        // Find numbers new line
        unmatchedNumCurIdx = findNumbers(currentLine)
        // Get maatched and unmatched numbers new line vs upper line
        matchedNumCurIdx, unmatchedNumCurIdx = matchAdjacent(unmatchedNumCurIdx, upperLine)
        // Get matches and unmatched numbers new line vs new line
        auxMatched, unmatchedNumCurIdx = matchAdjacent(unmatchedNumCurIdx, currentLine)
        matchedNumCurIdx = append(matchedNumCurIdx, auxMatched...)
        // Sum matched upper line to results
        result += sumMatchedNumbers(matchedNumUpperIdx, upperLine)
	}
    // Sum result of last line
    result += sumMatchedNumbers(matchedNumCurIdx, currentLine)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return result
}

func findNumbers(line string) [][]int {
    return findIndex(line, `\d+`)
}

func findIndex(line, regex string) [][]int {
    return regexp.MustCompile(regex).FindAllStringIndex(line, -1)
}

func matchAdjacent(unmatchedNumsIdx [][]int, line string) (matched, unmatched [][]int) {
    //fmt.Printf("unmatched: %v, \n", unmatchedNumsIdx)
    matched = make([][]int, 0)
    unmatched = make([][]int, 0)
    for _, indexes := range unmatchedNumsIdx {
        // If beginning of line, don't check left
        leftLimit := indexes[0] - 1
        if indexes[0] == 0 {
            leftLimit = 0
        }
        // If end of line, don't check right
        rightLImit := indexes[1] + 1
        if indexes[1] == len(line) {
            rightLImit = len(line)
        }

        if substring := line[leftLimit : rightLImit]; isMatch(`[^\d\.]`, substring){
           // fmt.Printf("found match\n")
            matched = append(matched, indexes)
            //fmt.Printf("%v\n", matched)
        } else {
            //fmt.Printf("not found\n")
            unmatched = append(unmatched, indexes)
        }
    }
    return
}

func isMatch(regex, line string) bool {
    return regexp.MustCompile(regex).MatchString(line)
}

func sumMatchedNumbers(numbersIdx [][]int, line string) int {
    total := 0
    for _, idx := range numbersIdx {
        number, err := strconv.Atoi(line[idx[0]:idx[1]])
        if err != nil {
            panic(err)
        }
       // fmt.Printf("%v\n", number)
        total += number
    }
    return total
}
