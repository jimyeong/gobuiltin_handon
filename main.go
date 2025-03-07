package main

import (
	"fmt"
	"strings"

	"example.com/m/v2/hand_on/helper"
)

func main() {
	k := "hellllllllllo worlllllld"
	v := helper.Delete_empty_strings(strings.Split(k, ""))
	fmt.Println(v)
	fmt.Println("size is %s", len(v))

}
