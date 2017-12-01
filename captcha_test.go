package captcha

import "testing"
// Learning how to test with: https://www.binpress.com/tutorial/getting-started-with-go-and-test-driven-development/160

// Here's my test input and expected output
// This builds a literal hash map in Go
var tests = map[int]int{
  1122 : 3,
  1111 : 4,
  1234 : 0,
  91212129 : 9,
}

func TestCaptchaFunc(t *testing.T) {
  fn := CaptchaFunc()
  for i, v := range tests {
    if val := fn(); val != v {
      t.Fatalf("at index %d, expected %d, got %d", i, v, val)
    }
  }
}



