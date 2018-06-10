package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func readNumber(reader *bufio.Reader) int {
	var input, _, _ = reader.ReadLine()
	var number, _ = strconv.Atoi(string(input))
	return number
}

func readCoordinate(reader *bufio.Reader) (int, int) {
	var input, _, _ = reader.ReadLine()
	var coordinateValues []string = strings.Split(string(input), " ")
	var x, _ = strconv.Atoi(coordinateValues[0])
	var y, _ = strconv.Atoi(coordinateValues[1])
	return x, y
}

func getNumberOfMovesToCompleteGame(positions []int, shoots map[int]int, numberOfTurns int) int {
	fmt.Println(positions)
	var nextPositions []int
	for _, position := range positions {
		for i := 0; i < 6; i++ {
			var space = position - i
			if space == 1 {
				return numberOfTurns
			}
			end, ok := shoots[space]
			if ok == true {
				nextPositions = append(nextPositions, end)
			}
		}
		var moveDiceRoll = position - 6
		if moveDiceRoll <= 1 {
			return numberOfTurns
		}
		nextPositions = append(nextPositions, position-6)
	}
	if numberOfTurns > 4 {
		return numberOfTurns
	}
	return getNumberOfMovesToCompleteGame(nextPositions, shoots, numberOfTurns+1)
}

func getLaddersAndSnakes(reader *bufio.Reader) map[int]int {
	var shoots map[int]int = make(map[int]int)
	var numberOfLadders = readNumber(reader)
	for ladderCnt := 0; ladderCnt < numberOfLadders; ladderCnt++ {
		var start, end = readCoordinate(reader)
		shoots[end] = start
	}
	var numberOfSnakes = readNumber(reader)
	for snakeCnt := 0; snakeCnt < numberOfSnakes; snakeCnt++ {
		var start, end = readCoordinate(reader)
		shoots[end] = start
	}
	return shoots
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	var numberOfProblems = readNumber(reader)
	for i := 0; i < numberOfProblems; i++ {
		var shoots map[int]int = getLaddersAndSnakes(reader)
		fmt.Println(shoots)
		var minNumberOfTurns = getNumberOfMovesToCompleteGame([]int{100}, shoots, 1)
		fmt.Println(minNumberOfTurns)
	}
}

func TestOne(t *testing.T) {

}
