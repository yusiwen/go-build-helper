package version

import (
	"fmt"
	"testing"
)

func TestGit(t *testing.T) {
	o, err := Version("d:\\git\\go-build-helper")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(o)

	o, err = Version(".")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(o)
}
