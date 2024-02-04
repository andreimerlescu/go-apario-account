# Apario Account

This package is used for creating an account data type that has an associated Reputation blockchain connected 
with GPG and Email support. This package does not deliver Email or encrypt/decrypt GPG, but rather will utilize
the `$PATH` value of `gpg` to perform the encrypt/decrypt of the given messages AND potentially SIGN or ENCRYPT a 
body message of a possible email communication.

Within the Apario Account, the expectation is the `Account` struct will be saved inside the Apario Identifier Valet,
inside the database called `accounts.db`. From there, Reputation is tracked/maintained using a blockchain hashing
technique for validation and consistency as the blocks increase in size; but is designed to not allow Reputation 
history to bloat the application and ruin the experience. The blockchain allows 8,192 Reputation transactions to be 
stored/expressed/accessed in a slice called `Pending` and every multiple of 8,192 Reputation transactions are 
hashed and flattened as each Block is added. 

Inside this `accounts.db` directory, the identifier of the account will be used to store the actual `account.json`
file through the Valet's version control system. Meaning; within the `accounts.db` and after the fibonacci directory
split identifier sits another folder called `account.json` which contains a root sub folder (always) called v0, 
another subdirectory (always) called 0 and then finally a last subdirectory called 1 (always 1 never 0 in this last 
octet). Inside this subdir sits the actual `account.json` at the given version number. For an example identifier of 
20241 (1st of 2024), the `account.json` will be located at (assuming the valet is stored at /valet): 

```shell
/valet/accounts.db/2024/1/account.json/v0/0/1/account.json
```

Then; emails too... and GPG public keys... and other data included copies of contributed works to the instance.

Many more details will be added to this README as the project progresses further into the future.

## Installation 

```shell
go get -u github.com/andreimerlescu/go-apario-account
```

## Usage

```go
package main

import (
	
	aa `github.com/andreimerlescu/go-apario-account`
	gpg `github.com/andreimerlescu/go-apario-account/gpg`
	email `github.com/andreimerlescu/go-apario-account/email`
	wallet `github.com/andreimerlescu/go-apario-account/wallet`
	ai `github.com/andreimerlescu/go-apario-identifier`
)

func main() {
	databasePath := "/valet/documents.db"
	id, err := ai.NewIdentifier(databasePath, 6, 3, 9)
	if err != nil {
		panic(err)
    }
	req := aa.AccountRequest{
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
	a := aa.NewAccount(id, &req)
	// do something with the new account
}

```

There is a test file that can be executed with `go test ./...` on this directory.

## Contributing / Developing

When downloading this repo, its important for you NOT to work off the main branch or even within this repository
itself. Instead, you must fork the repo to your account, commit your changes/tests there, and when ready, submit a 
pull request and your contribution will be considered for integration into the larger code-base. 

```shell
git clone git@github.com:andreimerlescu/go-apario-account.git
cd go-apario-account
git fetch --all
git checkout main
git pull origin -u main
git checkout -b development
go test ./...
```

For contributions from developers outside of Project Apario LLC, it is required that you have 100% code coverage in
your test files. Omitting test files or refusing to provide documentation/implementation/usage tests for a given
routine may result in that routine being the root cause for a pull request rejection. All pull requests from 
non-Project Apario developers must have 100% code coverage in there tests for any Go code submitted.

## Notes

This package is in a very early stage at the moment and should be considered in-development and not-safe-for-production
use yet. 

## License

This package is released under the GPL-3 license and was created by Project Apario LLC. 

## Intended Use Case

This package will be the account/wallet/blockchain of the decentralized Project Apario ecosystem. This is the first of
many, many commits of code that will happen with this idea. This package will be integrated into apario-reader and 
apario-writer plus other apario- related packages corresponding to the wider ecosystem of the Apario Network.
