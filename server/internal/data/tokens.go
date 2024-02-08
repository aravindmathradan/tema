package data

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base32"
	"fmt"
	"time"

	"github.com/aravindmathradan/tema/internal/validator"
)

const (
	ScopeActivation     = "activation"
	ScopeAuthentication = "authentication"
	ScopePasswordReset  = "password-reset"
	ScopeRefresh        = "refresh"
)

type Token struct {
	Plaintext string    `json:"token"`
	Hash      []byte    `json:"-"`
	UserID    int64     `json:"-"`
	Expiry    time.Time `json:"expiry"`
	Scope     string    `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func generateToken(userID int64, ttl time.Duration, scope string) (*Token, error) {
	token := &Token{
		UserID:    userID,
		Expiry:    time.Now().Add(ttl),
		Scope:     scope,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	randomBytes := make([]byte, 16)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]

	return token, nil
}

func ValidateTokenPlaintext(v *validator.Validator, tokenPlaintext string) {
	v.Check(validator.NotBlank(tokenPlaintext), "token", validator.EBLANKFIELD, "must be provided")
	v.Check(len(tokenPlaintext) == 26, "token", validator.EINVALIDTOKEN, "must be 26 bytes long")
}

func ValidateTokenScope(v *validator.Validator, scope string) {
	v.Check(validator.NotBlank(scope), "scope", validator.EBLANKFIELD, "must be provided")
	v.Check(
		validator.PermittedValue[string](scope, ScopeAuthentication, ScopeRefresh),
		"scope",
		validator.EVALUENOTPERMITTED,
		fmt.Sprintf("must be %s or %s", ScopeAuthentication, ScopeRefresh),
	)
}

type TokenModel struct {
	DB *sql.DB
}

func (m TokenModel) New(userID int64, ttl time.Duration, scope string) (*Token, error) {
	token, err := generateToken(userID, ttl, scope)
	if err != nil {
		return nil, err
	}

	err = m.Insert(token)
	return token, err
}

func (m TokenModel) Insert(token *Token) error {
	query := `
		INSERT INTO tokens (hash, user_id, expiry, scope, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

	args := []any{token.Hash, token.UserID, token.Expiry, token.Scope, token.CreatedAt, token.UpdatedAt}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, args...)
	return err
}

func (m TokenModel) DeleteAllForUser(scope string, userID int64) error {
	query := `
		DELETE FROM tokens
		WHERE scope = $1 AND user_id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, scope, userID)
	return err
}
