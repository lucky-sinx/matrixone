// +build amd64

package sub

import "golang.org/x/sys/cpu"

func int8SubAvx2Asm([]int8, []int8, []int8)
func int8SubScalarAvx2Asm(int8, []int8, []int8)
func int8SubByScalarAvx2Asm(int8, []int8, []int8)
func int16SubAvx2Asm([]int16, []int16, []int16)
func int16SubScalarAvx2Asm(int16, []int16, []int16)
func int16SubByScalarAvx2Asm(int16, []int16, []int16)
func int32SubAvx2Asm([]int32, []int32, []int32)
func int32SubScalarAvx2Asm(int32, []int32, []int32)
func int32SubByScalarAvx2Asm(int32, []int32, []int32)
func int64SubAvx2Asm([]int64, []int64, []int64)
func int64SubScalarAvx2Asm(int64, []int64, []int64)
func int64SubByScalarAvx2Asm(int64, []int64, []int64)
func uint8SubAvx2Asm([]uint8, []uint8, []uint8)
func uint8SubScalarAvx2Asm(uint8, []uint8, []uint8)
func uint8SubByScalarAvx2Asm(uint8, []uint8, []uint8)
func uint16SubAvx2Asm([]uint16, []uint16, []uint16)
func uint16SubScalarAvx2Asm(uint16, []uint16, []uint16)
func uint16SubByScalarAvx2Asm(uint16, []uint16, []uint16)
func uint32SubAvx2Asm([]uint32, []uint32, []uint32)
func uint32SubScalarAvx2Asm(uint32, []uint32, []uint32)
func uint32SubByScalarAvx2Asm(uint32, []uint32, []uint32)
func uint64SubAvx2Asm([]uint64, []uint64, []uint64)
func uint64SubScalarAvx2Asm(uint64, []uint64, []uint64)
func uint64SubByScalarAvx2Asm(uint64, []uint64, []uint64)
func float32SubAvx2Asm([]float32, []float32, []float32)
func float32SubScalarAvx2Asm(float32, []float32, []float32)
func float32SubByScalarAvx2Asm(float32, []float32, []float32)
func float64SubAvx2Asm([]float64, []float64, []float64)
func float64SubScalarAvx2Asm(float64, []float64, []float64)
func float64SubByScalarAvx2Asm(float64, []float64, []float64)

func int8SubAvx512Asm([]int8, []int8, []int8)
func int8SubScalarAvx512Asm(int8, []int8, []int8)
func int8SubByScalarAvx512Asm(int8, []int8, []int8)
func int16SubAvx512Asm([]int16, []int16, []int16)
func int16SubScalarAvx512Asm(int16, []int16, []int16)
func int16SubByScalarAvx512Asm(int16, []int16, []int16)
func int32SubAvx512Asm([]int32, []int32, []int32)
func int32SubScalarAvx512Asm(int32, []int32, []int32)
func int32SubByScalarAvx512Asm(int32, []int32, []int32)
func int64SubAvx512Asm([]int64, []int64, []int64)
func int64SubScalarAvx512Asm(int64, []int64, []int64)
func int64SubByScalarAvx512Asm(int64, []int64, []int64)
func uint8SubAvx512Asm([]uint8, []uint8, []uint8)
func uint8SubScalarAvx512Asm(uint8, []uint8, []uint8)
func uint8SubByScalarAvx512Asm(uint8, []uint8, []uint8)
func uint16SubAvx512Asm([]uint16, []uint16, []uint16)
func uint16SubScalarAvx512Asm(uint16, []uint16, []uint16)
func uint16SubByScalarAvx512Asm(uint16, []uint16, []uint16)
func uint32SubAvx512Asm([]uint32, []uint32, []uint32)
func uint32SubScalarAvx512Asm(uint32, []uint32, []uint32)
func uint32SubByScalarAvx512Asm(uint32, []uint32, []uint32)
func uint64SubAvx512Asm([]uint64, []uint64, []uint64)
func uint64SubScalarAvx512Asm(uint64, []uint64, []uint64)
func uint64SubByScalarAvx512Asm(uint64, []uint64, []uint64)
func float32SubAvx512Asm([]float32, []float32, []float32)
func float32SubScalarAvx512Asm(float32, []float32, []float32)
func float32SubByScalarAvx512Asm(float32, []float32, []float32)
func float64SubAvx512Asm([]float64, []float64, []float64)
func float64SubScalarAvx512Asm(float64, []float64, []float64)
func float64SubByScalarAvx512Asm(float64, []float64, []float64)

