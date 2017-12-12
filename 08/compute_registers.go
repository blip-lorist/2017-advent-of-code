package main

import (
  "bufio"
  "strings"
  "github.com/apaxa-go/eval"
  "strconv"
  "fmt"
  "os"
)

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  fileBuf := bufio.NewScanner(file)
  fmt.Println(findLargestRegisterValue(fileBuf))
}

func findLargestRegisterValue(input *bufio.Scanner) int {
  registersLog := make(map[string]int)
  largestValueEver := 0

  for input.Scan() {
    line := input.Text()
    instructions := strings.Fields(line)
    latestLargestValue := processInstructionsLargestValue(registersLog, instructions)
    if latestLargestValue > largestValueEver {
      largestValueEver = latestLargestValue
    }
  }

  //var largestValueSeen int
  //for _,v := range registersLog {
  //  //if largestValueSeen == nil {
  //  //  largestValueSeen = v
  //  //}

  //  if v > largestValueSeen {
  //    largestValueSeen = v
  //  }
  //}

  return largestValueEver
}

func processInstructionsLargestValue(registersLog map[string]int, instructions []string) int {
  largestValueEver := 0
  subject := instructions[0]
  verb := instructions[1]
  alteration := instructions[2]
  leftOperand := instructions[4]
  operator := instructions[5]
  rightOperand := instructions[6]

  // If these registers aren't logged, than create them with a 0 starting value
  currentRegisters := []string{subject, leftOperand}
  for _, register := range currentRegisters {
    _, present := registersLog[register]
    if !present {
      registersLog[register] = 0
    }
  }

  // Check if evaluation is true
  leftOperandString := strconv.Itoa(registersLog[leftOperand])
  evalString :=  leftOperandString + " " + operator + " " + rightOperand
  expr, err := eval.ParseString(evalString,"")
  if err!=nil{
    fmt.Println(err)
  }
  evalResult, err := expr.EvalToInterface(nil)

  if err!=nil{
    fmt.Println(err)
  }

  if evalResult.(bool) {
    // Alter register
    subjectValue := registersLog[subject]
    alterationValue, _ := strconv.Atoi(alteration)

    if verb == "dec" {
      newValue := subjectValue - alterationValue
      if newValue > largestValueEver {
        largestValueEver = newValue
      }

      registersLog[subject] = newValue
    }
    if verb == "inc" {
      newValue := subjectValue + alterationValue
      if newValue > largestValueEver {
        largestValueEver = newValue
      }

      registersLog[subject] = newValue
    }
  }
  return largestValueEver
}
