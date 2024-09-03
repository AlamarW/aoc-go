package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// type Game struct {
// 	GameId    int
// 	RedCube   int
// 	GreenCube int
// 	BlueCube  int
// }

//go:embed raw.txt
var input string

var gameNumbers = []int{}

var LIMITS = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func sliceString() []string {
	text := strings.Split(input, "\n")
	return text
}

func getGameNumbers(text string) {
	re := regexp.MustCompile(`\d+`)
	gameStrId := re.FindString(text)
	gameId, err := strconv.Atoi(gameStrId)
	if err != nil {
		fmt.Println(err)
	}
	gameNumbers = append(gameNumbers, gameId)
}

func checkValidGame(text string) {
	re := regexp.MustCompile(`\d+ (.+)`)
	gameInstances := re.FindAllString(text, -1)
	gameInstancesSplit := strings.Split(gameInstances[0], ";")
	for _, el := range gameInstancesSplit {
		setSplit := strings.Split(el, ",")
		test := setSplit[0]
		test2 := test[0]

		fmt.Println(test)
		fmt.Println("Leading character")
		fmt.Println(string(test2))
	}
}

// create a funct for possible games, store the game ID of those games ONLY, and then sum
// We actually don't need a struct with all of these, we only realy need to hold onto the int of the game
// but that can just be stored or +='d

func main() {
	bobby := sliceString()
	// fmt.Println(bobby[19])
	getGameNumbers(bobby[19])
	fmt.Println(gameNumbers)
	checkValidGame("Game 100: 4 red, 2 blue, 4 green; 2 green, 1 red, 1 blue; 3 green, 4 blue, 5 red; 18 red, 2 blue; 9 red, 5 green, 4 blue")
}
