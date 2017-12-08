package main

import (
  //"strings"
  "strconv"
)

func main() {
}

func reallocationCycles(input string) int {
  //slice := strings.Fields(input)
  //log := make(map[string]bool)

  // Until basecase is found
  // Log the bank state in a set and perform dup check
  //logState(currentBanks, log)
  // Find the largest bank
  // Zero it out, then redistribute the blocks
  return 0
}

func logUniqueState(currentBanks []int, log map[string]bool) bool {
  // If logging is successful, return true, if duplicate is found return false

  currentBanksString := ""
  for _,  bank := range currentBanks {
    currentBanksString += strconv.Itoa(bank)
  }

  _, present := log[currentBanksString]
  if !present {
    log[currentBanksString] = true
    return true
  } else {
    return false
  }
}
