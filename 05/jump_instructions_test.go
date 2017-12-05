package main

import "testing"

// Here's my test input and expected output
// This builds a literal hash map in Go
var exitStepsCountTests = map[string]int{
  "0\n3\n0\n1\n-3" : 5,
}

func TestExitStepsCount(t *testing.T) {
  for input, expected := range exitStepsCountTests {
    actual := exitStepsCount(input)
    if expected != actual {
      t.Errorf("for input %s, expected %v, got %v", input, expected, actual)
    }
  }
}
