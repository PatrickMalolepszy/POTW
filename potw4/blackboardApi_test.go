package main

import "testing"

func addEdge(graph map[int][]int, a, b int) {
	graph[a] = append(graph[a], b)
}

func TestCase1_detectCycle(t *testing.T) {
	graph := make(map[int][]int)
	addEdge(graph, 140, 141)
	addEdge(graph, 141, 140)
	if canGraduate(1, 2, graph) {
		t.Errorf("should not be able to graduate")
	}
}

func TestCase2(t *testing.T) {
	graph := make(map[int][]int)
	addEdge(graph, 140, 141)
	if !canGraduate(2, 2, graph) {
		t.Error()
	}
}

func TestCase3(t *testing.T) {
	graph := make(map[int][]int)
	addEdge(graph, 10, 20)
	addEdge(graph, 20, 30)
	addEdge(graph, 30, 10)
	if canGraduate(1, 3, graph) {
		t.Error()
	}
}

func TestCase4(t *testing.T) {
	graph := make(map[int][]int)
	addEdge(graph, 10, 20)
	addEdge(graph, 10, 30)
	addEdge(graph, 20, 30)
	graph[30] = make([]int, 0)
	if !canGraduate(3, 3, graph) {
		t.Error()
	}
}

func TestCase5_detectCycle(t *testing.T) {
	graph := make(map[int][]int)

	addEdge(graph, 10,20)
	addEdge(graph, 10,30)
	addEdge(graph, 10,40)

	addEdge(graph, 30, 60)
	addEdge(graph, 30,50)

	addEdge(graph, 60,20)

	addEdge(graph, 50,40)

	addEdge(graph, 40,10)

	if canGraduate(1, 6, graph) {
		t.Error()
	}
}
