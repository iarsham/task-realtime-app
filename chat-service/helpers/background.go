package helpers

import "go.uber.org/zap"

func Background(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				zap.S().Errorf("panic in background: %v", err)
			}
		}()
		f()
	}()
}