func init() {
	if cpu.X86.HasAVX512 {
		Int8Sub = int8SubAvx512
		Int8SubScalar = int8SubScalarAvx512
		Int8SubByScalar = int8SubByScalarAvx512
		Int16Sub = int16SubAvx512
		Int16SubScalar = int16SubScalarAvx512
		Int16SubByScalar = int16SubByScalarAvx512
		Int32Sub = int32SubAvx512
		Int32SubScalar = int32SubScalarAvx512
		Int32SubByScalar = int32SubByScalarAvx512
		Int64Sub = int64SubAvx512
		Int64SubScalar = int64SubScalarAvx512
		Int64SubByScalar = int64SubByScalarAvx512
		Uint8Sub = uint8SubAvx512
		Uint8SubScalar = uint8SubScalarAvx512
		Uint8SubByScalar = uint8SubByScalarAvx512
		Uint16Sub = uint16SubAvx512
		Uint16SubScalar = uint16SubScalarAvx512
		Uint16SubByScalar = uint16SubByScalarAvx512
		Uint32Sub = uint32SubAvx512
		Uint32SubScalar = uint32SubScalarAvx512
		Uint32SubByScalar = uint32SubByScalarAvx512
		Uint64Sub = uint64SubAvx512
		Uint64SubScalar = uint64SubScalarAvx512
		Uint64SubByScalar = uint64SubByScalarAvx512
		Float32Sub = float32SubAvx512
		Float32SubScalar = float32SubScalarAvx512
		Float32SubByScalar = float32SubByScalarAvx512
		Float64Sub = float64SubAvx512
		Float64SubScalar = float64SubScalarAvx512
		Float64SubByScalar = float64SubByScalarAvx512
	} else if cpu.X86.HasAVX2 {
		Int8Sub = int8SubAvx2
		Int8SubScalar = int8SubScalarAvx2
		Int8SubByScalar = int8SubByScalarAvx2
		Int16Sub = int16SubAvx2
		Int16SubScalar = int16SubScalarAvx2
		Int16SubByScalar = int16SubByScalarAvx2
		Int32Sub = int32SubAvx2
		Int32SubScalar = int32SubScalarAvx2
		Int32SubByScalar = int32SubByScalarAvx2
		Int64Sub = int64SubAvx2
		Int64SubScalar = int64SubScalarAvx2
		Int64SubByScalar = int64SubByScalarAvx2
		Uint8Sub = uint8SubAvx2
		Uint8SubScalar = uint8SubScalarAvx2
		Uint8SubByScalar = uint8SubByScalarAvx2
		Uint16Sub = uint16SubAvx2
		Uint16SubScalar = uint16SubScalarAvx2
		Uint16SubByScalar = uint16SubByScalarAvx2
		Uint32Sub = uint32SubAvx2
		Uint32SubScalar = uint32SubScalarAvx2
		Uint32SubByScalar = uint32SubByScalarAvx2
		Uint64Sub = uint64SubAvx2
		Uint64SubScalar = uint64SubScalarAvx2
		Uint64SubByScalar = uint64SubByScalarAvx2
		Float32Sub = float32SubAvx2
		Float32SubScalar = float32SubScalarAvx2
		Float32SubByScalar = float32SubByScalarAvx2
		Float64Sub = float64SubAvx2
		Float64SubScalar = float64SubScalarAvx2
		Float64SubByScalar = float64SubByScalarAvx2
	} else {
		Int8Sub = int8Sub
		Int8SubScalar = int8SubScalar
		Int8SubByScalar = int8SubByScalar
		Int16Sub = int16Sub
		Int16SubScalar = int16SubScalar
		Int16SubByScalar = int16SubByScalar
		Int32Sub = int32Sub
		Int32SubScalar = int32SubScalar
		Int32SubByScalar = int32SubByScalar
		Int64Sub = int64Sub
		Int64SubScalar = int64SubScalar
		Int64SubByScalar = int64SubByScalar
		Uint8Sub = uint8Sub
		Uint8SubScalar = uint8SubScalar
		Uint8SubByScalar = uint8SubByScalar
		Uint16Sub = uint16Sub
		Uint16SubScalar = uint16SubScalar
		Uint16SubByScalar = uint16SubByScalar
		Uint32Sub = uint32Sub
		Uint32SubScalar = uint32SubScalar
		Uint32SubByScalar = uint32SubByScalar
		Uint64Sub = uint64Sub
		Uint64SubScalar = uint64SubScalar
		Uint64SubByScalar = uint64SubByScalar
		Float32Sub = float32Sub
		Float32SubScalar = float32SubScalar
		Float32SubByScalar = float32SubByScalar
		Float64Sub = float64Sub
		Float64SubScalar = float64SubScalar
		Float64SubByScalar = float64SubByScalar
	}

	Int8SubSels = int8SubSels
	Int8SubScalarSels = int8SubScalarSels
	Int8SubByScalarSels = int8SubByScalarSels
	Int16SubSels = int16SubSels
	Int16SubScalarSels = int16SubScalarSels
	Int16SubByScalarSels = int16SubByScalarSels
	Int32SubSels = int32SubSels
	Int32SubScalarSels = int32SubScalarSels
	Int32SubByScalarSels = int32SubByScalarSels
	Int64SubSels = int64SubSels
	Int64SubScalarSels = int64SubScalarSels
	Int64SubByScalarSels = int64SubByScalarSels
	Uint8SubSels = uint8SubSels
	Uint8SubScalarSels = uint8SubScalarSels
	Uint8SubByScalarSels = uint8SubByScalarSels
	Uint16SubSels = uint16SubSels
	Uint16SubScalarSels = uint16SubScalarSels
	Uint16SubByScalarSels = uint16SubByScalarSels
	Uint32SubSels = uint32SubSels
	Uint32SubScalarSels = uint32SubScalarSels
	Uint32SubByScalarSels = uint32SubByScalarSels
	Uint64SubSels = uint64SubSels
	Uint64SubScalarSels = uint64SubScalarSels
	Uint64SubByScalarSels = uint64SubByScalarSels
	Float32SubSels = float32SubSels
	Float32SubScalarSels = float32SubScalarSels
	Float32SubByScalarSels = float32SubByScalarSels
	Float64SubSels = float64SubSels
	Float64SubScalarSels = float64SubScalarSels
	Float64SubByScalarSels = float64SubByScalarSels
}

