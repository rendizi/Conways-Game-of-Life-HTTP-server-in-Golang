package main

func intopos(pos, n int) int {
	for pos <= -n {
		pos += n
	}
	for pos >= n {
		pos -= n
	}
	return pos
}

func Rotate(data []int, pos int) []int {
	n := len(data)
	if pos == 0 || n == 0 {
		return data
	}
	pos = intopos(pos, n)

	if pos > 0 {
		return append(data[pos+1:], data[:pos+1]...)
	}

	pos = -pos
	return append(data[pos:], data[:pos]...)
}
