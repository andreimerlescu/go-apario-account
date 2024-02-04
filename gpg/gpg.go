package gpg

import (
	`time`
)

type Key []byte

type GPG struct {
	Public    Key
	Private   Key
	ID        []byte
	CreatedAt time.Time
}
