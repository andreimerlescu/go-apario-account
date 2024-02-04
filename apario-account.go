package go_apario_account

import (
	`time`

	ai `github.com/andreimerlescu/go-apario-identifier`

	`github.com/andreimerlescu/go-apario-account/account`
	`github.com/andreimerlescu/go-apario-account/email`
	`github.com/andreimerlescu/go-apario-account/gpg`
	`github.com/andreimerlescu/go-apario-account/wallet`
)

type AccountRequest struct {
	Username   string
	Email      email.Email
	Password   account.Keb
	GPG        gpg.GPG
	Reputation account.Reputation
}

// NewAccount requires a bound ai.Identifier first and an AccountRequest to generate a new account.Account
func NewAccount(identifier *ai.Identifier, newAccount *AccountRequest) account.Account {
	a := account.Account{
		Identifier: *identifier,
		Username:   newAccount.Username,
		Email:      newAccount.Email,
		Password:   newAccount.Password,
		GPG:        newAccount.GPG,
		Version: ai.Version{
			Major: 0,
			Minor: 0,
			Patch: 1,
		},
		Reputation: account.Reputation{},
	}

	r := account.Reputation{
		Account:                      &a,
		Blocks:                       []wallet.Block{},
		BlockSize:                    8192,
		Pending:                      []wallet.Reputation{},
		BlocksTotalPoints:            0,
		BlocksTotalPointsVerifiedAt:  time.Now().UTC(),
		PendingTotalPoints:           0,
		PendingTotalPointsVerifiedAt: time.Now().UTC(),
		PositiveActionCount:          0,
		NegativeActionCount:          0,
		Version: ai.Version{
			Major: 0,
			Minor: 0,
			Patch: 1,
		},
	}

	a.Reputation = r
	return a
}
