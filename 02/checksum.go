package main

import (
  "strings"
  "strconv"
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
    diff := difference(line)
    sum += diff
  }
  fmt.Println(sum)
}

func difference(input string) int {
  slice := strings.Fields(input)
  firstValue := slice[0]

  // Largest and smallest values are initially 
  // the first value
  largest, _ := strconv.Atoi(firstValue)
  smallest, _ := strconv.Atoi(firstValue)

  for _, num := range slice {
    current, _ := strconv.Atoi(num)

    if current > largest {
      largest = current
    }

    if current < smallest {
      smallest = current
    }
  }

  diff := largest - smallest
  return diff
}

func evenDivision(input string) int {
  // There are a few ways to solve this, I think
  // Prime factor comparisons
  // Even vs odd
  // On*n division

  slice := strings.Fields(input)
  var intSlice[]int

  // Convert all strings to ints
  for _, num := range slice {
    num, _ := strconv.Atoi(num)
    intSlice = append(intSlice, num)
  }

  // Sort from smallest to largest
  sort.Ints(intSlice)

  for i, num := range intSlice {
    for j, largerNum := range intSlice {
      if i == j {
        // don't divide a number by itself
        continue
      }

      // Divide each num into all other numbers
      if largerNum % num == 0 {
        quotient := largerNum / num
        return quotient
      }
    }
  }

  err := -1
  return err
}
