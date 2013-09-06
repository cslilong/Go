// sqrt_test.go
package simplemath

import "testing"

import "fmt"

func TestSqrt1(t *testing.T) {
	v := Sqrt(16)
	fmt.Println(v)
	if v != 4 {
		t.Errorf("Sqrt(16) failed. Got %v, expected 4.", v)
	}
}
