package user

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"

	"glintfed.org/ent"
	"glintfed.org/internal/data"
)

type Model struct {
	*ent.UserClient
}

func NewModel(client *data.Client) *Model {
	return &Model{
		UserClient: client.Ent.User,
	}
}

// CreateUserParams holds the fields required to create a new user.
type CreateUserParams struct {
	Name            string
	Username        string
	Email           string
	Password        string // plaintext; hashed before storing
	AppRegisterIP   string
	RegisterSource  string
	EmailVerifiedAt time.Time
}

// Create
//
//	INSERT INTO users (name, username, email, password, app_register_ip, register_source, email_verified_at)
//	VALUES (?, ?, ?, ?, ?, ?, ?)
func (m *Model) Create(ctx context.Context, params CreateUserParams) (*ent.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return m.UserClient.Create().
		SetNillableName(nullableString(params.Name)).
		SetUsername(params.Username).
		SetEmail(params.Email).
		SetPassword(string(hashed)).
		SetNillableAppRegisterIP(nullableString(params.AppRegisterIP)).
		SetNillableRegisterSource(nullableString(params.RegisterSource)).
		SetEmailVerifiedAt(params.EmailVerifiedAt).
		Save(ctx)
}

func nullableString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
