package template_pattern

import (
	"fmt"
)

type Game interface {
	Start()
	TakeTurn()
	HaveWinner() bool
	WinningPlayer() int
}

func PlayGame(g Game) {
	g.Start()
	for !g.HaveWinner() {
		g.TakeTurn()
	}

	fmt.Printf("Player  %d wins. \n", g.WinningPlayer())
}

type chess struct {
	turn, maxTurns, currentPlayer int
}

func NewGameOfChess() Game {
	return &chess{1, 10, 0}
}

func (c *chess) Start() {
	fmt.Println("Starting a new game of chess.")
}

func (c *chess) TakeTurn() {
	c.turn++
	fmt.Printf("Turn %d  taken by player %d\n", c.turn, c.currentPlayer)
	c.currentPlayer = 1 - c.currentPlayer
}

func (c *chess) HaveWinner() bool {
	return c.turn == c.maxTurns
}

func (c *chess) WinningPlayer() int {
	return c.currentPlayer
}

func PlayGame2(start, takeTurn func(), haveWinner func() bool, winningPlayer func() int) {
	for !haveWinner() {
		takeTurn()
	}
	fmt.Printf("Player %d wins.  \n", winningPlayer())
}

func ExecuteTemplateFunction(){
	fmt.Println("template function pattern")
	turn, maxTurns, currentPlayer := 1, 10, 0

	start := func() {
		fmt.Println("Starting a game of chess.")
	  }
	
	  takeTurn := func() {
		turn++
		fmt.Printf("Turn %d taken by player %d\n",
		  turn, currentPlayer)
		currentPlayer = (currentPlayer + 1) % 2
	  }
	
	  haveWinner := func()bool {
		return turn == maxTurns
	  }
	
	  winningPlayer := func()int {
		return currentPlayer
	  }
	
	  PlayGame2(start, takeTurn, haveWinner, winningPlayer)
}

func ExecuteTemplateMethod() {
	chess := NewGameOfChess()
	PlayGame(chess)

}
