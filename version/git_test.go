package version

import (
	"fmt"
	"testing"
)

func TestGit(t *testing.T) {
	o, err := Version("d:\\git\\EasyDarwin")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(o)
}
