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
  actual := CaptchaFunc()
  for i, expected := range tests {
    if expected != actual {
      t.Errorf("at index %d, expected %d, got %d", i, expected, actual)
    }
  }
}



