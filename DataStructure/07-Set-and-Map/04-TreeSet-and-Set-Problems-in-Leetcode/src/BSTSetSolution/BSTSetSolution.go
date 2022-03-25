package main

import (
	"Play-with-Data-Structures/07-Set-and-Map/01-Set-Basics-and-BSTSet/src/BSTSet"
	"bytes"
	"fmt"
)

// Go 中没有 set 类型，这里使用 map 实现
func uniqueMorseRepresentations(words []string) int {
	morseCodes := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	buffer := bytes.Buffer{}
	uniqueWordSet := BSTSet.Constructor()
	for _, word := range words {
		buffer.Reset()
		for _, letter := range word {
			buffer.WriteString(morseCodes[letter-'a'])
		}
		uniqueWordSet.Add(buffer.String())
	}

	return uniqueWordSet.GetSize()
}

func main() {
	words := []string{"gin", "zen", "gig", "msg"}

	fmt.Println(uniqueMorseRepresentations(words))
}
