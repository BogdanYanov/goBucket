package main

import (
	"fmt"
	"github.com/BogdanYanov/goBucket/constructor"
)

const (
	BucketSize     = 5
	BucketListSize = 5
	MaxIdxProduct  = 5
)

func main() {
	bucket := constructor.Bucket(BucketSize, BucketListSize, MaxIdxProduct)
	fmt.Println("Bucket before searching:")
	bucket.ShowBucket()
	bucket.Order(1, 2, 3, 5)
	fmt.Printf("\nBucket after searching by function 'Order':\n")
	bucket.ShowBucket()
	bucket.Order2(1, 2, 3, 5)
	fmt.Printf("\nBucket after searching by function 'Order2':\n")
	bucket.ShowBucket()
	bucket.FasterOrder(1, 2, 3, 5)
	fmt.Printf("\nBucket after searching by function 'FasteOrder':\n")
	bucket.ShowBucket()
}
