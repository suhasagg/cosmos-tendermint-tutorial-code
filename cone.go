package shapes

type ConeNode struct {
    BaseCenter Point
    Radius     float64
    Height     float64
    Children   []*ConeNode
}

// Insert adds a new ConeNode as a child
func (node *ConeNode) Insert(child *ConeNode) {
    node.Children = append(node.Children, child)
}

// Search for a node with a specific base center
func (node *ConeNode) Search(baseCenter Point) *ConeNode {
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

