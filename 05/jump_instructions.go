package main

import (
  "bufio"
  "fmt"
  "os"
  "container/list"
  "strconv"
  "math"
)

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  fileBuf := bufio.NewScanner(file)
  fmt.Println("Sum is")
  fmt.Println(exitStepsCount(fileBuf))
}

func exitStepsCount(input *bufio.Scanner) int {

  stepsLinkedList := list.New()

  var latestNode *list.Element

  // Build linked list
  for input.Scan() {
    current := input.Text()
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
    previous = current
    current = moveToNewNode(previous, current)

    // Increment step
    stepSum += 1

    if current == nil {
      break
    }

    // update previous node
    previous.Value = updatedPrevValue(previous)
    //fmt.Printf("currentValue is %d \n", currentValue)
    //fmt.Printf("stepsum is %d \n", stepSum)
  }

  return stepSum
}

func moveToNewNode(previous *list.Element, current *list.Element) *list.Element {
 // Follow instructions

  instructions, _ := previous.Value.(int)

  if instructions < 0 {
    absValueInstructions := math.Abs(float64(instructions))

    for i := 1; i <= int(absValueInstructions); i++ {
      current = current.Prev()
      if current == nil {
        break
      }
    }
  } else {
    for i := 1; i <= instructions; i++ {
      current = current.Next()
      if current == nil {
        break
      }
    }
  }
  return current
}

func updatedPrevValue(previous *list.Element) int {
  var newPreviousValue int
  if previous.Value.(int) >= 3 {
    newPreviousValue = previous.Value.(int) - 1
  } else {
    newPreviousValue = previous.Value.(int) + 1
  }

  return newPreviousValue
}
