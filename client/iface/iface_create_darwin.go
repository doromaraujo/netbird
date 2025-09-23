//go:build !ios

package iface

import (
	"time"

	"github.com/cenkalti/backoff/v4"
)

// Create creates a new Wireguard interface, sets a given IP and brings it up.
// Will reuse an existing one.
// this function is different on Android
func (w *WGIface) Create() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	backOff := &backoff.ExponentialBackOff{
		InitialInterval: 20 * time.Millisecond,
		MaxElapsedTime:  500 * time.Millisecond,
		Stop:            backoff.Stop,
		Clock:           backoff.SystemClock,
	}

	operation := func() error {
		cfgr, err := w.tun.Create()
		if err != nil {
			return err
		}
		w.configurer = cfgr
		return nil
	}

	return backoff.Retry(operation, backOff)
}
