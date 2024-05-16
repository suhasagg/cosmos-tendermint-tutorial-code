package shapes

type Point struct {
    X, Y, Z float64
}

type SphereNode struct {
    Center   Point
    Radius   float64
    Children []*SphereNode
}

// Insert adds a new SphereNode as a child
func (node *SphereNode) Insert(child *SphereNode) {
    node.Children = append(node.Children, child)
}

// Search for a node within a certain radius
func (node *SphereNode) Search(center Point, radius float64) *SphereNode {
    if node.Center == center && node.Radius == radius {
        return node
    }
    for _, child := range node.Children {
        if result := child.Search(center, radius); result != nil {
            return result
        }
    }
    return nil
}

