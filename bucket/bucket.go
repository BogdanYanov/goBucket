package bucket

import (
	"fmt"
	"math/rand"
	"time"
)

// Bucket is a two-dimensional array with the addition of methods for finding the best product
// and creating a basket with randomly filled numbers.
type Bucket [][]int

// New create Bucket with number of list sizeL, and sizeI items in each list.
// maxProdIdx indicates the maximum product index.
func New(sizeL, sizeI, maxProdIdx int) *Bucket {
	rand.Seed(time.Now().UnixNano())

	b := make(Bucket, 0, sizeL)

	for i := 0; i < sizeL; i++ {
		b = append(b, make([]int, 0, sizeI))
		for j := 0; j < sizeI; j++ {
			b[i] = append(b[i], rand.Intn(maxProdIdx)+1)
		}
	}

	return &b
}

// ShowBucket prints lists with items in bucket.
func (b *Bucket) ShowBucket() {
	for i, bVal := range *b {
		fmt.Printf("List %d: %v\n", i, bVal)
	}
}

// ReceiveOrderSlow search order items in bucket.
// The algorithm search items by comparing values optimalItem and current found item.
func (b *Bucket) ReceiveOrderSlow(order ...int) Bucket {

	newBucket := copyBucket(*b)

	var (
		optimalList = -1
		optimalItem = -1
	)

	for _, val := range order {
		optimalList = -1
		optimalItem = -1
	SearchLoop:
		for i, list := range *b {
		SearchItemList:
			for j, item := range list {
				if val == item {
					if j >= optimalItem && optimalItem != -1 {
						break SearchItemList
					}

					optimalList = i
					optimalItem = j

					if j == 0 {
						break SearchLoop
					}

					break SearchItemList
				}
			}
		}

		if optimalList != -1 || optimalItem != -1 {
			newBucket[optimalList][optimalItem] = 0
		}
	}

	removeZero(&newBucket)
	return newBucket
}

// ReceiveOrderFast search order items in bucket vertically with 'for' and labels
// Works faster than other searching functions
func (b *Bucket) ReceiveOrderFast(order ...int) Bucket {

	newBucket := copyBucket(*b)

	var (
		itemPos int
		listPos int
	)

	for _, val := range order {
		itemPos = 0
		listPos = 0

	SearchLoop:
		for {
			if itemPos >= len(newBucket[listPos]) {
				if listPos == len(newBucket)-1 {
					break SearchLoop
				}

				listPos++
				continue SearchLoop
			}

			if val == newBucket[listPos][itemPos] {
				newBucket[listPos][itemPos] = 0
				break SearchLoop
			}

			if listPos < len(newBucket)-1 {
				listPos++
				continue SearchLoop
			}

			itemPos++
			listPos = 0
		}
	}

	removeZero(&newBucket)
	return newBucket
}

// ReceiveOrder search order items in bucket vertically with only one infinity 'for'
func (b *Bucket) ReceiveOrder(order ...int) Bucket {

	newBucket := copyBucket(*b)

	var (
		itemPos  int = 0
		listPos  int = 0
		orderPos int = 0
	)

SearchItems:
	for {
		if itemPos >= len(newBucket[listPos]) {
			if listPos == len(newBucket)-1 {
				if orderPos == len(order)-1 {
					break SearchItems
				}
				orderPos++
				listPos = -1
				itemPos = 0
			}

			listPos++
			continue SearchItems
		}

		if order[orderPos] == newBucket[listPos][itemPos] {
			newBucket[listPos][itemPos] = 0

			if orderPos == len(order)-1 {
				break SearchItems
			}

			orderPos++
			itemPos = 0
			listPos = 0
			continue SearchItems
		} else {
			if listPos == len(newBucket)-1 {
				listPos = -1
				itemPos++
			}

			listPos++
			continue SearchItems
		}
	}

	removeZero(&newBucket)
	return newBucket
}

func copyBucket(srcBucket Bucket) Bucket {
	newBucket := make(Bucket, len(srcBucket), cap(srcBucket))

	for i := range srcBucket {
		newBucket[i] = make([]int, len(srcBucket[i]), cap(srcBucket[i]))
		copy(newBucket[i], srcBucket[i])
	}

	return newBucket
}

func removeZero(bucket *Bucket) {
	for i := 0; i < len(*bucket); i++ {
		for j := 0; j < len((*bucket)[i]); j++ {
			if (*bucket)[i][j] == 0 {
				(*bucket)[i] = append((*bucket)[i][:j], (*bucket)[i][j+1:]...)
				j--
			}
		}
	}
}
