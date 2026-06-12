package main


import (
  "fmt"
 // "os"
  "math/rand/v2"
//	tea "github.com/charmbracelet/bubbletea"
  )
  
  type scenario struct{
    text   string
    goldChange int
    govWatchChange int
    populationChange int
    happinessChange int
  }
  
  var gold int = 30
  var govWatch int = 300
  var population int = 100
  var happiness = 100
  var scenarios [4]scenario
  var choiceGlobal int = 0
  
  func chooseRandomScenarios(){
    var randomIndex = rand.IntN(4)
    fmt.Print(" ", scenarios[randomIndex].text)
    choiceGen()
    if choiceGlobal == 1{
      gold = gold + scenarios[randomIndex].goldChange
      govWatch = govWatch + scenarios[randomIndex].govWatchChange
      population = population+ scenarios[randomIndex].populationChange
      happiness = scenarios[randomIndex].happinessChange
  }
  GameLoop()
  }

func GameLoop(){
  fmt.Print("This is a resource management simulator right now you have", gold, " gold, ", govWatch, " government surveillance ", population, " population and ", happiness, " happiness")
    chooseRandomScenarios()
}

  func choiceGen() {
    fmt.Print("your choices are: Yes or No (input 1 or 2 to choose)")
    var choice string 
    fmt.Scanln(&choice)
    if choice == "1"{
      choiceGlobal = 1
    } else {
      choiceGlobal = 2
    }
  }

  func main(){
    scenarios[0] = scenario{
    text:        "The tax collector has come to collect taxes.",
    goldChange:  -23,
    govWatchChange: 300,
    populationChange: 0,
    happinessChange: 0,
  }
  scenarios[1] = scenario{
    text:        "There is a protest going on right now. Do you crush it?",
    goldChange:  -10,
    govWatchChange: 300,
    populationChange: -5,
    happinessChange: -30,
  }
  
  scenarios[2] = scenario{
    text:        "A new arcade is being planned in the capitol. Do you want to support it?",
    goldChange:  -10,
    govWatchChange: -50,
    populationChange: 0,
    happinessChange: 100,
  }
  
  scenarios[3] = scenario{
    text:        "A zombie apocalypse has broken out. Do you spend time to make sure it is destroyed?",
    goldChange:  -23,
    govWatchChange: 1000,
    populationChange: 0,
    happinessChange: 0,
  }
    GameLoop()
  }
  