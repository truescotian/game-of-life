package main

import (
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// creates a random number from 0 <= n <= ARRAYSIZE
func getRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(ARRAYSIZE)
}

// clears the console
// NOTE: Only supports Linux.
func clearScreen() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
