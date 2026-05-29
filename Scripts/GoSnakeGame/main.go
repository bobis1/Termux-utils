package main

import (
	"fmt"
	"os"
  "math/rand/v2"
	tea "github.com/charmbracelet/bubbletea"
)


var rows int = 5
var collums int = 16
var score = 0
var maxRow = 4
var maxCollums = 15
var FoodR int = 3
var FoodC int = 9

type position struct {
	x int
	y int
}

type model struct {
  body []position
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
		if m.body[0].y > 0{
var newPos = position {
		  x: m.body[0].x,
		  y: m.body[0].y - 1,
		  }
		  m.body = append([]position{newPos}, m.body...)
		} else{
	var newPos = position {
		  x: m.body[0].x,
		  y: 0,
		  }
		  m.body = append([]position{newPos}, m.body...)
		}
		case "down":
		if m.body[0].y < maxRow{
    var newPos = position {
		  x: m.body[0].x,
		  y: m.body[0].y + 1,
		  }
		  m.body = append([]position{newPos}, m.body...)
		} else{
      var newPos = position {
		  x: m.body[0].x,
		  y: maxRow,
		  }
		  m.body = append([]position{newPos}, m.body...)
		}
		case "left":
		if m.body[0].x > 0{
      var newPos = position {
		  x: m.body[0].x - 1,
		  y: m.body[0].y,
		  }
		  m.body = append([]position{newPos}, m.body...)
		} else {
		  var newPos = position {
		  x: m.body[0].x,
		  y: 0,
		  }
		  m.body = append([]position{newPos}, m.body...)
		}
		case "right":
		if m.body[0].x < maxCollums{
		  var newPos = position {
		  x: m.body[0].x + 1,
		  y: m.body[0].y,
		  }
		  m.body = append([]position{newPos}, m.body...)
		} else {
		  var newPos = position{
		  x: maxCollums,
		  y: m.body[0].y,
		}
		m.body = append([]position{newPos}, m.body...)
		
		}
	}
	if FoodC == m.body[0].x && FoodR == m.body[0].y {
				  score++
				  FoodC = rand.IntN(maxCollums)
				  FoodR = rand.IntN(maxRow)
	
	} else {
	  m.body = m.body[:len(m.body)-1]
	}
	}
	return m, nil
}

func (m model) View() string {
	var board string

	for y := 0; y < rows; y++ {
		for x := 0; x < collums; x++ {
			

			isSnakeHere := false
			
			
			for i := 0; i < len(m.body); i++ {
				if m.body[i].x == x && m.body[i].y == y {
					isSnakeHere = true
				}
			}


			if isSnakeHere {
				board += "█"
			} else if FoodC == x && FoodR == y {
				board += "!"
			} else {
				board += "."
			}
			
		}
		board += "\n"
	}

	board += fmt.Sprintf("X: %d, Y: %d  Score: %d | Press 'q' to quit", m.body[0].x, m.body[0].y, score)
	return board
}


func eaten(x int, y int, m model){
    if m.body[0].x == x && m.body[0].y == y{
      score++
    }
}

func main() {
p := tea.NewProgram(model{
		body: []position{
			{x: 7, y: 2},
		},
	})

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}



