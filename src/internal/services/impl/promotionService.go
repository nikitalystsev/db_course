package impl

import (
	"SmartShopper-services/core/models"
	"SmartShopper-services/errs"
	"SmartShopper-services/intf"
	"SmartShopper-services/intfRepo"
	"context"
	"errors"
	"github.com/google/uuid"
)

type PromotionService struct {
	promotionRepo intfRepo.IPromotionRepo
}

func NewPromotionService(promotionRepo intfRepo.IPromotionRepo) intf.IPromotionService {
	return &PromotionService{promotionRepo: promotionRepo}
}

func (ps *PromotionService) Create(ctx context.Context, promotion *models.PromotionModel) error {
	err := ps.promotionRepo.Create(ctx, promotion)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PromotionService) GetByID(ctx context.Context, ID uuid.UUID) (*models.PromotionModel, error) {
	existingPromotion, err := ps.promotionRepo.GetByID(ctx, ID)
	if err != nil && errors.Is(err, errs.ErrPromotionDoesNotExists) {
		return nil, err
	}

	if existingPromotion == nil {
		return nil, errs.ErrPromotionDoesNotExists
	}

	return existingPromotion, nil
}
