package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  fileBuf := bufio.NewScanner(file)
  fmt.Println(findRoot(fileBuf))
}

func findRoot(input *bufio.Scanner) string {
  log := make(map[string]string)
  root := ""
  var programs []string

  for input.Scan() {
    line := input.Text()
    var branches []string
    lineSlice := strings.Fields(line)

    program := lineSlice[0]
    if len(lineSlice) > 2 && lineSlice[2] == "->" {
      branches = lineSlice[3:len(lineSlice)]
    }

    logBranchRoot(branches, program, log)

    programs = append(programs, program)
  }

  for _, program := range programs {
    _, present := log[program]
    if !present {
      root = program
    }
  }
  return root
}

func logBranchRoot(branches []string, program string, log map[string]string) {
  for _, branch := range branches {
    if strings.Contains(branch, ",") {
      branch = branch[:len(branch)-1]
    }
    log[branch] = program
  }
}
