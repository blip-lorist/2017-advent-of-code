package main

import "testing"

// Here's my test input and expected output
// This builds a literal hash map in Go
var validPassphrasesTests = map[string]bool{
  "abcde fghij" : true,
  "abcde xyz ecdab" : false,
  "a ab abc abd abf abj" : true,
  "iiii oiii ooii oooi oooo" : true,
  "oiii ioii iioi iiio" : false,
}

func TestIsValidPassphrase(t *testing.T) {
  for input, expected := range validPassphrasesTests{
    actual := isValidPassphrase(input)
    if expected != actual {
      t.Errorf("for input %s, expected %v, got %v", input, expected, actual)
    }
  }
}
