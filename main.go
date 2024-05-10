package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Collig0/gofuck/interpreter"
)

var bfProgramFile string
var bfInput = flag.String("input", "", "Predetermined input for the Brainfuck program")

func main() {
	brainfuckProgram := new(interpreter.BF)
	brainfuckProgram.LoadProgram(bfProgramFile)
	brainfuckProgram.Input = []byte(*bfInput)
	brainfuckProgram.Execute()
	print("\n")
	fmt.Println("Result:", brainfuckProgram.Result)
	fmt.Printf("Finished in %v cycles.\n", brainfuckProgram.CycleCounter)
}

func init() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("Please provide the name of exactly one brainfuck file to execute.")
		os.Exit(1)
	} else {
		bfProgramBytes, err := os.ReadFile(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		bfProgramFile = string(bfProgramBytes)
	}
}
