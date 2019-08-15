// not good enough
package channnelclose

func SafeSend(ch chan T, value T) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()

	ch <- value

	return false
}

func SafeClose(ch chan T) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = false
		}
	}()
	close(ch)
	return true
}
