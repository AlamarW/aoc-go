package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var runningNum = 0

func parseLine(text string) int {
	re := regexp.MustCompile(`\d`)
	list := re.FindAllString(text, -1)
	firstNum := list[0]
	lastNum := list[len(list)-1]
	if lastNum == "" {
		lastNum = firstNum
	}
	finalStrNum := firstNum + lastNum
	finalNum, err := strconv.Atoi(finalStrNum)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return finalNum
}

func runningSum(num int) {
	runningNum = runningNum + num
}

func readFile() {
	file, err := os.Open("raw.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		runningSum(parseLine(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	readFile()
	fmt.Println(runningNum)
}
