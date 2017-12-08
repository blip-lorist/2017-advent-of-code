package main

import (
  "strings"
  "strconv"
  "os"
  "fmt"
  "bufio"
)

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    fmt.Println(reallocationLoopCycles(line))
  }
}

func reallocationLoopCycles(input string) int {
  slice := strings.Fields(input)
  var intSlice []int
  for _,  str := range slice {
    integer, _ := strconv.Atoi(str)
    intSlice = append(intSlice, integer)
  }

  log := make(map[string]int)
  cycleCount := 0
  logUniqueState(intSlice, cycleCount, log) // Log the starting state
  hasUniqueState := true

  // Until basecase is found
  for hasUniqueState {
    // Find the largest bank
    largestBankIndex := findLargestBank(intSlice)
    // Zero it out, then redistribute the blocks
    reallocateBlocks(largestBankIndex, intSlice)

    cycleCount += 1

    // Log the bank state in a set and perform dup check
    hasUniqueState = logUniqueState(intSlice, cycleCount, log)
  }

  loopStateKey := getLogKey(intSlice)
  cycleFirstSeen := log[loopStateKey]
  loopCycles := cycleCount - cycleFirstSeen
  return loopCycles
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

  // Iterate backwards through the slice
  // Since ties are broken by the lowest index bank
  // This allows us to break the tie during the first pass, instead of searching and tiebreaking afterwards
  // Bug fixed: largestSeenValue needs to be updated when the current element is >= to largestSeenValue, not just >
  for i := len(intSlice) - 1; i >= 0; i-- {
    if intSlice[i] >= largestSeenValue {
      largestSeenValue = intSlice[i]
      largestSeenIndex = i
    }
  }

  return largestSeenIndex
}

func logUniqueState(currentBanks []int, cycleCount int, log map[string]int) bool {
  // If logging is successful, return true, if duplicate is found return false

  currentBanksString := getLogKey(currentBanks)

  _, present := log[currentBanksString]
  if !present {
    log[currentBanksString] = cycleCount
    return true
  } else {
    return false
  }
}

// Helper method to generate a log key
// eg: "4,2,5,1"
func getLogKey(banksSlice []int) string {
  currentBanksString := ""
  var delimiter string
  for i,  bank := range banksSlice {
    // It's important to a delimiter in this key
    // Otherwise I will get false positives!

    if i == len(banksSlice) - 1 {
      delimiter = ""
    } else {
      delimiter = ","
    }

    currentBanksString += (strconv.Itoa(bank) + delimiter)
  }

  return currentBanksString
}
