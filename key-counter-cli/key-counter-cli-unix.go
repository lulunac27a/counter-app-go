package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("stty", "raw", "-echo")
	cmd.Stdin = os.Stdin //standard input
	err := cmd.Run()     //run the command
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
	var counter int = 0              //set counter to 0
	var count int = 0                //set key count to 0
	key := bufio.NewReader(os.Stdin) //read standard input
	fmt.Print("Enter any key:")      //print the string
	for {
		text, _, errors := key.ReadRune() //read rune input
		if errors != nil {
			panic(errors)
		}
		if text == '\r' { //when enter key is pressed
			counter += count //increase counter by number of keys pressed
			count = 0        //set count to 0
			fmt.Printf("\nCounter: %d", counter)
		} else if text == 27 { //when ESC key is pressed
			break
		} else {
			count++ //increase key count when key is pressed
		}
	}
}
