package popcount

import "testing"

func BenchmarkPopCount1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PopCount1(123456)
    }
}

func BenchmarkPopCount2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PopCount2(123456)
    }
}

func BenchmarkPopCount3(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PopCount3(123456)
    }
}

func BenchmarkPopCount4(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PopCount4(123456)
    }
}
