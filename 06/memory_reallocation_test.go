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

func TestFindLargestBankTieBreaker(t *testing.T) {
  state := []int{3,1,2,3}
  expectedIndex := 0 // Ensure that tie-breaker works - should return the smallest-index bank
  actualIndex := findLargestBank(state)
  if expectedIndex != actualIndex {
      t.Errorf("expected %d, got %d", expectedIndex, actualIndex)
  }
}

func TestReallocateBlocks(t *testing.T) {
  state := []int{0,2,7,0}
  largestBankIndex := 2 //  The value at index 2 is the largest in the slice
  expectedNextState := []int{2,4,1,2}
  reallocateBlocks(largestBankIndex, state)

  for i, actual := range state {
    expected := expectedNextState[i]
    if expected != actual {
      t.Errorf("expected %d, got %d", expected, actual)
    }
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
