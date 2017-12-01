package captcha

import "testing"
// Learning how to test with: https://blog.alexellis.io/golang-writing-unit-tests/

// Here's my test input and expected output
// This builds a literal hash map in Go
var tests = map[int]int{
  1122 : 3,
  1111 : 4,
  1234 : 0,
  91212129 : 9,
}

func TestCaptchaFunc(t *testing.T) {
  for input, expected := range tests {
    actual := CaptchaFunc(input)
    if expected != actual {
      t.Errorf("for input %d, expected %d, got %d", input, expected, actual)
    }
  }
}



