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

  var sortedLettersSlice[]string

  // Sorted letters
  for _, current := range slice {
    lettersSlice := strings.Split(current, "")
    sort.Strings(lettersSlice)
    sortedLettersString := strings.Join(lettersSlice, "")
    sortedLettersSlice = append(sortedLettersSlice, sortedLettersString)
  }

  // Sorted words
  sort.Strings(sortedLettersSlice)

  for i, current := range sortedLettersSlice {
    // De-dup: This also finds anagrams next to each other
    if i < (len(sortedLettersSlice)-1) && current == sortedLettersSlice[i+1] {
      return false
    }
  }
  return true
}
