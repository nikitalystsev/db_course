package impl

import (
	"SmartShopper-services/core/dto"
	"SmartShopper-services/core/models"
	"SmartShopper-services/errs"
	"SmartShopper-services/intf"
	"SmartShopper-services/intfRepo"
	"SmartShopper-services/pkg/transact"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type SaleProductService struct {
	saleProductRepo    intfRepo.ISaleProductRepo
	supplierRepo       intfRepo.ISupplierRepo
	productRepo        intfRepo.IProductRepo
	promotionRepo      intfRepo.IPromotionRepo
	shopRepo           intfRepo.IShopRepo
	transactionManager transact.ITransactionManager
}

func NewSaleProductService(
	saleProductRepo intfRepo.ISaleProductRepo,
	supplierRepo intfRepo.ISupplierRepo,
	productRepo intfRepo.IProductRepo,
	promotionRepo intfRepo.IPromotionRepo,
	shopRepo intfRepo.IShopRepo,
	transactionManager transact.ITransactionManager,
) intf.ISaleProductService {
	return &SaleProductService{
		saleProductRepo:    saleProductRepo,
		supplierRepo:       supplierRepo,
		productRepo:        productRepo,
		promotionRepo:      promotionRepo,
		shopRepo:           shopRepo,
		transactionManager: transactionManager,
	}
}

type SupplierIDs struct {
	retailerID     uuid.UUID
	distributorID  uuid.UUID
	manufacturerID uuid.UUID
}

func (sps *SaleProductService) Create(ctx context.Context, saleProduct *dto.NewSaleProductDTO) error {
	fmt.Println("call Create")
	return sps.transactionManager.Do(ctx, func(ctx context.Context) error {
		fmt.Println("call inner func")
		supplierIDs, err := sps.addSuppliersIfNotExists(ctx, saleProduct.Suppliers)
		if err != nil {
			return err
		}
		fmt.Println("Добавили поставщиков")
		productID, err := sps.addProductIfNotExists(ctx, supplierIDs, saleProduct.Product)
		if err != nil {
			fmt.Println("НЕ Добавили товар")
			return err
		}
		fmt.Println("Добавили товар")

		promotionID, err := sps.addPromotionIfNotExists(ctx, saleProduct.Promotion)
		if err != nil {
			return err
		}

		sale := &models.SaleProductModel{
			ID:          uuid.New(),
			ShopID:      saleProduct.ShopID,
			ProductID:   productID,
			PromotionID: promotionID,
			Price:       saleProduct.Price.Price,
			Currency:    saleProduct.Price.Currency,
			SettingDate: saleProduct.Price.SettingDate,
		}

		err = sps.saleProductRepo.Create(ctx, sale)
		if err != nil {
			return err
		}

		return nil
	})
}

func (sps *SaleProductService) GetByProductID(ctx context.Context, productID uuid.UUID) ([]*models.SaleProductModel, error) {
	sales, err := sps.saleProductRepo.GetByProductID(ctx, productID)
	if err != nil && !errors.Is(err, errs.ErrSaleProductDoesNotExists) {
		return nil, err
	}

	if errors.Is(err, errs.ErrSaleProductDoesNotExists) {
		return nil, errs.ErrSaleProductDoesNotExists
	}

	return sales, nil
}

func (sps *SaleProductService) GetByShopID(ctx context.Context, shopID uuid.UUID) ([]*models.SaleProductModel, error) {
	sales, err := sps.saleProductRepo.GetByShopID(ctx, shopID)
	if err != nil && !errors.Is(err, errs.ErrSaleProductDoesNotExists) {
		return nil, err
	}

	if errors.Is(err, errs.ErrSaleProductDoesNotExists) {
		return nil, errs.ErrSaleProductDoesNotExists
	}

	return sales, nil
}

func (sps *SaleProductService) GetByID(ctx context.Context, ID uuid.UUID) (*models.SaleProductModel, error) {
	sale, err := sps.saleProductRepo.GetByID(ctx, ID)
	if err != nil && !errors.Is(err, errs.ErrSaleProductDoesNotExists) {
		return nil, err
	}

	if errors.Is(err, errs.ErrSaleProductDoesNotExists) {
		return nil, errs.ErrSaleProductDoesNotExists
	}

	return sale, nil
}

func (sps *SaleProductService) Update(ctx context.Context, saleProduct *models.SaleProductModel) error {
	err := sps.saleProductRepo.Update(ctx, saleProduct)
	if err != nil {
		return err
	}

	return nil
}

func (sps *SaleProductService) addSuppliersIfNotExists(ctx context.Context, suppliers [3]*dto.SupplierDTO) (SupplierIDs, error) {
	fmt.Println("call addSuppliersIfNotExists")

	var (
		supplierIDs SupplierIDs
		err         error
	)
	supplierIDs.retailerID, err = sps.addRetailerIfNotExists(ctx, suppliers[0])
	if err != nil {
		return SupplierIDs{}, err
	}

	fmt.Println("Добавили ритейлера")

	supplierIDs.distributorID, err = sps.addDistributorIfNotExists(ctx, suppliers[1])
	if err != nil {
		return SupplierIDs{}, err
	}
	fmt.Println("Добавили 2")

	supplierIDs.manufacturerID, err = sps.addManufacturerIfNotExists(ctx, suppliers[2])
	if err != nil {
		return SupplierIDs{}, err
	}

	fmt.Println("Добавили 3")
	err = sps.addRetailerDistributorIfNotExists(ctx, supplierIDs.retailerID, supplierIDs.distributorID)
	if err != nil {
		return SupplierIDs{}, err
	}
	fmt.Println("Добавили связь 1")
	err = sps.addDistributorManufacturerIfNotExists(ctx, supplierIDs.distributorID, supplierIDs.manufacturerID)
	if err != nil {
		return SupplierIDs{}, err
	}

	fmt.Println("Добавили связь 2")
	return supplierIDs, nil
}

func (sps *SaleProductService) addRetailerIfNotExists(ctx context.Context, retailerDTO *dto.SupplierDTO) (uuid.UUID, error) {
	existingRetailer, err := sps.supplierRepo.GetRetailerByAddress(ctx, retailerDTO.Address)
	if err != nil && !errors.Is(err, errs.ErrRetailerDoesNotExists) {
		return uuid.Nil, err
	}

	if existingRetailer != nil {
		return existingRetailer.ID, nil
	}

	retailer := &models.SupplierModel{
		ID:                uuid.New(),
		Title:             retailerDTO.Title,
		Address:           retailerDTO.Address,
		PhoneNumber:       retailerDTO.PhoneNumber,
		FioRepresentative: retailerDTO.FioRepresentative,
	}

	if err = sps.supplierRepo.CreateRetailer(ctx, retailer); err != nil {
		return uuid.Nil, err
	}

	return retailer.ID, nil
}

func (sps *SaleProductService) addDistributorIfNotExists(ctx context.Context, distributorDTO *dto.SupplierDTO) (uuid.UUID, error) {
	existingRetailer, err := sps.supplierRepo.GetDistributorByAddress(ctx, distributorDTO.Address)
	if err != nil && !errors.Is(err, errs.ErrRetailerDoesNotExists) {
		return uuid.Nil, err
	}

	if existingRetailer != nil {
		return existingRetailer.ID, nil
	}

	distributor := &models.SupplierModel{
		ID:                uuid.New(),
		Title:             distributorDTO.Title,
		Address:           distributorDTO.Address,
		PhoneNumber:       distributorDTO.PhoneNumber,
		FioRepresentative: distributorDTO.FioRepresentative,
	}

	if err = sps.supplierRepo.CreateDistributor(ctx, distributor); err != nil {
		return uuid.Nil, err
	}

	return distributor.ID, nil
}

func (sps *SaleProductService) addManufacturerIfNotExists(ctx context.Context, manufacturerDTO *dto.SupplierDTO) (uuid.UUID, error) {
	existingRetailer, err := sps.supplierRepo.GetManufacturerByAddress(ctx, manufacturerDTO.Address)
	if err != nil && !errors.Is(err, errs.ErrRetailerDoesNotExists) {
		return uuid.Nil, err
	}

	if existingRetailer != nil {
		return existingRetailer.ID, nil
	}

	manufacturer := &models.SupplierModel{
		ID:                uuid.New(),
		Title:             manufacturerDTO.Title,
		Address:           manufacturerDTO.Address,
		PhoneNumber:       manufacturerDTO.PhoneNumber,
		FioRepresentative: manufacturerDTO.FioRepresentative,
	}

	if err = sps.supplierRepo.CreateManufacturer(ctx, manufacturer); err != nil {
		return uuid.Nil, err
	}

	return manufacturer.ID, nil
}

func (sps *SaleProductService) addRetailerDistributorIfNotExists(ctx context.Context, retailerID, distributorID uuid.UUID) error {
	fmt.Println("call addRetailerDistributorIfNotExists")
	isExist, err := sps.supplierRepo.IfExistsRetailerDistributor(ctx, retailerID, distributorID)
	if err != nil {
		fmt.Println("Ошибка выполнения запроса")
		return err
	}
	if isExist {
		fmt.Println("Существуют")
		return nil
	}

	fmt.Println("Создаем")

	err = sps.supplierRepo.CreateRetailerDistributor(ctx, retailerID, distributorID)
	if err != nil {
		fmt.Println("Хуево создали, блять")
		return err
	}
	return nil
}

func (sps *SaleProductService) addDistributorManufacturerIfNotExists(ctx context.Context, distributorID, manufacturerID uuid.UUID) error {
	isExist, err := sps.supplierRepo.IfExistsDistributorManufacturer(ctx, distributorID, manufacturerID)
	if err != nil {
		return err
	}
	if isExist {
		return nil
	}

	err = sps.supplierRepo.CreateDistributorManufacturer(ctx, distributorID, manufacturerID)
	if err != nil {
		return err
	}

	return nil
}

func (sps *SaleProductService) addProductIfNotExists(ctx context.Context, supplierIDs SupplierIDs, productDTO dto.ProductDTO) (uuid.UUID, error) {
	product := &models.ProductModel{
		ID:             uuid.New(),
		RetailerID:     supplierIDs.retailerID,
		DistributorID:  supplierIDs.distributorID,
		ManufacturerID: supplierIDs.manufacturerID,
		Name:           productDTO.Name,
		Categories:     productDTO.Categories,
		Brand:          productDTO.Brand,
		Compound:       productDTO.Compound,
		GrossMass:      productDTO.GrossMass,
		NetMass:        productDTO.NetMass,
		PackageType:    productDTO.PackageType,
	}

	err := sps.productRepo.Create(ctx, product)
	if err != nil {
		return uuid.Nil, err
	}

	return product.ID, nil
}

func (sps *SaleProductService) addPromotionIfNotExists(ctx context.Context, promotionDTO dto.PromotionDTO) (uuid.UUID, error) {
	if promotionDTO.Type == "" {
		return uuid.Nil, nil
	}

	promotion := &models.PromotionModel{
		ID:           uuid.New(),
		Type:         promotionDTO.Type,
		Description:  promotionDTO.Description,
		DiscountSize: promotionDTO.DiscountSize,
		StartDate:    promotionDTO.StartDate,
		EndDate:      promotionDTO.EndDate,
	}

	err := sps.promotionRepo.Create(ctx, promotion)
	if err != nil {
		return uuid.Nil, err
	}

	return promotion.ID, nil
}
