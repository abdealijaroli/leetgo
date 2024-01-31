package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freq := [26]int{}

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}

	for i := 0; i < len(freq); i++ {
		if freq[i] != 0 {
			return false
		}
	}
	return true
}

func altIsAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    m := map[rune]int{}

    for _, char := range s {
        m[char]++
    }

    for _, char := range t {
        if _, c := m[char]; !c || m[char] == 0 {
            return false
        }
        m[char]--
    }
    return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(altIsAnagram("anagram", "nagaram"))
}
