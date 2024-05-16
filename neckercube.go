package shapes

type NeckerCubeNode struct {
    Center    Point
    Perspective int  // Perspective view of the cube
    Children  []*NeckerCubeNode
}

// Insert adds a new NeckerCubeNode as a child
func (node *NeckerCubeNode) Insert(child *NeckerCubeNode) {
    node.Children = append(node.Children, child)
}

// Search for a node with a specific center and perspective
func (node *NeckerCubeNode) Search(center Point, perspective int) *NeckerCubeNode {
    if node.Center == center && node.Perspective == perspective {
        return node
    }
    for _, child := range node.Children {
        if result := child.Search(center, perspective); result != nil {
            return result
        }
    }
    return nil
}

