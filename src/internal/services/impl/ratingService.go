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

type RatingService struct {
	ratingRepo intfRepo.IRatingRepo
}

func NewRatingService(ratingRepo intfRepo.IRatingRepo) intf.IRatingService {
	return &RatingService{ratingRepo: ratingRepo}
}

func (rs *RatingService) Create(ctx context.Context, rating *models.RatingModel) error {
	if rating == nil {
		return errors.New("rating is nil")
	}

	existingRating, err := rs.ratingRepo.GetByUserAndSale(ctx, rating.UserID, rating.SaleProductID)
	if err != nil && !errors.Is(err, errs.ErrRatingDoesNotExists) {
		return err
	}

	if existingRating != nil {
		return errs.ErrRatingAlreadyExist
	}

	err = rs.ratingRepo.Create(ctx, rating)
	if err != nil {
		return err
	}

	return nil
}

func (rs *RatingService) GetByID(ctx context.Context, ID uuid.UUID) (*models.RatingModel, error) {
	existingRating, err := rs.ratingRepo.GetByID(ctx, ID)
	if err != nil && errors.Is(err, errs.ErrRatingDoesNotExists) {
		return nil, err
	}

	if existingRating != nil {
		return nil, errs.ErrRatingAlreadyExist
	}
	return existingRating, nil
}

func (rs *RatingService) DeleteByID(ctx context.Context, ID uuid.UUID) error {
	existingRating, err := rs.ratingRepo.GetByID(ctx, ID)
	if err != nil && errors.Is(err, errs.ErrRatingDoesNotExists) {
		return err
	}

	if existingRating != nil {
		return errs.ErrRatingAlreadyExist
	}

	err = rs.ratingRepo.DeleteByID(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
