package main

import (
  "testing"
  "bufio"
  "strings"
)

var testInput =
  "pbga (66)\n" +
  "xhth (57)\n" +
  "ebii (61)\n" +
  "havc (66)\n" +
  "ktlj (57)\n" +
  "fwft (72) -> ktlj, cntj, xhth\n" +
  "qoyq (66)\n" +
  "padx (45) -> pbga, havc, qoyq\n" +
  "tknk (41) -> ugml, padx, fwft\n" +
  "jptl (61)\n" +
  "ugml (68) -> gyxo, ebii, jptl\n" +
  "gyxo (61)\n" +
  "cntj (57)\n"

var inputReader = strings.NewReader(testInput)
var testBuf = bufio.NewScanner(inputReader)

var findRootTest = map[*bufio.Scanner]string {
  testBuf : "tknk",
}

func TestFindRoot(t *testing.T) {
  for _, expected := range findRootTest {
    actual := findRoot(testBuf)
    if expected != actual {
      t.Errorf("Expected %s, got %s", expected, actual)
    }
  }
}

func TestLogBranchesRoot(t *testing.T) {
  log := make(map[string]string)
  branches := []string{
    "ktlj",
    "cntj",
  }

  program := "fwft"

  logBranchRoot(branches, program, log)
  actual := log[branches[0]]
  if actual != program {
    t.Errorf("Expected %s, got %s", program, actual) 
  }
}


