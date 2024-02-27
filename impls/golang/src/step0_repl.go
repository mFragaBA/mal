package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
)

func READ(line string) string {
  return line
}

func EVAL(line string) string {
  return line
}

func PRINT(line string) string {
  return line
}

func rep(line string) string {
  return PRINT(EVAL(READ(line)))
}

func main() {
  fmt.Printf("user> ")
  reader := bufio.NewReader(os.Stdin)
  for {
    line, err := reader.ReadString('\n')
    if err == io.EOF {
      break
    }

    fmt.Println(rep(line))
    fmt.Printf("user> ")
  }
}
