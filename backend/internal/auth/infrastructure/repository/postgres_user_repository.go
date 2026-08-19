package repository

import (
	"context"
	"errors"

	"github.com/Wenev/WeDone/backend/internal/auth/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

var _ domain.UserRepository = (*PostgresUserRepository)(nil)

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (repo *PostgresUserRepository) Create(ctx context.Context, user *domain.User) error {
	model := FromDomain(user)
	result := repo.db.WithContext(ctx).Create(model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {

	var model UserModel
	err := repo.db.WithContext(ctx).Where("email = ?", email).First(&model).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	return ToDomain(&model), nil
}

func (repo *PostgresUserRepository) FindByGoogleID(ctx context.Context, googleID string) (*domain.User, error) {
	var model UserModel
	err := repo.db.WithContext(ctx).Where("google_id = ?", googleID).First(&model).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	return ToDomain(&model), nil
}

func (repo *PostgresUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var model UserModel
	err := repo.db.WithContext(ctx).Where("id = ?", id).First(&model).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	return ToDomain(&model), nil
}

func (repo *PostgresUserRepository) FindByInitial(ctx context.Context, initial string) (*domain.User, error) {
	var model UserModel
	err := repo.db.WithContext(ctx).Where("initial = ?", initial).First(&model).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	return ToDomain(&model), nil
}

func (repo *PostgresUserRepository) Update(ctx context.Context, user *domain.User) error {
	model := FromDomain(user)
	result := repo.db.WithContext(ctx).Save(model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *PostgresUserRepository) PromoteToManager(ctx context.Context, id uuid.UUID) error {
	result := repo.db.WithContext(ctx).Model(&UserModel{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"role": string(domain.RoleManager), 
			"manager_id": nil,
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}

func (repo *PostgresUserRepository) UpdateAdminStatus(ctx context.Context, id uuid.UUID, isAdmin bool) error {
	result := repo.db.WithContext(ctx).Model(&UserModel{}).Where("id = ?", id).Update("is_admin", isAdmin)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}

func (repo *PostgresUserRepository) UpdatePassword(ctx context.Context, id uuid.UUID, password string) error {
	result := repo.db.WithContext(ctx).Model(&UserModel{}).Where("id = ?", id).Update("password", password)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}
