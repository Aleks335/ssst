package strx

import "fmt"

func XA(strSlice []string) int {
	strNul := 0

	fmt.Println(strSlice)
	for _, s := range strSlice {
		if s == "B" {
			strNul++
		}
	}
	return strNul
}
