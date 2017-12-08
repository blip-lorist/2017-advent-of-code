package main

import (
  "testing"
)

func TestLogState(t *testing.T) {
  stateToLog := []int{0,2,7,0}
  expectedKey := "0270"
  log := make(map[string]bool)
  logUniqueState(stateToLog, log)
  _, present := log[expectedKey]
  if !present {
      t.Errorf("expectedKey %s is not present", expectedKey)
  }

  // Test duplicate detection
  isUnique := logUniqueState(stateToLog, log)
  if isUnique {
    t.Errorf("duplicate key %s was not detected", expectedKey)
  }

}

func TestFindLargestBank(t *testing.T) {
  state := []int{0,2,7,0}
  expectedIndex := 2 //  The value at index 2 is the largest in the slice
  actualIndex := findLargestBank(state)
  if expectedIndex != actualIndex {
      t.Errorf("expected %d, got %d", expectedIndex, actualIndex)
  }
}

var memoryReallocationTest = map[string]int {
  "0 2 7 0" : 5,
}

func TestMemoryReallocation(t *testing.T) {
  for input, expected := range memoryReallocationTest{
    actual := reallocationCycles(input)
    if expected != actual {
      t.Errorf("for input %s, expected %v, got %v", input, expected, actual)
    }
  }
}
