package aoc2023

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

func calibrate(filePath string) int {
    
    file, err := os.Open(filePath)
    result := 0
    if err != nil {
        panic(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        var twoDigits string = ""
        for _, c := range line {
            if (unicode.IsNumber(c)) {
                twoDigits = twoDigits + string(c)
                break
            }
        }
        rl := []rune(line)
        for i := range rl {
            rc := rl[len(rl) - 1 - i]
            if (unicode.IsNumber(rc)) {
                twoDigits = twoDigits + string(rc)
                break
            }
        }
        if twoDigits == "" {
            break
        }
        number, err := strconv.Atoi(twoDigits)
        if err != nil {
            panic(err)
        }
        result += number
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
    return result
}
