package interpreter

import (
	"errors"
	"log"
)

// Much of the design of this project comes from the Wikipedia article on Brainfuck at https://en.wikipedia.org/wiki/Brainfuck
// To the many authors of the article, thank you!

// Program settings are all inside this constant block
const (
	fieldSize uint = 30000 // The size of the array that brainfuck may operate on. Using 30000 for now as the wikipedia article says is the minimum required.

	incPointerSymbol   rune = '>' // Move the data pointer to the right
	decPointerSymbol   rune = '<' // Move the data pointer to the left
	incFieldSymbol     rune = '+' // Increment the value at the data pointer
	decFieldSymbol     rune = '-' // Decrement the value at the data pointer
	outByteSymbol      rune = '.' // Output the byte at the data pointer
	inByteSymbol       rune = ',' // Take one byte as input and set the value at the data pointer to that byte
	jumpForwardSymbol  rune = '[' // If the byte at the data pointer is zero, jump to the command after the matching ]
	jumpBackwardSymbol rune = ']' // If the byte at the data pointer is not zero, jump to the command before the matching [
)

type bfProgram []rune

type BF struct {
	dataPointer        uint            // The data pointer, as described in the Wikipedia article
	instructionPointer uint            // The instruction pointer, as described in the Wikipedia article
	field              [fieldSize]byte // The long array, or "tape," that brainfuck uses to execute code.
	program            bfProgram       // The program that is currently running
	CycleCounter       uint            // Keeps count of how many cycles this program has taken to run
	Result             []byte          // A copy of the results, identical to what is printed
}

func (bf *BF) LoadProgram(program string) {
	if bf.program != nil {
		log.Fatal("Tried to load new program into already existing interpreter instance")
	}
	bf.program = []rune(program)
}

func (bf *BF) Execute() {
	for bf.instructionPointer < uint(len(bf.program)) {
		switch bf.program[bf.instructionPointer] {

		case incPointerSymbol:
			bf.incPointer()
		case decPointerSymbol:
			bf.decPointer()
		case incFieldSymbol:
			bf.incField()
		case decFieldSymbol:
			bf.decField()
		case outByteSymbol:
			bf.outByte()
		case inByteSymbol:
			bf.inByte(0)
		case jumpForwardSymbol:
			bf.jumpForward()
		case jumpBackwardSymbol:
			bf.jumpBackward()
		default:
			bf.instructionPointer++
			continue
		}
		// fmt.Printf("Data pointer, instruction pointer/current instruction, current byte/char after cycle #%v: %v, %v/%c, %X/%c\n", bf.CycleCounter, bf.dataPointer, bf.instructionPointer, bf.program[bf.instructionPointer], bf.readByte(), bf.readByte())
		bf.instructionPointer++
		bf.CycleCounter++
	}
}

// Interpreter helper functions

// Gets the byte at the data pointer
func (bf *BF) readByte() byte {
	return bf.field[bf.dataPointer]
}

// Gets the rune at the data pointer, assuming it is meant to be one
func (bf *BF) readRune() rune {
	return rune(bf.readByte())
}

// Gets the current instruction
func (bf *BF) currentInstruction() rune {
	return bf.program[bf.instructionPointer]
}

// Increment a uint within certain bounds, and
func incWithBound(i uint, bound uint) uint {
	return (i + 1) % bound
}

// Decrement a uint within certain bounds
func decWithBound(i uint, bound uint) uint {
	if i == 0 {
		return bound
	}
	return i - 1
}

// Actual brainfuck interpreting functions

// Move the data pointer to the right
func (bf *BF) incPointer() {
	bf.dataPointer = incWithBound(bf.dataPointer, fieldSize-1)
}

// Move the data pointer to the left
func (bf *BF) decPointer() {
	bf.dataPointer = decWithBound(bf.dataPointer, fieldSize-1)
}

// Increment the value at the data pointer
func (bf *BF) incField() {
	bf.field[bf.dataPointer]++
}

// Decrement the value at the data pointer
func (bf *BF) decField() {
	bf.field[bf.dataPointer]--
}

// Output the byte at the data pointer as a character
func (bf *BF) outByte() {
	print(string(bf.readRune()))
	bf.Result = append(bf.Result, bf.readByte())
}

// Take one byte as input and set the value at the data pointer to that byte
func (bf *BF) inByte(input byte) {
	// Not implemented yet. Do not use!
	panic(errors.New("used unimplemented code"))
}

// If the byte at the data pointer is zero, jump to the command after the matching ]
func (bf *BF) jumpForward() {
	if bf.readByte() == 0 {
		// Seeks for the end of the loop
		for bf.currentInstruction() != ']' {
			bf.instructionPointer++
			if int(bf.instructionPointer) > len(bf.program) {
				log.Fatal("error: couldn't find end point of loop")
			}
		}
	}
}

func (bf *BF) jumpBackward() {
	if bf.readByte() != 0 {
		// Seeks for the start of the loop
		for bf.currentInstruction() != '[' {
			if bf.instructionPointer == 0 {
				log.Fatal("error: couldn't find return point for loop")
			}
			bf.instructionPointer--
		}
	}
}

// Print functions
