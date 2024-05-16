package shapes

type HexagonNode struct {
    Center    Point
    SideLength float64
    Children  []*HexagonNode
}

// Insert adds a new HexagonNode as a child
func (node *HexagonNode) Insert(child *HexagonNode) {
    node.Children = append(node.Children, child)
}

// Search for a node with a specific center
func (node *HexagonNode) Search(center Point) *HexagonNode {
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

