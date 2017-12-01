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
    fmt.Println(captcha(str))
  }
}

func captcha(input string) int {
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


