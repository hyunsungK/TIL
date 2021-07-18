package main

import (
	"fmt"
	"strconv"
)

func main() {
	if i, err := ConvertBinary2Int("10000"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
}

func ConvertBinary2Int(binary string) (int64, error) {
	return strconv.ParseInt(binary, 2, 64)
}
