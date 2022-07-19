package logic1

import "github.com/feditools/go-lib/fedihelper"

func (l *Logic) Transport() *fedihelper.Transport {
	return l.transport
}
