package impl

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/dao"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/model"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	repo "github.com/KaminurOrynbek/BiznesAsh/internal/repository/interface"
)

type verificationRepositoryImpl struct {
	dao *dao.VerificationDAO
}

func NewVerificationRepository(dao *dao.VerificationDAO) repo.VerificationRepository {
	return &verificationRepositoryImpl{dao: dao}
}

func (r *verificationRepositoryImpl) SaveVerificationCode(ctx context.Context, v *entity.Verification) error {
	return r.dao.Save(ctx, model.FromEntityVerification(v))
}

func (r *verificationRepositoryImpl) VerifyCode(ctx context.Context, userID, code string) (bool, error) {
	return r.dao.CheckValid(ctx, userID, code)
}

func (r *verificationRepositoryImpl) GetVerificationCode(ctx context.Context, userID string) (*entity.Verification, error) {
	m, err := r.dao.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return m.ToEntity(), nil
}

func (r *verificationRepositoryImpl) UpdateVerificationStatus(ctx context.Context, userID string) error {
	return r.dao.MarkUsed(ctx, userID)
}

func (r *verificationRepositoryImpl) UpdateVerificationCode(ctx context.Context, userID string, newCode string) error {
	return r.dao.UpdateCode(ctx, userID, newCode)
}
