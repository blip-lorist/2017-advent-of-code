package main

import (
  //"strings"
  "strconv"
)

func main() {
}

func reallocationCycles(input string) int {
  //slice := strings.Fields(input)
  //var log map[string]bool

  // Until basecase is found
  // Log the bank state in a set and perform dup check
  //logState(currentBanks, log)
  // Find the largest bank
  // Zero it out, then redistribute the blocks
  return 0
}

func logState(currentBanks []int, log map[string]bool) map[string]bool {
  currentBanksString := ""
  for _,  bank := range currentBanks {
    currentBanksString += strconv.Itoa(bank)
  }

  log[currentBanksString] = true
  return log
}
