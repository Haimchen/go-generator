package main

import (
	"fmt"
	"math/rand"
)

// configure this level
const worldSize = 25
const fullnessFactor = 0.4

var roomTypes = []Room{
	Room{height: 3, width: 3},
	Room{height: 3, width: 4},
	Room{height: 4, width: 3},
	Room{height: 4, width: 4},
}

// types
type World struct {
	Data   [worldSize][worldSize]Block
	Chunks []Chunk
}

type Block int

type Chunk struct {
	top    Point
	bottom Point
}

type Point struct {
	x int
	y int
}

type Room struct {
	height int
	width  int
}

func main() {
	fmt.Println("starting world generation")

	// initialize empty world matrix
	world := &World{}
	printWorld(world)

	// recursively split world into chunks
	// stop when next split would lead to chunks below minimum size (depending on FULLNESS_FACTOR)
	chunkHeight, chunkWidth := chunkSize()
	fmt.Printf("Chunk dimensions: %d x %d (w x h)\n", chunkWidth, chunkHeight)

	world.Chunks = chunks(world, chunkWidth, chunkHeight)
	for i, chunk := range world.Chunks {
		fmt.Printf("%d. Chunk: (%d/%d) to (%d/%d)\n", i, chunk.top.x, chunk.top.y, chunk.bottom.x, chunk.bottom.y)
	}

	// create one room per chunk
	// for each room randomly pick a room type from the list of available types
	buildRooms(world)

	printWorld(world)

	// make a new maze with all the chunks
	// define a starting chunk
	// find a path that connects all chunks ( recursive backtracker?)
	// result: all connections between chunks
	chunkRows := worldSize / chunkHeight
	chunkCols := worldSize / chunkWidth
	connections := BuildConnections(chunkRows, chunkCols)
	fmt.Printf("Built %d connections:\n", len(connections))
	for _, points := range connections {
		fmt.Println(points)
	}

	// create random tunnels connecting the rooms for which a connection was defined
	// find entrance and exit fields for each room?

	// define a starting Room for the player
	// could be the room of the starting chunk?

	//------ Prettify -------------------------

	// define different ground types per block (from a set defined in the config)
	// add half blocks in decorative places

	// ----------- Populate -------------------

	// ???
}

// ----------------World Generation----------------

func isRoomInChunk(chunk Chunk, start Point, room Room) bool {
	top, bottom := getRoomCoords(start, room)

	overlap := chunk.top.x > top.x ||
		chunk.top.y > top.y ||
		chunk.bottom.x < bottom.x ||
		chunk.bottom.y < bottom.y
	return !overlap
}

func buildRooms(world *World) {
	for i, chunk := range world.Chunks {
		room := roomTypes[rand.Intn(len(roomTypes))]
		fmt.Printf("Chunk %d has room type %d x %d \n", i, room.width, room.height)
		isRoomOutside := true
		startingPoint := Point{}

		for isRoomOutside {
			startingPoint = Point{
				x: chunk.top.x + rand.Intn(chunk.bottom.x-chunk.top.x),
				y: chunk.top.y + rand.Intn(chunk.bottom.y-chunk.top.y),
			}
			isRoomOutside = !isRoomInChunk(chunk, startingPoint, room)
		}
		fmt.Printf("Found a valid starting Point: (%d/%d)\n", startingPoint.x, startingPoint.y)
		world.insertRoom(startingPoint, room)
	}
}

func (world *World) insertRoom(start Point, room Room) {
	top, bottom := getRoomCoords(start, room)
	fmt.Printf("Drawing new room(%dx%d): (%d/%d) to (%d/%d)", room.width, room.height, top.x, top.y, bottom.x, bottom.y)

	// iterate over horizontal blocks
	for i := top.y; i <= bottom.y; i++ {
		// iterate over vertical blocks
		for j := top.x; j <= bottom.x; j++ {
			world.Data[i][j] = 1
		}
	}
}

func (room Room) size() int {
	return room.height * room.width
}

func avgSize(rooms []Room) float64 {
	sumSizes := 0
	for _, room := range rooms {
		sumSizes += room.size()
	}
	return float64(sumSizes) / float64(len(rooms))
}

func getRoomCoords(start Point, room Room) (Point, Point) {
	// try to center room around starting point
	// this will still be biased because width/height can be even
	xTop := start.x - room.width/2
	yTop := start.y - room.height/2
	xBottom := xTop + room.width - 1
	yBottom := yTop + room.height - 1
	return Point{x: xTop, y: yTop}, Point{x: xBottom, y: yBottom}
}

func chunkSize() (int, int) {
	height, width := worldSize, worldSize
	lastHeight, lastWidth := height, width
	currentChunkSize := height * width
	avgRoomSize := avgSize(roomTypes)
	minSize := int(avgRoomSize * (1 / fullnessFactor))

	for currentChunkSize > minSize {
		// save last working dimensions
		lastHeight = height
		lastWidth = width

		// split remaining space
		if height == width {
			//split horizontally
			height /= 2
		} else {
			//split vertically
			width /= 2
		}
		currentChunkSize = height * width
	}
	return lastHeight, lastWidth
}

func chunks(world *World, width, height int) []Chunk {
	numHorizontal := worldSize / width
	numVertical := worldSize / height
	chunksCount := numVertical * numHorizontal

	chunks := make([]Chunk, 0, chunksCount)
	// iterate over horizontal chunks
	for i := 0; i < numHorizontal; i++ {
		// iterate over vertical chunks
		for j := 0; j < numVertical; j++ {
			// find top and bottom Points
			// make a new chunk and push it to chunks
			startWidth := i * width
			startHeight := j * height
			top := Point{x: startWidth, y: startHeight}
			bottom := Point{x: startWidth + (width - 1), y: startHeight + (height - 1)}
			newChunk := Chunk{top: top, bottom: bottom}
			chunks = append(chunks, newChunk)
		}
	}
	return chunks
}
