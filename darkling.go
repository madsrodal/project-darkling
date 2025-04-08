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

const (
	Strength     int = 0
	Dexterity    int = 1
	Constitution int = 2
	Intelligence int = 3
	Wisdom       int = 4
	Charisma     int = 5
)

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
	var stats = rollStats()
	var classStat = getClassStat(stats)
	var class = getClass(classStat)
	return character{
		class,
		stats[Strength],
		stats[Dexterity],
		stats[Constitution],
		stats[Intelligence],
		stats[Wisdom],
		stats[Charisma],
	}
}

func rollStats() []int {
	var stats = make([]int, 0)
	for range 6 {
		stats = append(stats, rollDice(3, 6))
	}
	return stats
}

func rollDice(amount int, sides int) int {
	var sum = 0
	for range amount {
		sum += rand.Intn(sides) + 1
	}
	return sum
}

func getClassStat(stats []int) int {
	var maxStat = 0
	var maxValue = 0
	for stat, value := range stats {
		if stat != Constitution && (value > maxValue || value == maxValue && rand.Intn(2) == 1) {
			maxStat = stat
			maxValue = value
		}
	}
	return maxStat
}

func getClass(classStat int) string {
	switch classStat {
	case Strength:
		return "Fighter"
	case Dexterity:
		return "Thief"
	case Intelligence:
		return "Wizard"
	case Wisdom:
		return "Cleric"
	case Charisma:
		return "Bard"
	}
	return "Joker"
}
