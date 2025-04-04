package main

import (
	"fmt"
	"math/rand"
)

type character struct {
	class        string
	strength     int
	dexterity    int
	constitution int
	intelligence int
	wisdom       int
	charisma     int
}

func main() {
	var character = generateCharacter()
	fmt.Printf("Class: %s\n", character.class)
	fmt.Printf("Str: %d\n", character.strength)
	fmt.Printf("Dex: %d\n", character.dexterity)
	fmt.Printf("Con: %d\n", character.constitution)
	fmt.Printf("Int: %d\n", character.intelligence)
	fmt.Printf("Wis: %d\n", character.wisdom)
	fmt.Printf("Cha: %d\n", character.charisma)
}

func generateCharacter() character {
	var statList = rollStatList()
	var maxStatIndex = getMaxStatIndex(statList)
	var class = getClass(maxStatIndex)
	var character = character{
		class,
		statList[0],
		statList[1],
		statList[2],
		statList[3],
		statList[4],
		statList[5],
	}
	return character
}

func rollStatList() []int {
	var statList = make([]int, 0)
	for range 6 {
		statList = append(statList, rollStat())
	}
	return statList
}

func rollStat() int {
	return rollDice(3, 6)
}

func rollDice(amount int, sides int) int {
	var sum = 0
	for range amount {
		sum += rand.Intn(sides) + 1
	}
	return sum
}

func getMaxStatIndex(statList []int) int {
	var maxIndex = 0
	var maxValue = 0
	for index, value := range statList {
		if index != 2 && (value > maxValue || value == maxValue && rand.Intn(2) == 1) {
			maxIndex = index
			maxValue = value
		}
	}
	return maxIndex
}

func getClass(index int) string {
	switch index {
	case 0:
		return "Fighter"
	case 1:
		return "Thief"
	case 3:
		return "Wizard"
	case 4:
		return "Cleric"
	case 5:
		return "Bard"
	}
	return "Joker"
}
