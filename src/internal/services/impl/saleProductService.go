package impl

import (
	"SmartShopper-services/core/models"
	"SmartShopper-services/intf"
	"SmartShopper-services/intfRepo"
	"context"
)

type SaleProductService struct {
	saleProductRepo intfRepo.ISaleProductRepo
}

func NewSaleProductService(saleProductRepo intfRepo.ISaleProductRepo) intf.ISaleProductService {
	return &SaleProductService{saleProductRepo: saleProductRepo}
}

func (sps *SaleProductService) Create(ctx context.Context, saleProduct *models.UserModel) {

}
