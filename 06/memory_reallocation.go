package main

import (
  "strings"
  "strconv"
)

func main() {
}

func reallocationCycles(input string) int {
  slice := strings.Fields(input)
  var intSlice []int
  for _,  str := range slice {
    integer, _ := strconv.Atoi(str)
    intSlice = append(intSlice, integer)
  }

  log := make(map[string]bool)
  hasUniqueState := true
  cycleCount := 0

  // Until basecase is found
  for hasUniqueState {
    cycleCount += 1

    // Log the bank state in a set and perform dup check
    hasUniqueState = logUniqueState(intSlice, log)

    // Find the largest bank
    //largestBankIndex := findLargestBank(intSlice)
    // Zero it out, then redistribute the blocks
    // reallocateBlocks(largestBankIndex)
  }
  return cycleCount
}

func reallocateBlocks(largestBankIndex int, intSlice []int) {
  blocksToDistribute := intSlice[largestBankIndex]
  intSlice[largestBankIndex] = 0
  maxIndex := len(intSlice) - 1
  var currentIndex int

  // Set current index to the one after the largest bank OR first
  if (largestBankIndex + 1) > maxIndex {
    currentIndex = 0
  } else {
    currentIndex = largestBankIndex + 1
  }

  for blocksToDistribute > 0 {
    intSlice[currentIndex] += 1
    blocksToDistribute -= 1


    if currentIndex == maxIndex {
      currentIndex = 0
    } else {
      currentIndex += 1
    }
  }
}

func findLargestBank(intSlice []int) int {
  largestSeenIndex := 0
  largestSeenValue := intSlice[0]

  for i, num := range intSlice {
    if num > largestSeenValue {
      largestSeenValue = num
      largestSeenIndex = i
    }
  }

  return largestSeenIndex
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
