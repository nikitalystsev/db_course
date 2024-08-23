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
	println("call addSaleProductRating")
	userIDStr, _, err := getReaderData(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		println("не получили юзера")
		return
	}
	println("получили юзера")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		println("не спарсили id юзера")
		return
	}
	println("спарсили id юзера")
	var ratingDTO dto.RatingDTO
	if err = c.BindJSON(&ratingDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		println("не спарсили dto оценки")
		return
	}

	println("спарсили dto оценки")

	rating := &models.RatingModel{
		ID:            uuid.New(),
		UserID:        userID,
		SaleProductID: ratingDTO.SaleProductID,
		Review:        ratingDTO.Review,
		Rating:        ratingDTO.Rating,
	}

	if rating == nil {
		println("чет указатель на модель рейтинга nil")
	} else {
		println("так, указатель на модель рейтинга в порядке")
	}

	err = h.ratingService.Create(c.Request.Context(), rating)
	if err != nil && errors.Is(err, errs.ErrRatingAlreadyExist) {
		c.AbortWithStatusJSON(http.StatusConflict, err.Error())
		println("уже есть оценка")
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		println("что то под капотом не то")
		return
	}

	println("все супер, добавили")

	c.Status(http.StatusCreated)
}
