package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("GoRogue")
	os.Chdir("path/does/not/exist")
}
