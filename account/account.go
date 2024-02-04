package account

import (
	ai `github.com/andreimerlescu/go-apario-identifier`

	`github.com/andreimerlescu/go-apario-account/email`
	`github.com/andreimerlescu/go-apario-account/gpg`
)

type Keb []rune

type Account struct {
	Identifier ai.Identifier
	Username   string
	Email      email.Email
	Password   Keb
	GPG        gpg.GPG
	Version    ai.Version
	Reputation Reputation
}

// Copy returns a struct Account with copied data from the *Account receiving the Copy()
func (a *Account) Copy() Account {
	return Account{
		Identifier: a.Identifier,
		Username:   a.Username,
		Email:      a.Email,
		Password:   a.Password,
		GPG:        a.GPG,
		Version:    a.Version,
		Reputation: a.Reputation,
	}
}

// VersionedIdentifier returns a new copy of ai.Identifier and uses the ai.Version from the Account and applies it to the new ai.Identifier as its version
func (a *Account) VersionedIdentifier() ai.Identifier {
	return ai.Identifier{
		Instance:  a.Identifier.Instance,
		Concierge: a.Identifier.Concierge,
		Table:     a.Identifier.Table,
		Year:      a.Identifier.Year,
		Fragment:  a.Identifier.Fragment,
		Version: &ai.Version{
			Major: a.Version.Major,
			Minor: a.Version.Minor,
			Patch: a.Version.Patch,
		},
	}
}
