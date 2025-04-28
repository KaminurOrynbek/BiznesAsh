package impl

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/dao"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	repo "github.com/KaminurOrynbek/BiznesAsh/internal/repository/interface"
	"github.com/jmoiron/sqlx"
)

type verificationRepositoryImpl struct {
	db *sqlx.DB
}

func NewVerificationRepository(db *sqlx.DB) repo.VerificationRepository {
	return &verificationRepositoryImpl{db: db}
}

func (r *verificationRepositoryImpl) SaveVerificationCode(ctx context.Context, verification *entity.Verification) error {
	query := `
		INSERT INTO verifications (user_id, email, code, expires_at, is_used)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.ExecContext(ctx, query,
		verification.UserID,
		verification.Email,
		verification.Code,
		verification.ExpiresAt,
		verification.IsUsed,
	)
	return err
}

func (r *verificationRepositoryImpl) VerifyCode(ctx context.Context, userID, code string) (bool, error) {
	query := `
		SELECT COUNT(*) FROM verifications 
		WHERE user_id = $1 AND code = $2 AND is_used = false AND expires_at > now()
	`
	var count int
	err := r.db.GetContext(ctx, &count, query, userID, code)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *verificationRepositoryImpl) GetVerificationCode(ctx context.Context, userID string) (*entity.Verification, error) {
	var v dao.Verification
	query := `
		SELECT id, user_id, email, code, expires_at, is_used FROM verifications WHERE user_id = $1
	`
	err := r.db.GetContext(ctx, &v, query, userID)
	if err != nil {
		return nil, err
	}

	return &entity.Verification{
		UserID:    v.UserID,
		Email:     v.Email,
		Code:      v.Code,
		ExpiresAt: v.ExpiresAt,
		IsUsed:    v.IsUsed,
	}, nil
}

func (r *verificationRepositoryImpl) UpdateVerificationStatus(ctx context.Context, userID string) error {
	query := `UPDATE verifications SET is_used = true WHERE user_id = $1`
	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}

func (r *verificationRepositoryImpl) UpdateVerificationCode(ctx context.Context, userID string, newCode string) error {
	query := `
		UPDATE verifications
		SET code = $1, expires_at = now() + interval '10 minutes', is_used = false
		WHERE user_id = $2
	`
	_, err := r.db.ExecContext(ctx, query, newCode, userID)
	return err
}
