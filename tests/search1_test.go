package tests

import (
	"github.com/BogdanYanov/goBucket/constructor"
	"testing"
)

func BenchmarkSearch(b *testing.B) {
	b.ResetTimer()
	b.SetBytes(2)
	Bucket := constructor.Bucket(MaxLists, MaxItems, MaxProdIdx)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Bucket.Order(1, 2, 3, 5)
	}
}
