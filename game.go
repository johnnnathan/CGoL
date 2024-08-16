package main

import (
	"fmt"
	"math/rand"
	"time"
)

const lines int = 40 
const cols int = 30
var board [cols][lines]int
var boardTemp [cols][lines]int 

func initBoard()  {
  for line:=0; line < lines; line ++ {
    for col:=0; col < cols; col++ {
      roll := rand.Intn(2) 
      board[col][line] = roll
    }
  }
}
func printBoard() {
    // Clear the screen (ANSI escape codes for terminal)
    fmt.Print("\033[H\033[2J")

    // Iterate over each row and column of the board
    for line := 0; line < lines; line++ {
        for col := 0; col < cols; col++ {
            if board[col][line] == 1 {
                fmt.Print("* ") // Print '*' for live cells
            } else {
                fmt.Print("  ") // Print space for dead cells
            }
        }
        fmt.Println() // Move to the next line after printing all columns in the current row
    }
    fmt.Println() // Extra line for better separation between updates
}
func getNeighbour(col int, line int)int{
  var pop int 
  for i:=-1; i < 2; i++{
    for j:=-1; j < 2; j++{
      if i == 0 && j ==0 {
        continue
      }
      pop += getPosValue(col,line,i,j)
    }
  }


  return pop

}

func getPosValue(col int, line int, dirHor int, dirVer int)int{
  flag := checkRange(col + dirHor, line + dirVer)
  if flag && board[col + dirHor][line + dirVer] == 1{
    return 1
  }
  return 0 
}

func nextCycleState(col int, line int, neighbourCount int, alive int)  {
  var aliveBool bool
  if alive == 1{
    aliveBool = true
  }else {
    aliveBool = false
  }
    switch {
    case aliveBool && (neighbourCount > 3 || neighbourCount < 2):
      boardTemp[col][line] = 0
      return
    case !aliveBool && (neighbourCount == 3):
      boardTemp[col][line] = 1
      return
    default:
      boardTemp[col][line] = board[col][line]
  }
  return
}


func update() {
  for line:=0; line < lines; line ++ {
    for col:=0; col < cols; col++ {
      pop := getNeighbour(col, line)
      nextCycleState(col, line, pop, board[col][line])
    }
  }
  board = boardTemp

}

func checkRange(col int, line int)bool  {
  if col >-1 && col < cols && line >-1 && line < lines{
    return true
  }
  return false

  
}

func main()  {
  initBoard()
  printBoard()
  for true{
    update()
    printBoard()
    time.Sleep(200* time.Millisecond)
  }
}



