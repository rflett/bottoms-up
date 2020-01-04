package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	lowProbability  = 10
	medProbability  = 20
	highProbability = 30
)

type game struct {
	ID          string
	Probability int
	Name        string
	Description string
}

func getAllGames() []game {
	games := []game{
		{"DICE", highProbability, "Dice Roll", ""},
		{"CHICKS", highProbability, "Chicks", "All girls take a drink"},
		{"DICKS", highProbability, "Dicks", "All guys take a drink"},
		{"THUMBS", medProbability, "Thumbs", "Last person to put their thumb on the table takes a drink"},
		{"HEAVEN", medProbability, "Heaven", "Last person to put their hand up takes a drink"},
		{"SPINNER", medProbability, "Spinner", "Spin the shot spinner"},
		{"RHYME", medProbability, "Rhyme", "Make a rhyme, can't repeat, first to stumble, suck the teat"},
		{"CATEGORIES", medProbability, "Categories", "Start off with a category, say something that doesn't match and drink"},
		{"LAVA", lowProbability, "Lava", "THE DECK IS LAVA!"},
		{"POOL", lowProbability, "Pool", "LAST IN IN THE POOL DRINKS"},
		{"NORTO", lowProbability, "Nortoooo", "Norto, take a drink"},
		{"SEVENS", lowProbability, "Sevens", "Pick a number, each person subtracts 7, first to fuck up drinks"},
		{"WATERFALL", lowProbability, "Waterfall", "Drink for as long as the person to your right is drinking"},
	}

	return games
}

func getAGame() (game, int) {
	games := getAllGames()
	var pool []game

	// append every game to the pool X times where X is its probability
	for _, g := range games {
		i := g.Probability
		for i > 0 {
			pool = append(pool, g)
			i--
		}
	}

	// get a random number between 0 and the number of games in the pool
	rand.Seed(time.Now().UnixNano())
	gameChoice := rand.Intn(len(pool))

	// return the choice
	chosenGame := pool[gameChoice : gameChoice+1][0]

	// roll the dice
	diceRole := rand.Intn(6)

	if chosenGame.ID == "DICE" {
		youOrMe := rand.Intn(2)

		if youOrMe == 1 {
			chosenGame.Description = fmt.Sprintf("YOU need to take %d drink", diceRole)
		} else {
			chosenGame.Description = fmt.Sprintf("NOMINATE someone to have %d drink", diceRole)
		}

		if diceRole > 1 {
			chosenGame.Description += "s"
		}
	}

	return chosenGame, diceRole
}

func main() {
	r := gin.Default()

	r.GET("/game", func(c *gin.Context) {
		chosenGame, diceRole := getAGame()
		c.JSON(200, gin.H{
			"name":        chosenGame.Name,
			"description": chosenGame.Description,
			"dice_role":   diceRole,
		})
	})

	r.Run(":9000")

}
