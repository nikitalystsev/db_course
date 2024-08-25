package handlers

import (
	"SmartShopper-services/core/dto"
	"SmartShopper-services/core/models"
	"SmartShopper-services/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (h *Handler) addSaleProductRating(c *gin.Context) {
	userIDStr, _, err := getReaderData(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var ratingDTO dto.RatingDTO
	if err = c.BindJSON(&ratingDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	rating := &models.RatingModel{
		ID:            uuid.New(),
		UserID:        userID,
		SaleProductID: ratingDTO.SaleProductID,
		Review:        ratingDTO.Review,
		Rating:        ratingDTO.Rating,
	}

	err = h.ratingService.Create(c.Request.Context(), rating)
	if err != nil && errors.Is(err, errs.ErrRatingAlreadyExist) {
		c.AbortWithStatusJSON(http.StatusConflict, err.Error())
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}
