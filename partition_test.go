package txdis

import (
	"fmt"
	"math"
	"testing"
)

func Test_partition_Calculate(t *testing.T) {
	salt := []byte(defaultSaltString)
	size := 10

	type fields struct {
		salt []byte
		size int
	}
	type args struct {
		id int64
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "case 1",
			fields: fields{
				salt: salt,
				size: size,
			},
			args: args{id: 3100140277596160},
			want: 0,
		},
		{
			name: "case 2",
			fields: fields{
				salt: salt,
				size: size,
			},
			args: args{id: 3679289639698433},
			want: 9,
		},
		{
			name: "case 3",
			fields: fields{
				salt: salt,
				size: size,
			},
			args: args{id: 3831135389876225},
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := newPartition(tt.fields.salt, tt.fields.size)
			if got := p.Calculate(tt.args.id); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 檢查計算出的值，落在的區間是否平均
func Test_partition_Calculate_location_average(t *testing.T) {
	location := make(map[int]int)
	calculateTime := 100000
	salt := []byte(defaultSaltString)
	size := 10

	p := newPartition(salt, size)

	for i := 0; i < calculateTime; i++ {
		loc := p.Calculate(_random.Int63())
		location[loc]++
	}

	// 檢查長度
	if len(location) != size {
		t.Errorf("partition size: got=%d, want=%d", len(location), size)
	}

	// 檢查位置
	for i := 0; i < size; i++ {
		if _, ok := location[i]; !ok {
			t.Errorf("location not found: want=%d", i)
		}
	}

	// 檢查誤差
	avgCount := calculateTime / size // 平均值
	rangeBuffer := avgCount / size   // 誤差容忍範圍
	for i := 0; i < len(location); i++ {
		count := location[i]
		if float64(rangeBuffer) > math.Abs(float64(count-avgCount)) {
			// 要看再放開
			//t.Log(fmt.Sprintf("loc: %d, range: %d, count: %d", i, rangeBuffer, count))
		} else {
			t.Error(fmt.Sprintf("loc: %d, range: %d, count: %d", i, rangeBuffer, count))
		}

	}

}

// BenchmarkBalance-8   	 5000000	       308 ns/op	      64 B/op	       1 allocs/op
func Benchmark_partition_Calculate(b *testing.B) {
	p := newPartition([]byte(defaultSaltString), 10)

	for i := 0; i < b.N; i++ {
		p.Calculate(3100140277596160)
	}
}
