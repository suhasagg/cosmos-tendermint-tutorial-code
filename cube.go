package shapes

type CubeNode struct {
    Center   Point
    SideLength float64
    Children []*CubeNode
}

// Insert adds a new CubeNode as a child
func (node *CubeNode) Insert(child *CubeNode) {
    node.Children = append(node.Children, child)
}

// Search for a node with a specific center
func (node *CubeNode) Search(center Point) *CubeNode {
    if node.Center == center {
        return node
    }
    for _, child := range node.Children {
        if result := child.Search(center); result != nil {
            return result
        }
    }
    return nil
}

