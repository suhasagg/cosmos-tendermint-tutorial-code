package shapes

type OctahedronNode struct {
    Vertices  [6]Point
    Children  []*OctahedronNode
}

// Insert adds a new OctahedronNode as a child
func (node *OctahedronNode) Insert(child *OctahedronNode) {
    node.Children = append(node.Children, child)
}

// Search for a node with specific vertices
func (node *OctahedronNode) Search(vertices [6]Point) *OctahedronNode {
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

