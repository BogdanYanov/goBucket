package constructor

import (
	"github.com/BogdanYanov/goBucket/model"
	"math/rand"
	"time"
)

// Bucket creating bucket with sizeB lists, and sizeL items in each list.
// The maximum product number is indicated in maxProdIdx
func Bucket(sizeB, sizeL, maxProdIdx int) *model.Bucket {
	rand.Seed(time.Now().UnixNano())

	b := make(model.Bucket, 0, sizeB)

	for i := 0; i < sizeB; i++ {
		b = append(b, make([]int, 0, sizeL))
		for j := 0; j < sizeL; j++ {
			b[i] = append(b[i], rand.Intn(maxProdIdx)+1)
		}
	}

	return &b
}
