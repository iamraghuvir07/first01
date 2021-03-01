package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []string{"abc", "bcvd", "qwe", "oiu", "qaz", "xcv", "mnb"}
	var choice string
	fmt.Scanf("%s", &choice)
	getSort(arr, choice)
}
func getSort(brr []string, c string) {
	sort.Strings(brr)
	x := len(brr)
	var b int
	b = x
	x--
	if c == "desc" {
		var crr = make([]string, b)
		count := b

		fmt.Printf("Strings in descending order\n")
		for i := 0; i < b; i++ {
			count--
			crr[count] = brr[i]

		}
		for i := 0; i <= x; i++ {
			//fmt.Printf("Strings in descending order")
			fmt.Printf("%s\n", crr[i])
		}

	} else {
		fmt.Printf("Strings in ascending order\n")
		for i := 0; i <= x; i++ {

			fmt.Printf("%s\n", brr[i])
		}

	}
}


