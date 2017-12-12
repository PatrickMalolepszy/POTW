// POTW #5 - Fall 2017
// Snake Word Search - https://potw.cs.uwindsor.ca/problem/2017/5/
// Patrick Malolepszy
package main

import (
	"bufio"
	"os"
	"fmt"
	"sort"
	"strings"
)

// Time to use some tries :D
type trieNode struct {
	wordEnd  bool
	children map[byte]*trieNode
}

// inserts word into trie by adding nodes as needed, and marking the last one as a word.
func (node *trieNode) insert(word []byte) {
	curr := node
	for i := range word {
		c := word[i]
		if curr.children == nil {
			curr.children = make(map[byte]*trieNode)
		}
		if curr.children[c] == nil {
			curr.children[c] = &trieNode{}
		}
		curr = curr.children[c]
	}
	curr.wordEnd = true
}

// Avoids having to convert buffered bytes to a string, and then to a int.
// Source: https://stackoverflow.com/questions/31333353/faster-input-scanning
func readInt(scanner *bufio.Scanner) (n int) {
	scanner.Scan()
	buf := scanner.Bytes()
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}

// Converts an ASCII character to its uppercase version (as bytes).
func toUpperByte(word []byte) {
	for i := range word {
		if word[i] > 'Z' {
			word[i] -= 32
		}
	}
}

// Returns the wordsearch puzzle, a case-insensitive list of read words, and the original set of words.
func readPotw5Input() ([][]byte, [][]byte, [][]byte) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	r := readInt(scanner)
	c := readInt(scanner)
	wordSearch := make([][]byte, r)
	for i := range wordSearch {
		wordSearch[i] = make([]byte, c)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			scanner.Scan()
			wordSearch[i][j] = scanner.Bytes()[0]
		}
	}
	n := readInt(scanner)
	words := make([][]byte, n)
	originalWords := make([][]byte, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		words[i] = scanner.Bytes()
		originalWords[i] = make([]byte, len(words[i]))
		copy(originalWords[i], words[i])
		toUpperByte(words[i])
	}
	return wordSearch, words, originalWords
}

// making some sweet sweet tries.
func createTrie(words [][]byte) trieNode {
	root := trieNode{}
	for i := range words {
		root.insert(words[i])
	}
	return root
}

// Traverse snakes out from each point of the wordsearch
func findSnakeWords(wordSearch [][]byte, trie *trieNode) []string {
	foundWords := make(map[string]bool)
	for i := range wordSearch {
		for j := range wordSearch[i] {
			visited := make([][]bool, len(wordSearch))
			for k := range visited {
				visited[k] = make([]bool, len(wordSearch[0]))
			}
			dfs(i, j, wordSearch, visited, trie, &foundWords, make([]byte, 0, 100))
		}
	}
	keySet := make([]string, 0, len(foundWords))
	for k := range foundWords {
		keySet = append(keySet, k)
	}
	return keySet
}

// Depth first search to traverse the wordsearch puzzle in all four directions, and keep track of all words found.
func dfs(x, y int, wordSearch [][]byte, visited [][]bool, trie *trieNode, foundWords *map[string]bool, curr []byte) {
	if visited[x][y] {
		return
	}

	visited[x][y] = true
	c := wordSearch[x][y]
	curr = append(curr, c)
	trie = trie.children[c]

	if trie == nil {
		return
	}

	if trie.wordEnd {
		(*foundWords)[string(curr)] = true
	}

	// copy not reference of visted nodes.
	if x+1 < len(wordSearch) {
		dfs(x+1, y, wordSearch, copy2dArray(visited), trie, foundWords, curr)
	}
	if x-1 >= 0 {
		dfs(x-1, y, wordSearch, copy2dArray(visited), trie, foundWords, curr)
	}
	if y+1 < len(wordSearch[x]) {
		dfs(x, y+1, wordSearch, copy2dArray(visited), trie, foundWords, curr)
	}
	if y-1 >= 0 {
		dfs(x, y-1, wordSearch, copy2dArray(visited), trie, foundWords, curr)
	}
}

func copy2dArray(a [][]bool) [][]bool {
	newA := make([][]bool, len(a))
	for i := range a {
		newA[i] = make([]bool, len(a[i]))
		copy(a[i], newA[i])
	}
	return newA
}

// defining a sort function which does not take capitals into account
type ByAlphaIgnoreCaps []string
func (s ByAlphaIgnoreCaps) Len() int {
	return len(s)
}
func (s ByAlphaIgnoreCaps) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByAlphaIgnoreCaps) Less(i, j int) bool {
	return strings.Compare(strings.ToLower(s[i]), strings.ToLower(s[j])) < 0
}

func sortSolution(words [][]byte, originalWords [][]byte, foundWords []string) []string {
	result := make([]string, 0, 10)
	for i := range foundWords {
		for j := range words { // max of 10 words, quick loop is faster then the overhead of setting up maps.
			if string(words[j]) == foundWords[i] {
				result = append(result, string(originalWords[j]))
				break
			}
		}
	}
	sort.Sort(ByAlphaIgnoreCaps(result))
	return result
}

// simple main => simple life
func main() {
	wordSearch, words, originalWords := readPotw5Input()
	trie := createTrie(words)
	foundWords := findSnakeWords(wordSearch, &trie)
	sortedSolutionWords := sortSolution(words, originalWords, foundWords)
	for i:= range sortedSolutionWords {
		fmt.Println(sortedSolutionWords[i])
	}
}
