# go-generator
A level generator that builds connected rooms written in (bad) go

### How it works

Based on settings for world size and fullness, creates an empty world grid.
Content is then created in several steps:

1. split world into chunks
1. create a room in each  (room dimensions are picked randomly from a list of room types)
1. use some maze algorithm to define which rooms should be connected
1. build bridges that connect the rooms 
1. ???


### Configuration

##### Settings available:
world size: dimensions of the world grid
fullness_factor: defines how much empty space the world should contain
roomTypes: a list of tuples defining the width/height combinations for rooms in this world

##### Potential other settings:
seed: for random number generator (atm the seed is not changed, meaning the result of random number generation is deterministic)
[See Docs](https://golang.org/pkg/math/rand/)
