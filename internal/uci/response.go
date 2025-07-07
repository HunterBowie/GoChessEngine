package uci

import "fmt"

// Respond to a "uci" command
func uci() {
	fmt.Println("id name Luna")
	fmt.Println("id author Hunter Bowie")
	fmt.Println("uciok")
}

// Respond to a "isready" command
func isready() {
	fmt.Println("readyok")
}

// Respond to a timed "go" command
func goTimed(message string) {
	fmt.Println("bestmove d7d6")
}

// Respond to an infinite "go" command
func goInfinite() {
	fmt.Println("bestmove d7d6")
}
