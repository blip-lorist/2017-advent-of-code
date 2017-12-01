package main

import "testing"
// Learning how to test with: https://blog.alexellis.io/golang-writing-unit-tests/

// Here's my test input and expected output
// This builds a literal hash map in Go
var partOneTests = map[string]int{
  "1122" : 3,
  "1111" : 4,
  "1234" : 0,
  "91212129" : 9,
}

var partTwoTests = map[string]int{
  "1212" : 6,
  "1221" : 0,
  "123425" : 4,
  "123123" : 12,
  "12131415" : 4,
}

func TestCaptchaNeighbor(t *testing.T) {
  for input, expected := range partOneTests {
    actual := captchaNeighbor(input)
    if expected != actual {
      t.Errorf("for input %s, expected %d, got %d", input, expected, actual)
    }
  }
}

func TestCaptchaHalf(t *testing.T) {
  for input, expected := range partTwoTests {
    actual := captchaHalf(input)
    if expected != actual {
      t.Errorf("for input %s, expected %d, got %d", input, expected, actual)
    }
  }
}




