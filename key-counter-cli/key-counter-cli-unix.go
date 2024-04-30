package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("stty", "raw", "-echo")
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	defer func() {
		cmds := exec.Command("stty", "-raw", "echo")
		cmds.Stdin = os.Stdin
		errs := cmds.Run()
		if errs != nil {
			panic(errs)
		}
	}()
	var counter int = 0
	var count int = 0
	key := bufio.NewReader(os.Stdin)
	fmt.Print("Enter any key:")
	for {
		text, _, errors := key.ReadRune()
		if errors != nil {
			panic(errors)
		}
		if text == '\r' {
			counter += count
			count = 0
			fmt.Printf("\nCounter: %d", counter)
		} else if text == 27 {
			break
		} else {
			count++
		}
	}
}