func int8SubAvx2(xs, ys, rs []int8) []int8 {
	n := len(xs) / 16
	int8SubAvx2Asm(xs[:n*16], ys[:n*16], rs[:n*16])
	for i, j := n*16, len(xs); i < j; i++ {
		rs[i] = xs[i] - ys[i]
	}
	return rs
}

func int8SubAvx512(xs, ys, rs []int8) []int8 {
	n := len(xs) / 16
	int8SubAvx512Asm(xs[:n*16], ys[:n*16], rs[:n*16])
	for i, j := n*16, len(xs); i < j; i++ {
		rs[i] = xs[i] - ys[i]
	}
	return rs
}

func int8SubScalarAvx2(x int8, ys, rs []int8) []int8 {
	n := len(ys) / 16
	int8SubScalarAvx2Asm(x, ys[:n*16], rs[:n*16])
	for i, j := n*16, len(ys); i < j; i++ {
		rs[i] = x - ys[i]
	}
	return rs
}

func int8SubScalarAvx512(x int8, ys, rs []int8) []int8 {
	n := len(ys) / 16
	int8SubScalarAvx512Asm(x, ys[:n*16], rs[:n*16])
	for i, j := n*16, len(ys); i < j; i++ {
		rs[i] = x - ys[i]
	}
	return rs
}

func int8SubByScalarAvx2(x int8, ys, rs []int8) []int8 {
	n := len(ys) / 16
	int8SubByScalarAvx2Asm(x, ys[:n*16], rs[:n*16])
	for i, j := n*16, len(ys); i < j; i++ {
		rs[i] = ys[i] - x
	}
	return rs
}

func int8SubByScalarAvx512(x int8, ys, rs []int8) []int8 {
	n := len(ys) / 16
	int8SubByScalarAvx512Asm(x, ys[:n*16], rs[:n*16])
	for i, j := n*16, len(ys); i < j; i++ {
		rs[i] = ys[i] - x
	}
	return rs
}

func int16SubAvx2(xs, ys, rs []int16) []int16 {
	n := len(xs) / 8
	int16SubAvx2Asm(xs[:n*8], ys[:n*8], rs[:n*8])
	for i, j := n*8, len(xs); i < j; i++ {
		rs[i] = xs[i] - ys[i]
	}
	return rs
}

