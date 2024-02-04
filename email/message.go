package email

import (
	`fmt`
	`time`

	`github.com/andreimerlescu/go-apario-account/gpg`
)

type Message struct {
	To         Email
	From       Email
	Subject    string
	Body       string
	Encrypted  bool
	Signed     bool
	Checksum   string
	SentAt     time.Time
	ReceivedAt time.Time
	Read       bool
	ReadAt     time.Time
	ReadCount  int64
	Forwarded  bool
	Public     gpg.Key
}

// String pretty prints the Message by omitting any sensitive information beyond pure metadata
func (m *Message) String() string {
	return fmt.Sprintf("*Message [%T] To: %v ; From: %v ; Encrypted: %v ; Signed %v ; Checksum = %v",
		m, m.To, m.From, m.Encrypted, m.Signed, m.Checksum)
}
