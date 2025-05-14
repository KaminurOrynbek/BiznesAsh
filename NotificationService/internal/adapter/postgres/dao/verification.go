package dao

import (
	"context"
	"time"

	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/model"
	"github.com/jmoiron/sqlx"
)

type VerificationDAO struct {
	db *sqlx.DB
}

func NewVerificationDAO(db *sqlx.DB) *VerificationDAO {
	return &VerificationDAO{db: db}
}

func (dao *VerificationDAO) Save(ctx context.Context, v *model.Verification) error {
	query := `INSERT INTO verifications (user_id, email, code, expires_at, is_used)
	          VALUES (:user_id, :email, :code, :expires_at, :is_used)`

	insert := struct {
		UserID    string    `db:"user_id"`
		Email     string    `db:"email"`
		Code      string    `db:"code"`
		ExpiresAt time.Time `db:"expires_at"`
		IsUsed    bool      `db:"is_used"`
	}{
		UserID:    v.UserID,
		Email:     v.Email,
		Code:      v.Code,
		ExpiresAt: v.ExpiresAt,
		IsUsed:    v.IsUsed,
	}

	_, err := dao.db.NamedExecContext(ctx, query, insert)
	return err
}

func (dao *VerificationDAO) UpdateCode(ctx context.Context, userID, newCode string) error {
	query := `UPDATE verifications SET code = $1, expires_at = now() + interval '10 minutes', is_used = false WHERE user_id = $2`
	_, err := dao.db.ExecContext(ctx, query, newCode, userID)
	return err
}

func (dao *VerificationDAO) MarkUsed(ctx context.Context, userID string) error {
	query := `UPDATE verifications SET is_used = true WHERE user_id = $1`
	_, err := dao.db.ExecContext(ctx, query, userID)
	return err
}

func (dao *VerificationDAO) GetByUserID(ctx context.Context, userID string) (*model.Verification, error) {
	var v model.Verification
	query := `SELECT id, user_id, email, code, expires_at, is_used FROM verifications WHERE user_id = $1`
	err := dao.db.GetContext(ctx, &v, query, userID)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func (dao *VerificationDAO) CheckValid(ctx context.Context, userID, code string) (bool, error) {
	query := `SELECT COUNT(*) FROM verifications WHERE user_id = $1 AND code = $2 AND is_used = false AND expires_at > now()`
	var count int
	err := dao.db.GetContext(ctx, &count, query, userID, code)
	return count > 0, err
}
