package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

type Node struct {
    id int
    rumors []int
    neighbors []*Node
}

func NewNode(id int) *Node {
    return &Node{id: id}
}

func (n *Node) addNeighbor(neighbor *Node) {
    n.neighbors = append(n.neighbors, neighbor)
}

func (n *Node) spreadRumor(wg *sync.WaitGroup) {
    defer wg.Done()
    rumor := n.rumors[rand.Intn(len(n.rumors))]
    neighbor := n.neighbors[rand.Intn(len(n.neighbors))]
    neighbor.receiveRumor(rumor)
}

func (n *Node) receiveRumor(rumor int) {
    for _, r := range n.rumors {
        if r == rumor {
            // rumor already heard, do not propagate
            return
        }
    }
    n.rumors = append(n.rumors, rumor)
    var wg sync.WaitGroup
    for _, neighbor := range n.neighbors {
        wg.Add(1)
        go neighbor.spreadRumor(&wg)
    }
    wg.Wait()
}

func main() {
    // create nodes
    nodes := make([]*Node, 10)
    for i := range nodes {
        nodes[i] = NewNode(i)
    }
    // add neighbors
    for i := range nodes {
        for j := range nodes {
            if i != j {
                nodes[i].addNeighbor(nodes[j])
            }
        }
    }
    // start rumor
    rand.Seed(time.Now().UnixNano())
    startingNode := nodes[rand.Intn(len(nodes))]
    startingNode.rumors = []int{42}
    // spread rumor
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        node := nodes[rand.Intn(len(nodes))]
        wg.Add(1)
        go node.spreadRumor(&wg)
    }
    wg.Wait()
    // print results
    for _, node := range nodes {
        fmt.Printf("Node %d heard rumors: %v\n", node.id, node.rumors)
    }
}
