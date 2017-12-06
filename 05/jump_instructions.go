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
      // Add first node
      latestNode = stepsLinkedList.PushBack(currentInt)
    } else {
      // Add new node and update latestNode
      latestNode = stepsLinkedList.InsertAfter(currentInt, latestNode)
    }
  }

  // Node we're leaving
  previous := stepsLinkedList.Front()
  current := stepsLinkedList.Front()
  stepSum := 0


  // wtf go why is this your while loop
  for current != nil {
    current = moveToNewNode(previous, current)

    if current == nil {
      fmt.Println("current node is nil")
      break
    } else {
      fmt.Println("current node")
      fmt.Println(current.Value)
    }

    // Increment step
    stepSum += 1

    // Increment previous node
    previous.Value = previous.Value.(int) + 1

    fmt.Println("updated previous to:")
    fmt.Println(previous.Value)
  }

  return stepSum
}

func moveToNewNode(previous *list.Element, current *list.Element) *list.Element {
 // Follow instructions
  instructions, _ := previous.Value.(int)

  for i := 1; i <= instructions; i++ {
    current = current.Next()
    if current == nil {
      break
    }
  }
  return current
}

