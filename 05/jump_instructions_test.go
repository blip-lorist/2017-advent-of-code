package main

import (
  "testing"
  "bufio"
  "strings"
)

// Here's my test input and expected output
// This builds a literal hash map in Go
var testInput =  strings.NewReader("0\n3\n0\n1\n-3")
var testBuf = bufio.NewScanner(testInput)
var exitStepsCountTests = map[*bufio.Scanner]int{
  testBuf : 10, // test for part 2
}

func TestExitStepsCount(t *testing.T) {
  for input, expected := range exitStepsCountTests {
    actual := exitStepsCount(input)
    if expected != actual {
      t.Errorf("for input %s, expected %v, got %v", input, expected, actual)
    }
  }
}
