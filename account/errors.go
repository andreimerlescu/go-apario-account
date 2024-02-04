package account

import (
	`errors`
)

var ErrBadChecksum = errors.New("unable to sha256sum the block")
