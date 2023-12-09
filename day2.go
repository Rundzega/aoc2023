package aoc2023

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func game(redMax, greenMax, blueMax int, fileInput string) int {
	file, err := os.Open(fileInput)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
    result := 0
	for scanner.Scan() {
		game := scanner.Text()
		id, err := strconv.Atoi(regexp.MustCompile(`^Game (\d+)`).FindStringSubmatch(game)[1])
		if err != nil {
			panic(err)
		}

        if findGreater(game, "red", redMax) || findGreater(game, "green", greenMax) || findGreater(game, "blue", blueMax) {
            continue
        }
        result += id
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return result
}

func findGreater(game, color string, qtd int) bool {
	for _, match := range regexp.MustCompile(`(\d+) ` + color).FindAllStringSubmatch(game, -1) {
		if greenQtd, _ := strconv.Atoi(match[1]); greenQtd > qtd {
            return true
		}
	}
    return false
}
