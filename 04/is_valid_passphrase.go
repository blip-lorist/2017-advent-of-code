package main

import (
  "strings"
  "bufio"
  "fmt"
  "os"
  "sort"
)

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  sum := 0
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    // This is how to predicate method in Go
    if isValidPassphrase(line) {
      sum += 1
    }
  }
  fmt.Println(sum)
}

func isValidPassphrase(input string) bool {
  slice := strings.Fields(input)

  // Sort
  sort.Strings(slice)

  for i, current := range slice {
    if i < (len(slice)-1) && current == slice[i+1] {
      return false
    }
  }
  return true
}
