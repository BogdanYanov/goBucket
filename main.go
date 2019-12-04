package main

import (
	"fmt"
	"github.com/BogdanYanov/goBucket/bucket"
)

const (
	bucketSize     = 10
	bucketListSize = 10
	maxIdxProduct  = 5
)

func main() {
	myBucket := bucket.New(bucketSize, bucketListSize, maxIdxProduct)
	newBucket := myBucket.ReceiveOrder(1, 2, 3, 4, 5)
	fmt.Println("Bucket before order")
	myBucket.ShowBucket()
	fmt.Println("\nBucket after order")
	newBucket.ShowBucket()
}
