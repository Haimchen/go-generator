package main

import (
	"fmt"
)

func printWorld(world *World) {
	fmt.Println("Next Iteration:")
	for _, row := range world.Data {
		printRow(row[:])
	}
}

func printRow(row []Block) {
	rowString := ""
	for _, block := range row {
		rowString = fmt.Sprintf("%s| %s ", rowString, printBlock(block))
	}
	fmt.Println(rowString + "|")
}

func printBlock(block Block) string {
	switch block {
	case 0:
		return "_"
	case 1:
		return "X"
	default:
		return "_"
	}
}
