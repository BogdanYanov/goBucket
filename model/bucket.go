package model

import "fmt"

type Bucket [][]int

// ShowBucket prints lists with items in bucket.
func (b *Bucket) ShowBucket() {
	for i, bVal := range *b {
		fmt.Printf("List %d: %v\n", i, bVal)
	}
}

// Order search order items in bucket.
// The algorithm search items by comparing values in oi (OptimalItem) and current found item.
func (b *Bucket) Order(order ...int) {
	oi := OptimalItem{}
	for _, val := range order {
		oi.ListNum = -1
		oi.ItemInList = -1
	SearchLoop:
		for i, list := range *b {
			for j, item := range list {
				if val == item {
					if j == 0 {
						oi.ListNum = i
						oi.ItemInList = j
						break SearchLoop
					} else if oi.ItemInList == -1 {
						oi.ListNum = i
						oi.ItemInList = j
					} else if j < oi.ItemInList {
						oi.ListNum = i
						oi.ItemInList = j
					}
				}
			}
		}

		if oi.ListNum != -1 || oi.ItemInList != -1 {
			(*b)[oi.ListNum][oi.ItemInList] = 0
		}
	}
}
