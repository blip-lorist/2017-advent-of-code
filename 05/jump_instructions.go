package main

import (
  "bufio"
  "fmt"
  "os"
  "container/list"
  "strings"
  "strconv"
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
    exitStepsCount(line)
  }
  fmt.Println(sum)
}

func exitStepsCount(input string) int {
  slice := strings.Fields(input)

  stepsLinkedList := list.New()

  var latestNode *list.Element

  // Build linked list
  for _, current := range slice {
    currentInt, _ := strconv.Atoi(current)

    if stepsLinkedList.Len() == 0 {
      latestNode = stepsLinkedList.PushBack(currentInt)
    } else {
      latestNode = stepsLinkedList.InsertAfter(currentInt, latestNode)
    }
  }

  // Node we're leaving
  previous := stepsLinkedList.Front()
  current := stepsLinkedList.Front()

  fmt.Println("first node value")
  fmt.Println(previous.Value)

  // Follow instructions
  instructions, _ := previous.Value.(int)

  for i := 1; i <= instructions; i++ {
    current = current.Next()
  }

  fmt.Println("current node value")
  fmt.Println(current.Value)
  // Increment previous node
  previous.Value = instructions + 1

  fmt.Println("first node after leaving")
  fmt.Println(previous.Value)

  return 0

}