func int16SubAvx512(xs, ys, rs []int16) []int16 {
	n := len(xs) / 8
	int16SubAvx512Asm(xs[:n*8], ys[:n*8], rs[:n*8])
	for i, j := n*8, len(xs); i < j; i++ {
		rs[i] = xs[i] - ys[i]
	}
	return rs
}

func int16SubScalarAvx2(x int16, ys, rs []int16) []int16 {
	n := len(ys) / 8
	int16SubScalarAvx2Asm(x, ys[:n*8], rs[:n*8])
	for i, j := n*8, len(ys); i < j; i++ {
		rs[i] = x - ys[i]
	}
	return rs
}

func int16SubScalarAvx512(x int16, ys, rs []int16) []int16 {
	n := len(ys) / 8
	int16SubScalarAvx512Asm(x, ys[:n*8], rs[:n*8])
	for i, j := n*8, len(ys); i < j; i++ {
		rs[i] = x - ys[i]
	}
	return rs
}

func int16SubByScalarAvx2(x int16, ys, rs []int16) []int16 {
	n := len(ys) / 8
	int16SubByScalarAvx2Asm(x, ys[:n*8], rs[:n*8])
	for i, j := n*8, len(ys); i < j; i++ {
		rs[i] = ys[i] - x
	}
	return rs
}

func int16SubByScalarAvx512(x int16, ys, rs []int16) []int16 {
	n := len(ys) / 8
	int16SubByScalarAvx512Asm(x, ys[:n*8], rs[:n*8])
	for i, j := n*8, len(ys); i < j; i++ {
		rs[i] = ys[i] - x
	}
	return rs
}

func int32SubAvx2(xs, ys, rs []int32) []int32 {
	n := len(xs) / 4
	int32SubAvx2Asm(xs[:n*4], ys[:n*4], rs[:n*4])
	for i, j := n*4, len(xs); i < j; i++ {
		rs[i] = xs[i] - ys[i]
	}
	return rs
}

func int32SubAvx512(xs, ys, rs []int32) []int32 {
	n := len(xs) / 4
	int32SubAvx512Asm(xs[:n*4], ys[:n*4], rs[:n*4])
	for i, j := n*4, len(xs); i < j; i++ {
		rs[i] = xs[i] - ys[i]
	}
	return rs
}

func int32SubScalarAvx2(x int32, ys, rs []int32) []int32 {
	n := len(ys) / 4
	int32SubScalarAvx2Asm(x, ys[:n*4], rs[:n*4])
	for i, j := n*4, len(ys); i < j; i++ {
		rs[i] = x - ys[i]
	}
	return rs
}

func int32SubScalarAvx512(x int32, ys, rs []int32) []int32 {
	n := len(ys) / 4
	int32SubScalarAvx512Asm(x, ys[:n*4], rs[:n*4])
	for i, j := n*4, len(ys); i < j; i++ {
		rs[i] = x - ys[i]
	}
	return rs
}

func int32SubByScalarAvx2(x int32, ys, rs []int32) []int32 {
	n := len(ys) / 4
	int32SubByScalarAvx2Asm(x, ys[:n*4], rs[:n*4])
	for i, j := n*4, len(ys); i < j; i++ {
		rs[i] = ys[i] - x
	}
	return rs
}

func int32SubByScalarAvx512(x int32, ys, rs []int32) []int32 {
	n := len(ys) / 4
	int32SubByScalarAvx512Asm(x, ys[:n*4], rs[:n*4])
	for i, j := n*4, len(ys); i < j; i++ {
		rs[i] = ys[i] - x
	}
	return rs
}

func int64SubAvx2(xs, ys, rs []int64) []int64 {
	n := len(xs) / 2
	int64SubAvx2Asm(xs[:n*2], ys[:n*2], rs[:n*2])
	for i, j := n*2, len(xs); i < j; i++ {
		rs[i] = xs[i] - ys[i]
	}
	return rs
}

func int64SubAvx512(xs, ys, rs []int64) []int64 {
	n := len(xs) / 2
	int64SubAvx512Asm(xs[:n*2], ys[:n*2], rs[:n*2])
	for i, j := n*2, len(xs); i < j; i++ {
		rs[i] = xs[i] - ys[i]
	}
	return rs
}

