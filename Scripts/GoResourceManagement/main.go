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
  const ScenarioNum = 9
  var gold int = 30
  var govWatch int = 300
  var population int = 100
  var happiness = 100
  var scenarios [ScenarioNum]scenario
  var choiceGlobal int = 0
  
  func chooseRandomScenarios(){
    var randomIndex = rand.IntN(ScenarioNum)
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
  if gold < -30{
    fmt.Print("\nThe debt collectors have come to your establishment. They demand payment GAME OVER")
    return
  }
  if govWatch > 1000{
    fmt.Print("\nYou created a surveillance state congratulations. GAME OVER")
    return
  }
  if happiness < -300{
    fmt.Print("\nthe people have decided to overthrow you because you are not making them happy GAME OVER")
    return
  }
  if population <= 0{
    fmt.Print("You don't have anymore population GAME OVER")
    return
  }
  if  govWatch <= 20{
    fmt.Print("The nation has fallen into anarchy GAME OVER")
    return
  }
  fmt.Print("This is a resource management simulator right now you have " , gold, " gold, ", govWatch, " government surveillance ", population, " population and ", happiness, " happiness")
    chooseRandomScenarios()
}

  func choiceGen() {
    fmt.Print("\nyour choices are: Yes or No (input 1 or 2 to choose)")
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
    text:        "\nThe tax collector has come to collect taxes.",
    goldChange:  -23,
    govWatchChange: 300,
    populationChange: 0,
    happinessChange: 0,
  }
  scenarios[1] = scenario{
    text:        "\nThere is a protest going on right now. Do you crush it?",
    goldChange:  -10,
    govWatchChange: 300,
    populationChange: -5,
    happinessChange: -30,
  }
  
  scenarios[2] = scenario{
    text:        "\nA new arcade is being planned in the capitol. Do you want to support it?",
    goldChange:  -10,
    govWatchChange: -50,
    populationChange: 0,
    happinessChange: 100,
  }
  
  scenarios[3] = scenario{
    text:        "\nA zombie apocalypse has broken out. Do you spend time to make sure it is destroyed?",
    goldChange:  -23,
    govWatchChange: 1000,
    populationChange: 0,
    happinessChange: 0,
  }
  scenarios[4] = scenario{
    text:        "\nA neighboring country has fallen into civil war. Do you intervene?",
    goldChange:  -30,
    govWatchChange: -10,
    populationChange: -10,
    happinessChange: 0,
  }
  scenarios[5] = scenario{
    text:        "\n The people give their gold to you.",
    goldChange:  300,
    govWatchChange: 0,
    populationChange: 0,
    happinessChange: 0,
  }
  scenarios[6] = scenario{
    text:        "\n The people give their gold to you.",
    goldChange:  300,
    govWatchChange: 0,
    populationChange: 0,
    happinessChange: 0,
  }
  scenarios[7] = scenario{
    text:        "\n Do you approve a party on the road next to the capital",
    goldChange:  -10,
    govWatchChange: 0,
    populationChange: 10,
    happinessChange: 200,
  }
    scenarios[8] = scenario{
    text:        "\n A group if refugees is outside your border. Do you let them in?",
    goldChange:  -10,
    govWatchChange: 0,
    populationChange: 100,
    happinessChange: -5,
  }
    GameLoop()
}  
  