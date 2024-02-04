package email

import (
	`time`

	`github.com/andreimerlescu/go-apario-account/gpg`
)

type Email struct {
	Current  string
	History  map[string]time.Time
	Verified bool
	gpg.GPG
}
