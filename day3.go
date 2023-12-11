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

func fixGear(fileInput string) int {
	file, err := os.Open(fileInput)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
    result := 0
    
    // Initialize all lines 
    scanner.Scan()
    topLine := ""
    currentLine := ""
    bottomLine := scanner.Text()
	for scanner.Scan() {
        topLine = currentLine
        currentLine = bottomLine
        bottomLine = scanner.Text()
        result += computeGearLine(currentLine, topLine, bottomLine)
	}
    // compute Last line
    result += computeGearLine(bottomLine, currentLine, "")

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return result

}

func findNumbers(line string) [][]int {
    return findIndex(line, `\d+`)
}

func findStars(line string) []int {
    starsIdxPair := findIndex(line, `\*`)
    starsIdx := make([]int, 0)
    
    for _, idx := range starsIdxPair {
        starsIdx = append(starsIdx, idx[0])
    }
    return starsIdx
}

func findIndex(line, regex string) [][]int {
    return regexp.MustCompile(regex).FindAllStringIndex(line, -1)
}

func matchAdjacent(unmatchedNumsIdx [][]int, line string) (matched, unmatched [][]int) {
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
            matched = append(matched, indexes)
        } else {
            unmatched = append(unmatched, indexes)
        }
    }
    return
}

func isMatch(regex, line string) bool {
    return regexp.MustCompile(regex).MatchString(line)
}

func computeGearLine(gearLine, aboveLine, belowLine string) int {
    lineTotal := 0
    for _, starIdx := range findStars(gearLine) {
        qtdAdj := 0
        gearRatio := 1
        qtdAdjLine := 0
        lineGearRatio := 1
        if (aboveLine != "") {
            aboveNums := findNumbers(aboveLine)
            qtdAdjLine, lineGearRatio = checkAdjacentToStar(starIdx, aboveNums, aboveLine)
            qtdAdj += qtdAdjLine
            gearRatio *= lineGearRatio
        }
        if (belowLine != "") {
            belowNums := findNumbers(belowLine)
            qtdAdjLine, lineGearRatio = checkAdjacentToStar(starIdx, belowNums, belowLine)
            qtdAdj += qtdAdjLine
            gearRatio *= lineGearRatio
        }
        gearLineNums := findNumbers(gearLine)
        qtdAdjLine, lineGearRatio = checkAdjacentToStar(starIdx, gearLineNums, gearLine)
        qtdAdj += qtdAdjLine
        gearRatio *= lineGearRatio
        if qtdAdj == 2 {
            lineTotal += gearRatio
        }
    }
    return lineTotal
}

func checkAdjacentToStar(starIdx int, numsPos [][]int, line string) (qtdAjd, gearRatio int) {
    qtdAjd = 0
    gearRatio = 1
    matchedNums := make([][]int, 0)
    for _, numPos := range numsPos {
        
        numStart := numPos[0]
        numEnd := numPos[1] - 1
        starIdxStart := starIdx - 1
        if starIdxStart < 0 {
            starIdxStart = 0
        }
        starIdxEnd := starIdx + 1
        if starIdxEnd == len(line) {
            starIdxEnd = starIdx
        }
        if (numStart >= starIdxStart && numStart <= starIdxEnd) ||
            (numEnd >= starIdxStart && numEnd <= starIdxEnd) ||
            (starIdxStart <= numEnd && starIdxStart >= numStart) ||
            (starIdxEnd <= numEnd && starIdxEnd >= numStart) {
                qtdAjd++
                matchedNums = append(matchedNums, numPos)
            }
    }
    if qtdAjd > 0 {
        gearRatio = multiplyMatchedNumbers(matchedNums, line)
    } 
    return
}


func sumMatchedNumbers(numbersIdx [][]int, line string) int {
    total := 0
    for _, idx := range numbersIdx {
        number, err := strconv.Atoi(line[idx[0]:idx[1]])
        if err != nil {
            panic(err)
        }
        total += number
    }
    return total
}

func multiplyMatchedNumbers(numbersIdx [][]int, line string) int {
    total := 1
    for _, idx := range numbersIdx {
        number, err := strconv.Atoi(line[idx[0]:idx[1]])
        if err != nil {
            panic(err)
        }
        total *= number
    }
    return total
}

