package appregister

import (
	"context"
	"time"

	"glintfed.org/ent"
	"glintfed.org/ent/appregister"
	"glintfed.org/internal/data"
)

type Model struct {
	*ent.AppRegisterClient
}

func NewModel(client *data.Client) *Model {
	return &Model{
		AppRegisterClient: client.Ent.AppRegister,
	}
}

// VerifyCodeExists
//
//	SELECT EXISTS(
//	  SELECT 1 FROM app_registers
//	  WHERE email = ? AND verify_code = ? AND created_at > ?
//	)
func (m *Model) VerifyCodeExists(ctx context.Context, email, code string) (bool, error) {
	return m.Query().
		Where(
			appregister.Email(email),
			appregister.VerifyCode(code),
			appregister.CreatedAtGT(time.Now().AddDate(0, 0, -90)),
		).
		Exist(ctx)
}

// DeleteByEmail
//
//	DELETE FROM app_registers WHERE email = ?
func (m *Model) DeleteByEmail(ctx context.Context, email string) error {
	_, err := m.Delete().
		Where(appregister.Email(email)).
		Exec(ctx)
	return err
}
