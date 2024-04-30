package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("stty", "raw", "-echo") //execute command from shell for UNIX
	cmd.Stdin = os.Stdin                        //standard input
	err := cmd.Run()                            //run the command
	logger := log.New(os.Stdout, "", 0)         //use logger without printing current date and time
	if err != nil {
		panic(err) //stop if there is an error due to failed to run command from UNIX
	}
	defer func() {
		cmds := exec.Command("stty", "-raw", "echo")
		cmds.Stdin = os.Stdin
		errs := cmds.Run()
		if errs != nil {
			panic(errs) //stop if there is an error due to failed to run command from UNIX
		}
	}()
	var counter int = 0                                  //set counter to 0
	var count int = 0                                    //set key count to 0
	var squareCount int = 0                              //set square count to 0
	var triangleCount int = 0                            //set triangle count to 0
	key := bufio.NewReader(os.Stdin)                     //read standard input
	logger.Printf("Enter any key. Press ESC to quit.\r") //print directions
	for {
		text, _, errors := key.ReadRune() //read rune input (single character)
		if errors != nil {
			panic(errors) //stop if there is an error due to failed to read rune input
		}
		if text == '\r' { //when enter key is pressed
			count += counter                                       //increase key count by number of keys pressed
			squareCount += counter * counter                       //increase square count by count times count
			triangleCount += counter * (counter + 1) / 2           //increase triangle count by triangular number of count
			counter = 0                                            //set count to 0
			logger.Printf("Counter: %d\r", count)                  //log the key count
			logger.Printf("Square Counter: %d\r", squareCount)     //log the square count
			logger.Printf("Triangle Counter: %d\r", triangleCount) //log the triangle count
		} else if text == 27 { //when ESC key is pressed
			break //stop the program
		} else {
			counter++ //increase key count when key is pressed
		}
	}
}
