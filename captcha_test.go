package captcha

import "testing"
// Learning how to test with: https://blog.alexellis.io/golang-writing-unit-tests/

// Here's my test input and expected output
// This builds a literal hash map in Go
var tests = map[string]int{
  "1122" : 3,
  "1111" : 4,
  "1234" : 0,
  "91212129" : 9,
}

func TestCaptcha(t *testing.T) {
  for input, expected := range tests {
    actual := captcha(input)
    if expected != actual {
      t.Errorf("for input %s, expected %d, got %d", input, expected, actual)
    }
  }
}



