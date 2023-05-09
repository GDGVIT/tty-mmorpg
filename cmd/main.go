package main

import (
	"fmt"
	"tty-mmorpg/pkg/game"
)

func main() {
	// Print the title screen
	game.PrintTitleScreen()

	// Prompt the player to create a character
	fmt.Println("Press any key to continue...")
	fmt.Scanln()

	// Print the character creation screen
	game.PrintCharacterCreationScreen()

	// TODO: Implement character creation logic
	game.RollAbilityScores()

	// Start the game loop
	for {
		// TODO: Implement game logic
	}
}
