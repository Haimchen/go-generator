package main

import (
	"fmt"
	"math/rand"
)

type MazeField struct {
	visited bool
}

type Maze struct {
	field [][]MazeField
}

// initialize a new maze with the given dimensions
func newMaze(rows int, cols int) *Maze {
	maze := &Maze{}
	rowSlice := make([][]MazeField, rows, rows)
	for i := range rowSlice {
		rowSlice[i] = make([]MazeField, cols, cols)
	}
	maze.field = rowSlice
	return maze
}

// set initialized to true for field related to given point
// return 1 if value changed and 0 if not
func (maze *Maze) initializeField(point Point) int {
	visited := maze.field[point.y][point.x].visited
	if visited {
		return 0
	}

	maze.field[point.y][point.x].visited = true
	return 1
}

// Find all free (= not visited) neighbor fields
// returns an array of 0-4 points
func (maze *Maze) findFreeNeighbors(point Point) []Point {
	neighbors := make([]Point, 0, 4)
	//top
	if point.y > 0 {
		top := maze.field[point.y-1][point.x]
		if !top.visited {
			neighbors = append(neighbors, Point{x: point.x, y: point.y - 1})
		}
	}
	//bottom
	if point.y < len(maze.field)-1 {
		bottom := maze.field[point.y+1][point.x]
		if !bottom.visited {
			neighbors = append(neighbors, Point{x: point.x, y: point.y + 1})
		}
	}
	//left
	if point.x > 0 {
		left := maze.field[point.y][point.x-1]
		if !left.visited {
			neighbors = append(neighbors, Point{x: point.x - 1, y: point.y})
		}
	}
	//right
	if point.x < len(maze.field[0])-1 {
		right := maze.field[point.y][point.x+1]
		if !right.visited {
			neighbors = append(neighbors, Point{x: point.x + 1, y: point.y})
		}
	}
	return neighbors
}

func buildPath(rows int, cols int) []Point {
	maze := newMaze(rows, cols)
	totalFields := rows * cols
	fmt.Printf("maze: %d x %d\n", len(maze.field[0]), len(maze.field))
	path := make([]Point, 0, 2*totalFields)

	// random starting point in first column
	currentField := Point{x: 0, y: rand.Intn(rows)}
	path = append(path, currentField)
	initializedFields := 0 // TODO; try initializing currentField first

	counter := 0
	for initializedFields < totalFields {
		neighbors := maze.findFreeNeighbors(currentField)
		// there are no free neighbors, go back one step
		if len(neighbors) == 0 {
			counter -= 1
			currentField = path[counter]
			path = append(path, currentField)
			fmt.Println(currentField)
		} else {
			// there are some free neighbors, choose one as next
			nextField := neighbors[rand.Intn(len(neighbors))]
			path = append(path, nextField)
			currentField = nextField
			counter = len(path) - 1
		}
		initializedFields += maze.initializeField(currentField)
	}
	return path
}

func connectionsFromPath(path []Point) [][2]Point {
	connections := make([][2]Point, 0, len(path))
	// go through all path elements until the second to last
	for i := 0; i < len(path)-1; i++ {
		connections = append(connections, [2]Point{path[i], path[i+1]})
	}
	return connections
}

func unique(connections [][2]Point) [][2]Point {
	unique := make([][2]Point, 0, len(connections))

	for _, conn := range connections {
		if !isIn(unique, conn) {
			unique = append(unique, conn)
		}
	}

	return unique
}

func isIn(collection [][2]Point, item [2]Point) bool {
	for _, elem := range collection {
		if (elem[0] == item[0] && elem[1] == item[1]) ||
			(elem[0] == item[1] && elem[1] == item[0]) {
			return true
		}
	}
	return false
}

func BuildConnections(rows int, cols int) [][2]Point {
	path := buildPath(rows, cols)
	connections := connectionsFromPath(path)
	return unique(connections)
}
