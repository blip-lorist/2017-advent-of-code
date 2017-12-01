package main

import (
  "strings"
  "strconv"
  "bufio"
  "fmt"
  "os"
)

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    str := scanner.Text()
    fmt.Println(captchaHalf(str))
  }
}

func captchaNeighbor(input string) int {
  slice := strings.Split(input, "")
  length := len(slice)
  sum := 0

  for i, num := range slice {
    current, _ := strconv.Atoi(num)

    var next int

    if i == (length - 1) {
      next, _ = strconv.Atoi(slice[0])
    } else {
      next, _ = strconv.Atoi(slice[i + 1])
    }

    if current == next {
      sum += current
    }
  }

  return sum
}

func captchaHalf(input string) int {
  slice := strings.Split(input, "")
  length := len(slice)
  halfDistance := length/2
  sum := 0

  for i, num := range slice {
    current, _ := strconv.Atoi(num)

    var next int

    nextI := i + halfDistance
    if nextI > (length - 1) {
      next, _ = strconv.Atoi(slice[nextI-length])
    } else {
      next, _ = strconv.Atoi(slice[nextI])
    }

    if current == next {
      sum += current
    }
  }

  return sum
}
