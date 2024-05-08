package main

import (
	"fmt"
)

// Define a structure for a 2D point
type Point struct {
	x, y float64
}

// Define a structure for a quadtree node
type QuadtreeNode struct {
	x, y, width, height float64
	points              []Point
	children            [4]*QuadtreeNode
}

// Create a new quadtree node
func NewQuadtreeNode(x, y, width, height float64) *QuadtreeNode {
	return &QuadtreeNode{
		x:       x,
		y:       y,
		width:   width,
		height:  height,
		points:  []Point{},
		children: [4]*QuadtreeNode{nil, nil, nil, nil},
	}
}

// Insert a point into the quadtree
func (node *QuadtreeNode) Insert(point Point) {
	if len(node.points) < 4 {
		node.points = append(node.points, point)
		return
	}

	// If the node doesn't have children, create them
	if node.children[0] == nil {
		node.subdivide()
	}

	// Insert the point into one of the child nodes
	for i := 0; i < 4; i++ {
		child := node.children[i]
		if point.x >= child.x && point.x < child.x+child.width &&
			point.y >= child.y && point.y < child.y+child.height {
			child.Insert(point)
			break
		}
	}
}

// Subdivide the node into four equal parts
func (node *QuadtreeNode) subdivide() {
	childWidth := node.width / 2
	childHeight := node.height / 2
	x := node.x
	y := node.y

	node.children[0] = NewQuadtreeNode(x, y, childWidth, childHeight)
	node.children[1] = NewQuadtreeNode(x+childWidth, y, childWidth, childHeight)
	node.children[2] = NewQuadtreeNode(x, y+childHeight, childWidth, childHeight)
	node.children[3] = NewQuadtreeNode(x+childWidth, y+childHeight, childWidth, childHeight)
}

// Query points within a given range
func (node *QuadtreeNode) QueryRange(x, y, width, height float64) []Point {
	var result []Point

	// If the range doesn't intersect with the node, return empty result
	if x > node.x+node.width || x+width < node.x || y > node.y+node.height || y+height < node.y {
		return result
	}

	// Check points within the node
	for _, point := range node.points {
		if point.x >= x && point.x <= x+width && point.y >= y && point.y <= y+height {
			result = append(result, point)
		}
	}

	// Recursively query child nodes
	for _, child := range node.children {
		if child == nil {
			continue
		}
		result = append(result, child.QueryRange(x, y, width, height)...)
	}

	return result
}

func main() {
	// Create a quadtree with bounds (0, 0, 100, 100)
	quadtree := NewQuadtreeNode(0, 0, 100, 100)

	// Insert some points into the quadtree
	quadtree.Insert(Point{10, 10})
	quadtree.Insert(Point{20, 20})
	quadtree.Insert(Point{80, 80})
	quadtree.Insert(Point{90, 90})

	// Query points within a range
	queryResult := quadtree.QueryRange(0, 0, 50, 50)
	fmt.Println("Points within range:", queryResult)
}

