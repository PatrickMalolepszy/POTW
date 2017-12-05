package main

import (
	"testing"
	"fmt"
)

func TestCase1(t *testing.T) {
	words := [][]byte{[]byte("AWY"), []byte("ZDW"), []byte("AWYR"), []byte("BLAH")}
	originalWords := [][]byte{[]byte("awy"), []byte("ZDw"), []byte("zWYR"), []byte("blah")}
	trie := createTrie(words)
	wordSearch := [][]byte{{byte('A'), byte('W'), byte('H'), byte('V')},
		{byte('W'), byte('O'), byte('Z'), byte('L')},
		{byte('Y'), byte('R'), byte('D'), byte('W')}}
	result := findSnakeWords(wordSearch, &trie)
	fmt.Println(result)
	sortedSolutionWords := sortSolution(words, originalWords, result)
	for i:= range sortedSolutionWords {
		fmt.Println(sortedSolutionWords[i])
	}
}
