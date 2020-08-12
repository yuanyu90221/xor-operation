package xor_op

func xorOperation(n int, start int) int {
	result := 0
	idx := 0
	for idx < n {
		result ^= (start + 2*idx)
		idx++
	}
	return result
}