func int64SubScalarAvx2(x int64, ys, rs []int64) []int64 {
	n := len(ys) / 2
	int64SubScalarAvx2Asm(x, ys[:n*2], rs[:n*2])
	for i, j := n*2, len(ys); i < j; i++ {
		rs[i] = x - ys[i]
	}
	return rs
}

func int64SubScalarAvx512(x int64, ys, rs []int64) []int64 {
	n := len(ys) / 2
	int64SubScalarAvx512Asm(x, ys[:n*2], rs[:n*2])
	for i, j := n*2, len(ys); i < j; i++ {
		rs[i] = x - ys[i]
	}
	return rs
}

func int64SubByScalarAvx2(x int64, ys, rs []int64) []int64 {
	n := len(ys) / 2
	int64SubByScalarAvx2Asm(x, ys[:n*2], rs[:n*2])
	for i, j := n*2, len(ys); i < j; i++ {
		rs[i] = ys[i] - x
	}
	return rs
}

func int64SubByScalarAvx512(x int64, ys, rs []int64) []int64 {
	n := len(ys) / 2
	int64SubByScalarAvx512Asm(x, ys[:n*2], rs[:n*2])
	for i, j := n*2, len(ys); i < j; i++ {
		rs[i] = ys[i] - x
	}
	return rs
}

func uint8SubAvx2(xs, ys, rs []uint8) []uint8 {
	n := len(xs) / 16
	uint8SubAvx2Asm(xs[:n*16], ys[:n*16], rs[:n*16])
	for i, j := n*16, len(xs); i < j; i++ {
		rs[i] = xs[i] - ys[i]
	}
	return rs
}

func uint8SubAvx512(xs, ys, rs []uint8) []uint8 {
	n := len(xs) / 16
	uint8SubAvx512Asm(xs[:n*16], ys[:n*16], rs[:n*16])
	for i, j := n*16, len(xs); i < j; i++ {
		rs[i] = xs[i] - ys[i]
	}
	return rs
}

func uint8SubScalarAvx2(x uint8, ys, rs []uint8) []uint8 {
	n := len(ys) / 16
	uint8SubScalarAvx2Asm(x, ys[:n*16], rs[:n*16])
	for i, j := n*16, len(ys); i < j; i++ {
		rs[i] = x - ys[i]
	}
	return rs
}

func uint8SubScalarAvx512(x uint8, ys, rs []uint8) []uint8 {
	n := len(ys) / 16
	uint8SubScalarAvx512Asm(x, ys[:n*16], rs[:n*16])
	for i, j := n*16, len(ys); i < j; i++ {
		rs[i] = x - ys[i]
	}
	return rs
}

func uint8SubByScalarAvx2(x uint8, ys, rs []uint8) []uint8 {
	n := len(ys) / 16
	uint8SubByScalarAvx2Asm(x, ys[:n*16], rs[:n*16])
	for i, j := n*16, len(ys); i < j; i++ {
		rs[i] = ys[i] - x
	}
	return rs
}

func uint8SubByScalarAvx512(x uint8, ys, rs []uint8) []uint8 {
	n := len(ys) / 16
	uint8SubByScalarAvx512Asm(x, ys[:n*16], rs[:n*16])
	for i, j := n*16, len(ys); i < j; i++ {
		rs[i] = ys[i] - x
	}
	return rs
}

func uint16SubAvx2(xs, ys, rs []uint16) []uint16 {
	n := len(xs) / 8
	uint16SubAvx2Asm(xs[:n*8], ys[:n*8], rs[:n*8])
	for i, j := n*8, len(xs); i < j; i++ {
		rs[i] = xs[i] - ys[i]
	}
	return rs
}

func uint16SubAvx512(xs, ys, rs []uint16) []uint16 {
	n := len(xs) / 8
	uint16SubAvx512Asm(xs[:n*8], ys[:n*8], rs[:n*8])
	for i, j := n*8, len(xs); i < j; i++ {
		rs[i] = xs[i] - ys[i]
	}
	return rs
}

func uint16SubScalarAvx2(x uint16, ys, rs []uint16) []uint16 {
	n := len(ys) / 8
	uint16SubScalarAvx2Asm(x, ys[:n*8], rs[:n*8])
	for i, j := n*8, len(ys); i < j; i++ {
		rs[i] = x - ys[i]
	}
	return rs
}

func uint16SubScalarAvx512(x uint16, ys, rs []uint16) []uint16 {
	n := len(ys) / 8
	uint16SubScalarAvx512Asm(x, ys[:n*8], rs[:n*8])
	for i, j := n*8, len(ys); i < j; i++ {
		rs[i] = x - ys[i]
	}
	return rs
}

