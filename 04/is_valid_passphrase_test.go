package main

import "testing"

// Here's my test input and expected output
// This builds a literal hash map in Go
var validPassphrasesTests = map[string]bool{
  "aa bb cc dd ee" : true,
  "aa bb cc dd aa" : false,
  "aa bb cc dd aaa" : true,
}

func TestIsValidPassphrase(t *testing.T) {
  for input, expected := range validPassphrasesTests{
    actual := isValidPassphrase(input)
    if expected != actual {
      t.Errorf("for input %s, expected %v, got %v", input, expected, actual)
    }
  }
}

