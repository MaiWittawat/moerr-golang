package main

import (
	"fmt"
	"moerr/moerr"
)

// how to use
func main() {
	err1 := moerr.New("something went wrong")
	err2 := moerr.NewErrorf("invalid value: %v", 42)

	fmt.Println(err1)
	fmt.Println(err2)
}
