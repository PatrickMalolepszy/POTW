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

func TestCase2(t *testing.T) {
	words := [][]byte{[]byte("vlwdryw")}
	originalWords := [][]byte{[]byte("vlwdryw")}
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

func TestCase3(t *testing.T) {
	words := [][]byte{[]byte("ABC"), []byte("ADC")}
	originalWords := [][]byte{[]byte("ABC"), []byte("ADC")}
	trie := createTrie(words)
	wordSearch := [][]byte{{byte('A'), byte('B'), byte('H'), byte('V')},
							{byte('D'), byte('C'), byte('Z'), byte('L')},
							{byte('Y'), byte('R'), byte('D'), byte('W')}}
	result := findSnakeWords(wordSearch, &trie)
	fmt.Println(result)
	sortedSolutionWords := sortSolution(words, originalWords, result)
	for i:= range sortedSolutionWords {
		fmt.Println(sortedSolutionWords[i])
	}
}
