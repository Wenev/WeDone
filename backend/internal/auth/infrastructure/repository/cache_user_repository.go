package repository

import (
	"context"
	"encoding/json"
	"log/slog"
	"math/rand"
	"time"

	"github.com/Wenev/WeDone/backend/internal/auth/application/utility"
	"github.com/Wenev/WeDone/backend/internal/auth/domain"
	"github.com/google/uuid"
	"golang.org/x/sync/singleflight"
)

const (
	baseTTL   = 15 * time.Minute
	jitterMax = 120 * time.Second
)

func userIDKey(id uuid.UUID) string {
	return "auth:user:id:" + id.String()
}

func userEmailKey(email string) string {
	return "auth:user:email:" + email
}

func userInitialKey(initial string) string {
	return "auth:user:initial:" + initial
}

func userGoogleIDKey(gid string) string {
	return "auth:user:google_id:" + gid
}

func ttlWithJitter() time.Duration {
	jitter := time.Duration(rand.Int63n(int64(jitterMax)))
	return baseTTL + jitter
}

type CacheUserRepository struct {
	userRepo domain.UserRepository
	cache    utility.UserCachePort
	sf       singleflight.Group
}

func (repo *CacheUserRepository) setCache(ctx context.Context, user *domain.User) error {
	raw, err := json.Marshal(user)
	if err != nil {
		slog.Warn("cache: marshal failed", "err", err)
		return err
	}
	ttl := ttlWithJitter()
	blob := string(raw)

	if err := repo.cache.Set(ctx, userIDKey(user.ID), blob, ttl); err != nil {
		return err
	}

	idStr := user.ID.String()

	_ = repo.cache.Set(ctx, userEmailKey(user.Email), idStr, ttl)
	_ = repo.cache.Set(ctx, userInitialKey(user.Initial), idStr, ttl)
	if user.GoogleID != "" {
		_ = repo.cache.Set(ctx, userGoogleIDKey(user.GoogleID), idStr, ttl)
	}

	return nil
}

func (repo *CacheUserRepository) invalidate(ctx context.Context, id uuid.UUID, email, initial, googleID string) error {
	keys := []string{userIDKey(id)}
	if email != "" {
		keys = append(keys, userEmailKey(email))
	}
	if initial != "" {
		keys = append(keys, userInitialKey(initial))
	}
	if googleID != "" {
		keys = append(keys, userGoogleIDKey(googleID))
	}
	if err := repo.cache.Delete(ctx, keys...); err != nil {
		return err
	}
	return nil
}

func (repo *CacheUserRepository) Create(ctx context.Context, user *domain.User) error {
	return repo.userRepo.Create(ctx, user)
}

func (repo *CacheUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	if idStr, err := repo.cache.Get(ctx, userEmailKey(email)); err == nil {
		if id, parseErr := uuid.Parse(idStr); parseErr == nil {
			return repo.FindByID(ctx, id)
		}
	}
	user, err := repo.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	repo.setCache(ctx, user)
	return user, nil
}

func (repo *CacheUserRepository) FindByGoogleID(ctx context.Context, googleID string) (*domain.User, error) {
	if idStr, err := repo.cache.Get(ctx, userGoogleIDKey(googleID)); err == nil {
		if id, parseErr := uuid.Parse(idStr); parseErr == nil {
			return repo.FindByID(ctx, id)
		}
	}
	user, err := repo.userRepo.FindByGoogleID(ctx, googleID)
	if err != nil {
		return nil, err
	}
	repo.setCache(ctx, user)
	return user, nil
}

func (repo *CacheUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	key := userIDKey(id)
	if raw, err := repo.cache.Get(ctx, key); err == nil {
		var user domain.User
		if jsonErr := json.Unmarshal([]byte(raw), &user); jsonErr == nil {
			return &user, nil
		}
	}
	val, err, _ := repo.sf.Do(key, func() (interface{}, error) {
		user, dbErr := repo.userRepo.FindByID(ctx, id)
		if dbErr != nil {
			return nil, dbErr
		}
		repo.setCache(ctx, user)
		return user, nil
	})
	if err != nil {
		return nil, err
	}
	return val.(*domain.User), nil
}

func (repo *CacheUserRepository) FindByInitial(ctx context.Context, initial string) (*domain.User, error) {
	if idStr, err := repo.cache.Get(ctx, userInitialKey(initial)); err == nil {
		if id, parseErr := uuid.Parse(idStr); parseErr == nil {
			return repo.FindByID(ctx, id)
		}
	}
	user, err := repo.userRepo.FindByInitial(ctx, initial)
	if err != nil {
		return nil, err
	}
	repo.setCache(ctx, user)
	return user, nil
}

func (repo *CacheUserRepository) Update(ctx context.Context, user *domain.User) error {
	if err := repo.userRepo.Update(ctx, user); err != nil {
		return err
	}
	repo.invalidate(ctx, user.ID, user.Email, user.Initial, user.GoogleID)
	return nil
}

func (repo *CacheUserRepository) UpdateAdminStatus(ctx context.Context, id uuid.UUID, isAdmin bool) error {
	if err := repo.userRepo.UpdateAdminStatus(ctx, id, isAdmin); err != nil {
		return err
	}
	repo.invalidate(ctx, id, "", "", "")
	return nil
}

func (repo *CacheUserRepository) UpdatePassword(ctx context.Context, id uuid.UUID, password string) error {
	if err := repo.userRepo.UpdatePassword(ctx, id, password); err != nil {
		return err
	}
	repo.invalidate(ctx, id, "", "", "")
	return nil
}
