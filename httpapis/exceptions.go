package httpapis

import (
	"errors"
)

var ErrClientInited = errors.New("client already inited")
var ErrClientNotInit = errors.New("client not init")
var ErrHttpQueryNotOK = errors.New("http query not ok")
