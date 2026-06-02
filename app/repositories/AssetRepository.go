package repositories

import (
	"edart-core/app/models"
	"edart-core/database"

	"gorm.io/gorm"
)

type AssetRepository struct {
	DB *gorm.DB
}

func NewAssetRepo() *AssetRepository{
	db := database.DB
	return &AssetRepository{DB: db}
}

func (r*AssetRepository) FindBySymbol(symbol string) (*string, error) {
	var category string
	err := r.DB.Model(&models.AssetListing{}).Select("category").Where(`"symbol"`, "= ?", symbol).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}