//go:build ios

package internal

func (e *Engine) getFileDescriptor() (int, error) {
	return int(e.mobileDep.FileDescriptor), nil
}
