package go_apario_account

import (
	`os`
	`testing`
	`time`

	ai `github.com/andreimerlescu/go-apario-identifier`

	`github.com/andreimerlescu/go-apario-account/account`
	`github.com/andreimerlescu/go-apario-account/email`
	`github.com/andreimerlescu/go-apario-account/gpg`
)

func TestAparioAccount(t *testing.T) {
	databasePath, mkErr := os.MkdirTemp("", "documents.db")
	if mkErr != nil {
		t.Errorf("os.MkdirTemp received mkErr %v", mkErr)
		return
	}
	id, err := ai.NewIdentifier(databasePath, 6, 3, 9)
	if err != nil {
		t.Errorf("ai.NewIdentifier received err %v", err)
		return
	}
	req := AccountRequest{
		Username: "Webmaster",
		Email: email.Email{
			Current:  "webmaster@example.com",
			History:  map[string]time.Time{},
			Verified: false,
			GPG:      gpg.GPG{},
		},
		Password:   account.Keb("Password123$"),
		GPG:        gpg.GPG{},
		Reputation: account.Reputation{},
	}
	a := NewAccount(id, &req)
	vid := a.VersionedIdentifier()
	if vid.Version.Patch != a.Version.Patch {
		t.Errorf("vid.Version.Patch != a.Version.Patch triggered error in test ; expected == ; got != ")
		return
	}
}
