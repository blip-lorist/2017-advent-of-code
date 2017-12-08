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

//var memoryReallocationTest = map[string]int {
//  "0 2 7 0" : 5,
//}
//
//func TestMemoryReallocation(t *testing.T) {
//  for input, expected := range memoryReallocationTest{
//    actual := reallocationCycles(input)
//    if expected != actual {
//      t.Errorf("for input %s, expected %v, got %v", input, expected, actual)
//    }
//  }
//}
