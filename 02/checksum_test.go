package main

import "testing"
// Learning how to test with: https://blog.alexellis.io/golang-writing-unit-tests/

// Here's my test input and expected output
// This builds a literal hash map in Go
var differenceTests = map[string]int{
  "5 1 9 5" : 8,
  "7 5 3" : 4,
  "2 4 6 8" : 6,
}

var evenDivisionTests = map[string]int{
  "5 9 2 8" : 4,
  "9 4 7 3" : 3,
  "3 8 6 5" : 2,
}

func TestDifference(t *testing.T) {
  for input, expected := range differenceTests {
    actual := difference(input)
    if expected != actual {
      t.Errorf("for input %s, expected %d, got %d", input, expected, actual)
    }
  }
}

func TestEvenDivision(t *testing.T) {
  for input, expected := range evenDivisionTests {
    actual := evenDivision(input)
    if expected != actual {
      t.Errorf("for input %s, expected %d, got %d", input, expected, actual)
    }
  }
}
