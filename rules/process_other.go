// +build !darwin,!linux

package rules

import (
	C "github.com/Dreamacro/clash/constant"
)

func NewProcess(process string, adapter string) (C.Rule, error) {
	return nil, ErrPlatformNotSupport
}
