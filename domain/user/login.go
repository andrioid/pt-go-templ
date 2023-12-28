package user

import (
	"app/internal/session"
	"context"
	"errors"
)

const SALT = "3wfUSLdaDPLe4q"

func Login(ctx context.Context, user, password string) error {
	// TODO: Replace with real login code
	// TODO: Fetch user row
	// TODO: Verify hash
	u, err := GetUser(ctx, user)
	if err != nil {
		return errors.New("login_failed")
	}

	session.Manager.Put(ctx, "user_id", u.ID)
	return nil
}

// https://pkg.go.dev/golang.org/x/crypto@v0.17.0/argon2#hdr-Argon2id
func validateLogin(u UserModel, in string) error {
	if u.hashType != "argon2id" {
		return errors.New("unsupported hash algorithm")
	}
	return nil
}

func Logout(ctx context.Context) error {
	if err := session.Manager.Destroy(ctx); err != nil {
		return err
	}
	return nil
}

func PasswordHash(in string) (string, error) {
	return in, nil
}
