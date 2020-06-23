package gocrush

const (
	maxValue      = int64(0xFFFFFFFF)
	crushHashSeed = int64(1315423911)
)

type triple struct {
	a int64
	b int64
	c int64
}

func hash1(a int64) int64 {
	var hash = xor(crushHashSeed, a)
	var x = int64(231232)
	var y = int64(1232)
	b := a
	b, x, hash = hashMix(b, x, hash)
	y, a, hash = hashMix(y, a, hash)
	return hash
}

func hash2(a, b int64) int64 {
	var hash = xor(xor(crushHashSeed, a), b)
	var x = int64(231232)
	var y = int64(1232)
	a, b, hash = hashMix(a, b, hash)
	x, a, hash = hashMix(x, a, hash)
	b, y, hash = hashMix(b, y, hash)
	return hash
}

func hash3(a, b, c int64) int64 {
	var hash = xor(xor(xor(crushHashSeed, a), b), c)
	var x = int64(231232)
	var y = int64(1232)
	a, b, hash = hashMix(a, b, hash)
	c, x, hash = hashMix(c, x, hash)
	y, a, hash = hashMix(y, a, hash)
	b, x, hash = hashMix(b, x, hash)
	_, _, hash = hashMix(y, c, hash)

	return hash
}

func hash4(a, b, c, d int64) int64 {
	var hash = xor(xor(xor(xor(crushHashSeed, a), b), c), d)
	var x = int64(231232)
	var y = int64(1232)
	a, b, hash = hashMix(a, b, hash)
	c, x, hash = hashMix(c, d, hash)
	c, x, hash = hashMix(a, x, hash)
	y, a, hash = hashMix(y, b, hash)
	b, x, hash = hashMix(c, x, hash)
	_, _, hash = hashMix(y, d, hash)

	return hash
}

func hashMix(a, b, c int64) (int64, int64, int64) {
	a = subtract(a, b)
	a = subtract(a, c)
	a = xor(a, c>>13)
	b = subtract(b, c)
	b = subtract(b, a)
	b = xor(b, leftShift(a, 8))
	c = subtract(c, a)
	c = subtract(c, b)
	c = xor(c, (b >> 13))
	a = subtract(a, b)
	a = subtract(a, c)
	a = xor(a, (c >> 12))
	b = subtract(b, c)
	b = subtract(b, a)
	b = xor(b, leftShift(a, 16))
	c = subtract(c, a)
	c = subtract(c, b)
	c = xor(c, (b >> 5))
	a = subtract(a, b)
	a = subtract(a, c)
	a = xor(a, (c >> 3))
	b = subtract(b, c)
	b = subtract(b, a)
	b = xor(b, leftShift(a, 10))
	c = subtract(c, a)
	c = subtract(c, b)
	c = xor(c, (b >> 15))
	return a, b, c
}

func subtract(val, subtract int64) int64 {
	return (val - subtract) & maxValue
}

func leftShift(val, shift int64) int64 {
	return (val << uint64(shift)) & maxValue
}

func xor(val, xor int64) int64 {
	return (val ^ xor) & maxValue
}
