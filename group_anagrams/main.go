package main

import (
	"fmt"
)

func groupAnagrams(strs []string) [][]string {
	// create a map of key: integer array of length 26 (count array) and val: array of strings

	// for every string in strs array we are going to have an equivalent count array that will have the count of each and every letter of that string

	// the count array will be our key to finding more strings in the strs array

	// we will append them in the map as we keep finding more strings that match the count array

	// at last, the vals of our map will have multiple string arrays now all of those arrays will be appended to a new array which will be our result

	m := make(map[[26]int][]string)

	for _, s := range strs {
		count := [26]int{}

		for _, c := range s {
			count[c-'a']++
		}

		m[count] = append(m[count], s)

	}

	var res [][]string

	for _, val := range m {
		res = append(res, val)
	}

	return res

}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
