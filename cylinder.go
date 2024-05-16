package shapes

type CylinderNode struct {
    BaseCenter Point
    Radius     float64
    Height     float64
    Children   []*CylinderNode
}

// Insert adds a new CylinderNode as a child
func (node *CylinderNode) Insert(child *CylinderNode) {
    node.Children = append(node.Children, child)
}

// Search for a node with a specific base center
func (node *CylinderNode) Search(baseCenter Point) *CylinderNode {
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

