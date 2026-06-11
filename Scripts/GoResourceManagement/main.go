package main


import (
  "fmt"
 // "os"
//  "math/rand/v2"
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
  var scenarios [1]scenario
  var choice int = 0
  
  func chooseRandomScenarios(){
    var randomIndex = rand.IntN(10)
    fmt.Print(" ", scenarios[0].text)
    choiceGen()
    if choice == 1{
      gold = gold + scenarios[randomIndex].goldChange
      govWatch = govWatch + scenarios[randomIndex].govWatchChange
      population = population+ scenarios[randomIndex].populationChange
      happiness = scenarios[randomIndex].happinessChange
    } else {
      
    }
  }

  func choiceGen() int {
    fmt.Print("your choices are: Yes or No (input 1 or 2 to choose)")
    var choice string 
    fmt.Scanln(&choice).
    if choice == "1"{
      choice = 1
    } else {
      return
      choice = 2
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
  
    fmt.Print("This is a resource management simulator right now you have", gold, " gold, ", govWatch, " government surveillance ", population, " population and ", happiness, " happiness")
    chooseRandomScenarios()
  }
  