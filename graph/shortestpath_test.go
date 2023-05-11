package graph

import (
	"fmt"
	"testing"
)

// === RUN   TestShortestPath
// Dijkstra ========>
// FromNode=A | ToNode=E | Distance=7
// FromNode=A | ToNode=A | Distance=0
// FromNode=A | ToNode=B | Distance=6
// FromNode=A | ToNode=C | Distance=1
// FromNode=A | ToNode=D | Distance=5
// Floyd ========>
// FromNode=E | ToNode=B | Distance=9223372036854775807
// FromNode=E | ToNode=C | Distance=9223372036854775807
// FromNode=E | ToNode=D | Distance=9223372036854775807
// FromNode=E | ToNode=E | Distance=0
// FromNode=E | ToNode=A | Distance=9223372036854775807
// FromNode=A | ToNode=E | Distance=7
// FromNode=A | ToNode=A | Distance=0
// FromNode=A | ToNode=B | Distance=6
// FromNode=A | ToNode=C | Distance=1
// FromNode=A | ToNode=D | Distance=5
// FromNode=B | ToNode=D | Distance=7
// FromNode=B | ToNode=E | Distance=9
// FromNode=B | ToNode=A | Distance=9223372036854775807
// FromNode=B | ToNode=B | Distance=0
// FromNode=B | ToNode=C | Distance=3
// FromNode=C | ToNode=E | Distance=6
// FromNode=C | ToNode=A | Distance=9223372036854775807
// FromNode=C | ToNode=B | Distance=9223372036854775807
// FromNode=C | ToNode=C | Distance=0
// FromNode=C | ToNode=D | Distance=4
// FromNode=D | ToNode=C | Distance=9223372036854775807
// FromNode=D | ToNode=D | Distance=0
// FromNode=D | ToNode=E | Distance=2
// FromNode=D | ToNode=A | Distance=9223372036854775807
// FromNode=D | ToNode=B | Distance=9223372036854775807
// --- PASS: TestShortestPath (0.00s)
// PASS
//
// Process finished with the exit code 0
func TestShortestPath(t *testing.T) {
	g := NewGraph[string]()
	A := g.AddNode("A")
	B := g.AddNode("B")
	C := g.AddNode("C")
	D := g.AddNode("D")
	E := g.AddNode("E")

	g.AddEdge(A, B, 6)
	g.AddEdge(A, C, 1)
	g.AddEdge(B, C, 3)
	g.AddEdge(B, D, 7)
	g.AddEdge(C, D, 4)
	g.AddEdge(C, E, 9)
	g.AddEdge(D, E, 2)

	dijkstraDp := g.Dijkstra(A)
	fmt.Println("Dijkstra ========> ")
	for to, distance := range dijkstraDp {
		fmt.Printf("FromNode=%s | ToNode=%s | Distance=%d \n", A.value, to.value, distance)
	}

	floydDp := g.Floyd()
	fmt.Println("Floyd ========> ")
	for form, toDistanceMap := range floydDp {
		for to, distance := range toDistanceMap {
			fmt.Printf("FromNode=%s | ToNode=%s | Distance=%d \n", form.value, to.value, distance)
		}
	}
}
