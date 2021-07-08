package main

import (
	"fmt"
	"strings"
)

// max occurance of words by size of words = ?

func main() {
	str := "SAdfdas sd qw sads sdfds qdasd sd sd as ds as as asd sdf sdf"
	words := strings.Split(str, " ")
	var maxOccurance map[int]int = map[int]int{}
	for _, word := range words {
		maxOccurance[len(word)]++
	}
	maxValue := 0
	maxKey := 0
	for key, value := range maxOccurance {
		if value > maxValue {
			maxValue = value
			maxKey = key
		}
		//fmt.Printf("%d letter word : %d occurance\n", key, value)
	}
	fmt.Printf("%d letter word : %d occurance\n", maxKey, maxValue)
}
