package model

import "fmt"

type Bucket [][]int

// ShowBucket prints lists with items in bucket.
func (b *Bucket) ShowBucket() {
	for i, bVal := range *b {
		fmt.Printf("List %d: %v\n", i, bVal)
	}
}
