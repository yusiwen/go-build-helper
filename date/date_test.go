package date

import (
	"fmt"
	"testing"
)

func TestDate(t *testing.T) {
	o, _ := Date("")
	fmt.Println(o)
	o, _ = Date("2006.01.02")
	fmt.Println(o)
}
