package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"open-api/internal/models"
)

func (svc *Service) CreateAdmin(ctx context.Context, admin *models.Admin) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin.Password = string(hashedPassword)

	err = svc.repo.CreateAdmin(ctx, admin)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) GetAdminByEmail(ctx context.Context, email string) (models.Admin, error) {
	var admin models.Admin
	admin, err := svc.repo.GetAdminByEmail(ctx, email)
	if err != nil {
		return admin, err
	}

	return admin, nil
}

func (svc *Service) UpdateAdmin(ctx context.Context, admin *models.Admin) error {
	err := svc.repo.UpdateAdmin(ctx, admin)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) DeleteAdminByEmail(ctx context.Context, email string) error {
	err := svc.repo.DeleteAdminByEmail(ctx, email)
	if err != nil {
		return err
	}

	return nil
}
