package other


func ByteSliceExtend(in []byte, n int) (head, tail []byte) {
	total := len(in)+n
	if cap(in) >= total {
		head = in[:total]
	} else {
		head = make([]byte,total)
	}
	tail = head[len(in):]
	return
}
