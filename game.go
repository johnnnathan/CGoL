package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const (
    lines = 50
    cols  = 120
    halfLines = lines / 2 
    halfCols = cols / 2
    quarterLines = lines / 4
    quarterCols = cols / 4
)

var board [cols][lines]int
var boardTemp [cols][lines]int

func initBoard() {
    for line := 0; line < lines; line++ {
        for col := 0; col < cols; col++ {
            board[col][line] = rand.Intn(2)
        }
    }
}

func printBoard() string {
    output := "<table style='border-collapse: collapse;'>"
    for line := 0; line < lines; line++ {
        output += "<tr>"
        for col := 0; col < cols; col++ {
            if board[col][line] == 1 {
                output += "<td style='width: 15px; height: 15px; background-color: black;'></td>"
            } else {
                output += "<td style='width: 15px; height: 15px; background-color: white;'></td>"
            }
        }
        output += "</tr>"
    }
    output += "</table>"
    return output
}

func dynamicHandler(w http.ResponseWriter, r *http.Request) {
  //set content type
  w.Header().Add("Content-Type", "text/html")
  //set response code 
  w.WriteHeader(http.StatusOK)
  //print the board state to the writer
  fmt.Fprintln(w, printBoard())
}

func updateBoard() {
  var wg = sync.WaitGroup{}
  wg.Add(4)
  go updateQuad(&wg, 0, halfLines, 0, halfCols) 
  go updateQuad(&wg, 0, halfLines, halfCols, cols) 
  go updateQuad(&wg, halfLines, lines, 0, halfCols)
  go updateQuad(&wg, halfLines, lines, halfCols, cols) 

  wg.Wait()

  board = boardTemp
}

func updateQuad(wg *sync.WaitGroup,startLine int, endLine int, startCol int , endCol int){
  defer wg.Done()
    for line := startLine; line < endLine; line++ {
        for col := startCol; col < endCol; col++ {
            pop := getNeighbour(col, line)
            nextCycleState(col, line, pop, board[col][line])
        }
    }
}

func getNeighbour(col int, line int) int {
    var pop int
    for i := -1; i < 2; i++ {
        for j := -1; j < 2; j++ {
            if i == 0 && j == 0 {
                continue
            }
            if checkRange(col+i, line+j) && board[col+i][line+j] == 1 {
                pop++
            }
        }
    }
    return pop
}

func checkRange(col int, line int) bool {
    return col > -1 && col < cols && line > -1 && line < lines
}

func nextCycleState(col int, line int, neighbourCount int, alive int) {
    if alive == 1 {
        if neighbourCount > 3 || neighbourCount < 2 {
            boardTemp[col][line] = 0
        } else {
            boardTemp[col][line] = 1
        }
    } else {
        if neighbourCount == 3 {
            boardTemp[col][line] = 1
        } else {
            boardTemp[col][line] = 0
        }
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    initBoard()

    //set htmx file directory
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

    //operate on tag in htmx file
    http.HandleFunc("/dynamic", func(w http.ResponseWriter, r *http.Request) {
        updateBoard()
        dynamicHandler(w, r)
    })

    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
