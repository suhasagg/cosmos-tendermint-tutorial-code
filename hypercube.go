package shapes

type HypercubeNode struct {
    Center   []float64  // For n-dimensional space
    SideLength float64
    Children []*HypercubeNode
}

// Insert adds a new HypercubeNode as a child
func (node *HypercubeNode) Insert(child *HypercubeNode) {
    node.Children = append(node.Children, child)
}

// Search for a node with a specific center
func (node *HypercubeNode) Search(center []float64) *HypercubeNode {
    if compareCenters(node.Center, center) {
        return node
    }
    for _, child := range node.Children {
        if result := child.Search(center); result != nil {
            return result
        }
    }
    return nil
}

// Helper function to compare centers
func compareCenters(c1, c2 []float64) bool {
    if len(c1) != len(c2) {
        return false
    }
    for i := range c1 {
        if c1[i] != c2[i] {
            return false
        }
    }
    return true
}

