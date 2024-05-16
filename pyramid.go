package shapes

type PyramidNode struct {
    BaseCenter Point
    BaseLength float64
    Height     float64
    Children   []*PyramidNode
}

// Insert adds a new PyramidNode as a child
func (node *PyramidNode) Insert(child *PyramidNode) {
    node.Children = append(node.Children, child)
}

// Search for a node with a specific base center
func (node *PyramidNode) Search(baseCenter Point) *PyramidNode {
    if node.BaseCenter == baseCenter {
        return node
    }
    for _, child := range node.Children {
        if result := child.Search(baseCenter); result != nil {
            return result
        }
    }
    return nil
}

