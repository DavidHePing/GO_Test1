package threadSafe

import "testing"

func BenchmarkSyncMapReadWriteDiffKey(b *testing.B) {
	SyncMapReadWriteDiffKey(b.N)
}

func BenchmarkConncurrentMapReadWriteDiffKey(b *testing.B) {
	ConncurrentMapReadWriteDiffKey(b.N)
}

func BenchmarkSyncMapReadWriteSameKey(b *testing.B) {
	SyncMapReadWriteSameKey(b.N)
}

func BenchmarkConncurrentMapReadWriteSameKey(b *testing.B) {
	ConncurrentMapReadWriteSameKey(b.N)
}

func BenchmarkSyncMapRead(b *testing.B) {
	SyncMapRead(b.N)
}

func BenchmarkConncurrentMapRead(b *testing.B) {
	ConncurrentMapRead(b.N)
}
