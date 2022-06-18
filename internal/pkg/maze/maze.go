package maze

import "fmt"

type Maze struct {
	Walls Layer
	Foods Layer
	Characters Layer
}

func NewMaze ()(Maze, error) {
	var m Maze
	if err := m.Walls.Reset("Walls"); err != nil {
		return Maze{}, err
	}
	m.Foods.Reset("Foods")
	m.Characters.Reset("Characters")
	return m, nil
}