package appregister

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"glintfed.org/internal/data"
)

func TestModel_VerifyCodeExists(t *testing.T) {
	client, cleanup, err := data.NewTestClient(t)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	model := NewModel(client)
	ctx := context.Background()

	email := "test@example.com"
	code := "123456"

	t.Run("Record exists and is recent", func(t *testing.T) {
		_, err := client.Ent.AppRegister.Create().
			SetEmail(email).
			SetVerifyCode(code).
			SetCreatedAt(time.Now().AddDate(0, 0, -10)).
			Save(ctx)
		assert.NoError(t, err)

		exists, err := model.VerifyCodeExists(ctx, email, code)
		assert.NoError(t, err)
		assert.True(t, exists)
	})

	t.Run("Record exists but is old (91 days ago)", func(t *testing.T) {
		oldEmail := "old@example.com"
		_, err := client.Ent.AppRegister.Create().
			SetEmail(oldEmail).
			SetVerifyCode(code).
			SetCreatedAt(time.Now().AddDate(0, 0, -91)).
			Save(ctx)
		assert.NoError(t, err)

		exists, err := model.VerifyCodeExists(ctx, oldEmail, code)
		assert.NoError(t, err)
		assert.False(t, exists)
	})

	t.Run("Record exists but code is different", func(t *testing.T) {
		exists, err := model.VerifyCodeExists(ctx, email, "wrong")
		assert.NoError(t, err)
		assert.False(t, exists)
	})

	t.Run("Record exists but email is different", func(t *testing.T) {
		exists, err := model.VerifyCodeExists(ctx, "wrong@example.com", code)
		assert.NoError(t, err)
		assert.False(t, exists)
	})

	t.Run("Record does not exist", func(t *testing.T) {
		exists, err := model.VerifyCodeExists(ctx, "nonexistent@example.com", "000000")
		assert.NoError(t, err)
		assert.False(t, exists)
	})
}

func TestModel_DeleteByEmail(t *testing.T) {
	client, cleanup, err := data.NewTestClient(t)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	model := NewModel(client)
	ctx := context.Background()

	email := "delete@example.com"
	_, err = client.Ent.AppRegister.Create().
		SetEmail(email).
		SetVerifyCode("123456").
		Save(ctx)
	assert.NoError(t, err)

	t.Run("Delete existing record", func(t *testing.T) {
		err := model.DeleteByEmail(ctx, email)
		assert.NoError(t, err)

		exists, err := model.VerifyCodeExists(ctx, email, "123456")
		assert.NoError(t, err)
		assert.False(t, exists)
	})

	t.Run("Delete non-existent record", func(t *testing.T) {
		err := model.DeleteByEmail(ctx, "nonexistent@example.com")
		assert.NoError(t, err)
	})
}
