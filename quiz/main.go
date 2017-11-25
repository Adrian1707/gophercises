package main

import (
  "fmt"
  "encoding/csv"
  "os"
  "bufio"
  "io"
  "time"
)


func countLines() int64{
  f, _ := os.Open("problems.csv")

  s := bufio.NewScanner(f)
  lc := int64(0)
  for s.Scan(){
    lc++
  }
  return lc
}


func main(){
  timer1 := time.NewTimer(time.Second * 8)
  f, _ := os.Open("problems.csv")

  lines := countLines()
  r := csv.NewReader(bufio.NewReader(f))
  correct := 0
  questions := 0

  for {
    go func(){
      <-timer1.C
      fmt.Println("Times up! You got ", correct, "out of", questions, "answered", "and", lines, "questions in total")
      os.Exit(0)
    }()
    record, err := r.Read()
    if err == io.EOF {
      break
    }
    fmt.Println("Hey User, what is", record[0])
    fmt.Print("Enter text: ")
    var input string
    fmt.Scanln(&input)
    if record[1] == input {
      correct += 1
    }
    questions += 1
  }
   fmt.Println("Congratulations, you got ", correct, "out of", questions)
 }

