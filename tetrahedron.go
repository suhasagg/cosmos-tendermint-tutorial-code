package shapes

type TetrahedronNode struct {
    Vertices  [4]Point
    Children  []*TetrahedronNode
}

// Insert adds a new TetrahedronNode as a child
func (node *TetrahedronNode) Insert(child *TetrahedronNode) {
    node.Children = append(node.Children, child)
}

// Search for a node with specific vertices
func (node *TetrahedronNode) Search(vertices [4]Point) *TetrahedronNode {
    if node.Vertices == vertices {
        return node
    }
    for _, child := range node.Children {
        if result := child.Search(vertices); result != nil {
            return result
        }
    }
    return nil
}

