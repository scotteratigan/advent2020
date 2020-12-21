// define data structures and logic to read input

// 1. build out a slice of instructions
// instructions are structs - instruction/value/visited
// 2. when executing an instruction, set the visited bool to true
// 3. if visited is already set when you get to an instruction, halt (not implemented in this file)

package fakeassembly

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Code represents an instruction type: acc, jmp, nop
type Code int

const (
	// ACC is an accumulator
	ACC Code = iota
	// JMP is a jump statement
	JMP
	// NOP is no operation
	NOP
)

// InstructionBlock represents an instruction Code and value pair, along with a bool to track execution
type InstructionBlock struct {
	C    Code // acc, jmp, or nop
	Val  int
	Exec bool // whether or not a line has been executed
}

// ReadInputToStruct reads the struct
func ReadInputToStruct(filename string) []InstructionBlock {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	reg := regexp.MustCompile(`^(\w+) ([+\-]{1}\d+)$`)
	var ibSlice []InstructionBlock

	// For each line in input file, convert into InstructionBlock struct, add to slice
	for scanner.Scan() {
		m := reg.FindStringSubmatch(scanner.Text())
		instruction := m[1]
		value := m[2]
		var ib InstructionBlock

		switch instruction {
		case "acc":
			ib.C = ACC
		case "jmp":
			ib.C = JMP
		case "nop":
			ib.C = NOP
		default:
			log.Fatal("Tried to assign an invalid instruction code.")
		}

		i, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		ib.Val = i
		ibSlice = append(ibSlice, ib)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return ibSlice
}
