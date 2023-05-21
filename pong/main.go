package main

import (
	"log"
	"time"

	"github.com/nsf/termbox-go"
)

// Game constants
const (
	boardWidth     = 60
	boardHeight    = 20
	paddleWidth    = 1
	paddleHeight   = 4
	ballChar       = '●'
	paddleChar     = '█'
	borderChar     = '═'
	borderTop      = '╔'
	borderBottom   = '╚'
	borderLeft     = '║'
	borderRight    = '║'
	ballSpeed      = time.Millisecond * 50
	paddleSpeed    = 2
	initialBallX   = boardWidth / 2
	initialBallY   = boardHeight / 2
	initialPaddleY = (boardHeight - paddleHeight) / 2
)

// Game variables
var (
	ballX, ballY         int
	ballXDirection       = 1
	ballYDirection       = 1
	playerPaddleY        int
	opponentPaddleY      int
	playerScore, opponentScore int
	gameOver             = false
)

func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	// Initialize game state
	resetGame()

	// Game loop
	for !gameOver {
		drawBoard()
		updateBall()
		updatePaddle()
		checkCollision()
		checkGameOver()
		termbox.Flush()
		time.Sleep(ballSpeed)
	}
}

func resetGame() {
	ballX = initialBallX
	ballY = initialBallY
	playerPaddleY = initialPaddleY
	opponentPaddleY = initialPaddleY
	playerScore = 0
	opponentScore = 0
	gameOver = false
}

func drawBoard() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Draw border
	for row := 0; row < boardHeight+2; row++ {
		termbox.SetCell(0, row, borderLeft, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(boardWidth+1, row, borderRight, termbox.ColorDefault, termbox.ColorDefault)
	}
	for col := 1; col < boardWidth+1; col++ {
		termbox.SetCell(col, 0, borderTop, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(col, boardHeight+1, borderBottom, termbox.ColorDefault, termbox.ColorDefault)
	}

	// Draw ball
	termbox.SetCell(ballX, ballY, ballChar, termbox.ColorDefault, termbox.ColorDefault)

	// Draw paddles
	for i := 0; i < paddleHeight; i++ {
		termbox.SetCell(1, playerPaddleY+i, paddleChar, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(boardWidth, opponentPaddleY+i, paddleChar, termbox.ColorDefault, termbox.ColorDefault)
	}

	// Draw scores
	termbox.SetCell(boardWidth/2-1, 0, '0'+rune(playerScore), termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(boardWidth/2+1, 0, '0'+rune(opponentScore), termbox.ColorDefault, termbox.ColorDefault)
}

func updateBall() {
	ballX += ballXDirection
	ballY += ballYDirection
}

func updatePaddle() {
	ev := termbox.PollEvent()
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyArrowUp:
			playerPaddleY -= paddleSpeed
		case termbox.KeyArrowDown:
			playerPaddleY += paddleSpeed
		}
	}
}

func checkCollision() {
	// Collision with top or bottom wall
	if ballY <= 0 || ballY >= boardHeight+1 {
		ballYDirection *= -1
	}

	// Collision with player paddle
	if ballX == 1 && ballY >= playerPaddleY && ballY < playerPaddleY+paddleHeight {
		ballXDirection *= -1
	}

	// Collision with opponent paddle
	if ballX == boardWidth && ballY >= opponentPaddleY && ballY < opponentPaddleY+paddleHeight {
		ballXDirection *= -1
	}

	// Scoring
	if ballX <= 0 {
		opponentScore++
		resetBall()
	}

	if ballX >= boardWidth+1 {
		playerScore++
		resetBall()
	}
}

func resetBall() {
	ballX = initialBallX
	ballY = initialBallY
	ballXDirection = -1
	ballYDirection = -1
}

func checkGameOver() {
	if playerScore >= 10 || opponentScore >= 10 {
		gameOver = true
	}
}

