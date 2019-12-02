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

// Order2 search order items in bucket vertically using labels instead 'for'
func (b *Bucket) Order2(order ...int) {
	var (
		curList int
		curItem int
	)

Order:
	for _, val := range order {
		curList = 0
		curItem = 0

	Search:
		if val == (*b)[curList][curItem] {
			(*b)[curList][curItem] = 0
			continue Order
		} else {
			if curList < len(*b)-1 {
				curList++
				goto Search
			} else if curItem < len((*b)[0])-1 {
				curList = 0
				curItem++
				goto Search
			} else {
				continue Order
			}
		}
	}
}

// FasterOrder search order items in bucket vertically with 'for' and labels
// Works faster than other searching functions
func (b *Bucket) FasterOrder(order ...int) {

	itemPos := 0
	founded := false

	for _, val := range order {
		itemPos = 0
		founded = false
	SearchLoop:
		for _, list := range *b {
			if val == list[itemPos] {
				list[itemPos] = 0
				founded = true
				break SearchLoop
			}
		}
		if !founded {
			if itemPos < len((*b)[0])-1 {
				itemPos++
				goto SearchLoop
			}
		}
	}
}
