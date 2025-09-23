//go:build !android && !ios

package internal

import "fmt"

func (e *Engine) getFileDescriptor() (int, error) {
	return -1, fmt.Errorf("this function has not implemented on this platform")
}