func uint16SubByScalarAvx2(x uint16, ys, rs []uint16) []uint16 {
	n := len(ys) / 8
	uint16SubByScalarAvx2Asm(x, ys[:n*8], rs[:n*8])
	for i, j := n*8, len(ys); i < j; i++ {
		rs[i] = ys[i] - x
	}
	return rs
}

func uint16SubByScalarAvx512(x uint16, ys, rs []uint16) []uint16 {
	n := len(ys) / 8
	uint16SubByScalarAvx512Asm(x, ys[:n*8], rs[:n*8])
	for i, j := n*8, len(ys); i < j; i++ {
		rs[i] = ys[i] - x
	}
	return rs
}

func uint32SubAvx2(xs, ys, rs []uint32) []uint32 {
	n := len(xs) / 4
	uint32SubAvx2Asm(xs[:n*4], ys[:n*4], rs[:n*4])
	for i, j := n*4, len(xs); i < j; i++ {
		rs[i] = xs[i] - ys[i]
	}
	return rs
}

func uint32SubAvx512(xs, ys, rs []uint32) []uint32 {
	n := len(xs) / 4
	uint32SubAvx512Asm(xs[:n*4], ys[:n*4], rs[:n*4])
	for i, j := n*4, len(xs); i < j; i++ {
		rs[i] = xs[i] - ys[i]
	}
	return rs
}

func uint32SubScalarAvx2(x uint32, ys, rs []uint32) []uint32 {
	n := len(ys) / 4
	uint32SubScalarAvx2Asm(x, ys[:n*4], rs[:n*4])
	for i, j := n*4, len(ys); i < j; i++ {
		rs[i] = x - ys[i]
	}
	return rs
}

func uint32SubScalarAvx512(x uint32, ys, rs []uint32) []uint32 {
	n := len(ys) / 4
	uint32SubScalarAvx512Asm(x, ys[:n*4], rs[:n*4])
	for i, j := n*4, len(ys); i < j; i++ {
		rs[i] = x - ys[i]
	}
	return rs
}

func uint32SubByScalarAvx2(x uint32, ys, rs []uint32) []uint32 {
	n := len(ys) / 4
	uint32SubByScalarAvx2Asm(x, ys[:n*4], rs[:n*4])
	for i, j := n*4, len(ys); i < j; i++ {
		rs[i] = ys[i] - x
	}
	return rs
}

func uint32SubByScalarAvx512(x uint32, ys, rs []uint32) []uint32 {
	n := len(ys) / 4
	uint32SubByScalarAvx512Asm(x, ys[:n*4], rs[:n*4])
	for i, j := n*4, len(ys); i < j; i++ {
		rs[i] = ys[i] - x
	}
	return rs
}

func uint64SubAvx2(xs, ys, rs []uint64) []uint64 {
	n := len(xs) / 2
	uint64SubAvx2Asm(xs[:n*2], ys[:n*2], rs[:n*2])
	for i, j := n*2, len(xs); i < j; i++ {
		rs[i] = xs[i] - ys[i]
	}
	return rs
}

func uint64SubAvx512(xs, ys, rs []uint64) []uint64 {
	n := len(xs) / 2
	uint64SubAvx512Asm(xs[:n*2], ys[:n*2], rs[:n*2])
	for i, j := n*2, len(xs); i < j; i++ {
		rs[i] = xs[i] - ys[i]
	}
	return rs
}

func uint64SubScalarAvx2(x uint64, ys, rs []uint64) []uint64 {
	n := len(ys) / 2
	uint64SubScalarAvx2Asm(x, ys[:n*2], rs[:n*2])
	for i, j := n*2, len(ys); i < j; i++ {
		rs[i] = x - ys[i]
	}
	return rs
}

func uint64SubScalarAvx512(x uint64, ys, rs []uint64) []uint64 {
	n := len(ys) / 2
	uint64SubScalarAvx512Asm(x, ys[:n*2], rs[:n*2])
	for i, j := n*2, len(ys); i < j; i++ {
		rs[i] = x - ys[i]
	}
	return rs
}

func uint64SubByScalarAvx2(x uint64, ys, rs []uint64) []uint64 {
	n := len(ys) / 2
	uint64SubByScalarAvx2Asm(x, ys[:n*2], rs[:n*2])
	for i, j := n*2, len(ys); i < j; i++ {
		rs[i] = ys[i] - x
	}
	return rs
}

