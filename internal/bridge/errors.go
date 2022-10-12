package bridge

import "errors"

var (
	ErrVaultInsecure = errors.New("the vault is insecure")
	ErrVaultCorrupt  = errors.New("the vault is corrupt")

	ErrServeIMAP    = errors.New("failed to serve IMAP")
	ErrServeSMTP    = errors.New("failed to serve SMTP")
	ErrWatchUpdates = errors.New("failed to watch for updates")

	ErrNoSuchUser          = errors.New("no such user")
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrUserAlreadyLoggedIn = errors.New("the user is already logged in")
	ErrNotImplemented      = errors.New("not implemented")

	ErrSizeTooLarge = errors.New("file is too big")
)