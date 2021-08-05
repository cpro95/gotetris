package main

import (
	"time"
)

// Speeds
const slowestSpeed = 700 * time.Millisecond
const fastestSpeed = 60 * time.Millisecond

// Game play
const numSquares = 4
const numTypes = 7
const defaultLevel = 1
const maxLevel = 10
const rowsPerLevel = 5

// Pieces
var dxBank = [][]int{
	{},
	{0, 1, -1, 0},
	{0, 1, -1, -1},
	{0, 1, -1, 1},
	{0, -1, 1, 0},
	{0, 1, -1, 0},
	{0, 1, -1, -2},
	{0, 1, 1, 0},
}

var dyBank = [][]int{
	{},
	{0, 0, 0, 1},
	{0, 0, 0, 1},
	{0, 0, 0, 1},
	{0, 0, 1, 1},
	{0, 0, 1, 1},
	{0, 0, 0, 0},
	{0, 0, 1, 1},
}

type gameState int

const (
	gameIntro gameState = iota
	gameStarted
	gamePaused
	gameOver
)

// Struct Game contains all the game state.
type Game struct {
	board        [][]int // [y][x]
	state        gameState
	level        int
	numLines     int
	piece        int
	x            int
	y            int
	dx           []int
	dy           []int
	dxPrime      []int
	dyPrime      []int
	skyline      int
	fallingTimer *time.Timer
}

// NewGame returns a fully-initialized game.
func NewGame() *Game {
	g := new(Game)
	g.resetGame()
	return g
}

// Reset the game in order to play again.
func (g *Game) resetGame() {
	g.board = make([][]int, boardHeight)
	for y := 0; y < boardHeight; y++ {
		g.board[y] = make([]int, boardWidth)
		for x := 0; x < boardWidth; x++ {
			g.board[y][x] = 0
		}
	}

	g.state = gameIntro
	g.level = 1
	g.numLines = 0
	g.x = 1
	g.y = 1
	g.dx = []int{0, 0, 0, 0}
	g.dy = []int{0, 0, 0, 0}
	g.dxPrime = []int{0, 0, 0, 0}
	g.dyPrime = []int{0, 0, 0, 0}
	g.skyline = boardHeight - 1

	g.fallingTimer = time.NewTimer(time.Duration(1000000 * time.Second))
	g.fallingTimer.Stop()
}
