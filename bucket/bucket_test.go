package bucket

import "testing"

const (
	BucketSize     = 5
	BucketListSize = 5
	MaxIdxProduct  = 5
)

var myBucket = New(BucketSize, BucketListSize, MaxIdxProduct)

func BenchmarkBucket_ReceiveOrderSlow(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		myBucket.ReceiveOrderSlow(1, 2, 3, 4, 5)
	}
}

func BenchmarkBucket_ReceiveOrderFast(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		myBucket.ReceiveOrderFast(1, 2, 3, 4, 5)
	}
}

func BenchmarkBucket_ReceiveOrder(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		myBucket.ReceiveOrder(1, 2, 3, 4, 5)
	}
}
