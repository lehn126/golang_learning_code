package p3

import (
	"fmt"
	"os"
)

func TestEnv() {
	fmt.Printf("env GOROOT: %v\n", os.Getenv("GOROOT"))
}
