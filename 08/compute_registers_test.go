package main

import (
  "testing"
  "bufio"
  "strings"
)

var testInput =
  "b inc 5 if a > 1\n" +
  "a inc 1 if b < 5\n" +
  "c dec -10 if a >= 1\n" +
  "c inc -20 if c == 10\n"

var inputReader = strings.NewReader(testInput)
var testBuf = bufio.NewScanner(inputReader)

var findLargestRegisterValueTest = map[*bufio.Scanner]int {
  testBuf : 10,
}

func TestFindLargestRegisterValue (t *testing.T) {
  for _, expected := range findLargestRegisterValueTest {
    actual := findLargestRegisterValue(testBuf)
    if expected != actual {
      t.Errorf("Expected %d, got %d", expected, actual)
    }
  }
}
