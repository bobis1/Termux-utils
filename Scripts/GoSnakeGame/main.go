package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)


var rows int = 5
var collums int = 16


type model struct {
	x int
	y int
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "up":
		if m.y > 0{
		  m.y--
		} else{
		  m.y=0
		}
		case "down":
		if m.y < rows{
		  m.y++
		} else{
		  m.y = rows
		}
		case "left":
		if m.x > 0{
			m.x--
		} else {
		  m.x = 0
		}
		case "right":
		if m.x < collums{
			m.x++
		} else {
		  m.x = collums
		}
	}
	}
	return m, nil
}

func (m model) View() string {
	var board string

	
	for y := 0; y < rows; y++ {
		for x := 0; x < collums; x++ {
			if m.x == x && m.y == y {
				board += "█"
			} else {
				board += "."
			}
		}
		board += "\n"
	}

	board += fmt.Sprintf("X: %d, Y: %d | Press 'q' to quit", m.x, m.y)
	return board
}

func main() {
	p := tea.NewProgram(model{x: 7, y: 2})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}