func uint64SubByScalarAvx512(x uint64, ys, rs []uint64) []uint64 {
	n := len(ys) / 2
	uint64SubByScalarAvx512Asm(x, ys[:n*2], rs[:n*2])
	for i, j := n*2, len(ys); i < j; i++ {
		rs[i] = ys[i] - x
	}
	return rs
}

func float32SubAvx2(xs, ys, rs []float32) []float32 {
	n := len(xs) / 4
	float32SubAvx2Asm(xs[:n*4], ys[:n*4], rs[:n*4])
	for i, j := n*4, len(xs); i < j; i++ {
		rs[i] = xs[i] - ys[i]
	}
	return rs
}

func float32SubAvx512(xs, ys, rs []float32) []float32 {
	n := len(xs) / 4
	float32SubAvx512Asm(xs[:n*4], ys[:n*4], rs[:n*4])
	for i, j := n*4, len(xs); i < j; i++ {
		rs[i] = xs[i] - ys[i]
	}
	return rs
}

func float32SubScalarAvx2(x float32, ys, rs []float32) []float32 {
	n := len(ys) / 4
	float32SubScalarAvx2Asm(x, ys[:n*4], rs[:n*4])
	for i, j := n*4, len(ys); i < j; i++ {
		rs[i] = x - ys[i]
	}
	return rs
}

func float32SubScalarAvx512(x float32, ys, rs []float32) []float32 {
	n := len(ys) / 4
	float32SubScalarAvx512Asm(x, ys[:n*4], rs[:n*4])
	for i, j := n*4, len(ys); i < j; i++ {
		rs[i] = x - ys[i]
	}
	return rs
}

func float32SubByScalarAvx2(x float32, ys, rs []float32) []float32 {
	n := len(ys) / 4
	float32SubByScalarAvx2Asm(x, ys[:n*4], rs[:n*4])
	for i, j := n*4, len(ys); i < j; i++ {
		rs[i] = ys[i] - x
	}
	return rs
}

func float32SubByScalarAvx512(x float32, ys, rs []float32) []float32 {
	n := len(ys) / 4
	float32SubByScalarAvx512Asm(x, ys[:n*4], rs[:n*4])
	for i, j := n*4, len(ys); i < j; i++ {
		rs[i] = ys[i] - x
	}
	return rs
}

func float64SubAvx2(xs, ys, rs []float64) []float64 {
	n := len(xs) / 2
	float64SubAvx2Asm(xs[:n*2], ys[:n*2], rs[:n*2])
	for i, j := n*2, len(xs); i < j; i++ {
		rs[i] = xs[i] - ys[i]
	}
	return rs
}

func float64SubAvx512(xs, ys, rs []float64) []float64 {
	n := len(xs) / 2
	float64SubAvx512Asm(xs[:n*2], ys[:n*2], rs[:n*2])
	for i, j := n*2, len(xs); i < j; i++ {
		rs[i] = xs[i] - ys[i]
	}
	return rs
}

func float64SubScalarAvx2(x float64, ys, rs []float64) []float64 {
	n := len(ys) / 2
	float64SubScalarAvx2Asm(x, ys[:n*2], rs[:n*2])
	for i, j := n*2, len(ys); i < j; i++ {
		rs[i] = x - ys[i]
	}
	return rs
}

func float64SubScalarAvx512(x float64, ys, rs []float64) []float64 {
	n := len(ys) / 2
	float64SubScalarAvx512Asm(x, ys[:n*2], rs[:n*2])
	for i, j := n*2, len(ys); i < j; i++ {
		rs[i] = x - ys[i]
	}
	return rs
}

func float64SubByScalarAvx2(x float64, ys, rs []float64) []float64 {
	n := len(ys) / 2
	float64SubByScalarAvx2Asm(x, ys[:n*2], rs[:n*2])
	for i, j := n*2, len(ys); i < j; i++ {
		rs[i] = ys[i] - x
	}
	return rs
}

func float64SubByScalarAvx512(x float64, ys, rs []float64) []float64 {
	n := len(ys) / 2
	float64SubByScalarAvx512Asm(x, ys[:n*2], rs[:n*2])
	for i, j := n*2, len(ys); i < j; i++ {
		rs[i] = ys[i] - x
	}
	return rs
}