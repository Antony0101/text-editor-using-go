package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

// Set terminal to raw mode (works on Unix-based systems like Linux or macOS)
func enableRawMode() {
	cmd := exec.Command("stty", "raw", "-echo")
	cmd.Stdin = os.Stdin
	cmd.Run()
}

// Reset terminal to cooked mode (standard terminal mode)
func disableRawMode() {
	cmd := exec.Command("stty", "-raw", "echo")
	cmd.Stdin = os.Stdin
	cmd.Run()
}

// Clear the screen
func clearScreen() {
	fmt.Print("\033[2J")
}

// Move the cursor to a specific position
func moveCursor(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
}

// Read a single character from standard input
func readChar() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	char, err := reader.ReadByte()
	if err != nil {
		return "", err
	}
	return string(char), nil
}

func main() {
	// Clear the terminal screen and hide the cursor
	clearScreen()
	fmt.Print("\033[?25l")

	defer func() {
		// Clean up: reset the screen and show the cursor
		disableRawMode()
		clearScreen()
		fmt.Print("\033[?25h")
	}()

	// Set terminal to raw mode
	enableRawMode()

	// Starting position for the "player"
	x, y := 10, 10

	for {
		// Move the cursor to the current position of the "player"
		moveCursor(y, x)
		fmt.Print("@")

		// Wait for a key press (blocking)
		char, err := readChar()
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Handle input for movement
		switch char {
		case "w": // Move up
			y--
		case "s": // Move down
			y++
		case "a": // Move left
			x--
		case "d": // Move right
			x++
		case "q": // Quit the program
			return
		}

		// Clear the old position
		clearScreen()
	}
}
