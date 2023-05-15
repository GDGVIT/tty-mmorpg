package game

import (
	"fmt"
	"math/rand"
	"time"
)

// Race represents a character's race
type Race struct {
	Name        string
	Description string
}

// Class represents a character's class
type Class struct {
	Name        string
	Description string
}

// AbilityScore represents an ability score
type AbilityScore struct {
	Name  string
	Value int
}

// Character represents a player character
type Character struct {
	Name          string
	Race          *Race
	Class         *Class
	AbilityScores map[string]*AbilityScore
}

// RollAbilityScores rolls ability scores for a character
func RollAbilityScores() map[string]*AbilityScore {
	scores := make(map[string]*AbilityScore)
	abilities := []string{"strength", "dexterity", "constitution", "intelligence", "wisdom", "charisma"}
	for _, ability := range abilities {
		fmt.Printf("\nRolling for %s!\n", ability)
		scores[ability] = &AbilityScore{Name: ability, Value: rollAbilityScore()}
		fmt.Printf("%s: %d\n", ability, scores[ability].Value)
	}
	return scores
}

// rollAbilityScore rolls a single ability score
func rollAbilityScore() int {
	rolls := []int{rollDie(), rollDie(), rollDie(), rollDie()}
	minIndex := 0
	for i := 0; i < len(rolls); i++ {
		fmt.Printf("Roll %d: %d\n", i, rolls[i])
	}
	for j := 0; j < len(rolls); j++ {
		for i := 0; i < len(rolls); i++ {
			if rolls[i] < rolls[minIndex] {
				minIndex = i
			}
		}
	}
	return rolls[0] + rolls[1] + rolls[2] + rolls[3] - rolls[minIndex]
}

// rollDie rolls a six-sided die
func rollDie() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}

// TODO: Add more rules and game mechanics as needed
