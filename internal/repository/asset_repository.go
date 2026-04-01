package repository

import (
	"database/sql"

	"github.com/Werteryudf/sentinel-engine/internal/domain"
)

type AssetRepository struct {
	db *sql.DB 
}

func NewAssetRepository(db *sql.DB) (*AssetRepository){
	
	
	return &AssetRepository{db: db}
}

func (r *AssetRepository) GetAll() ([]domain.Asset, error){
    rows, err := r.db.Query("SELECT * FROM assets")
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	
	var assets []domain.Asset

	for rows.Next(){
		var asset domain.Asset
		err := rows.Scan(&asset.ID, &asset.Ticker, &asset.Name, &asset.AssetType, &asset.CreatedAt)
		if err != nil {
        	return nil, err
    	}
		assets = append(assets, asset)

		
	}	
	return assets, nil
}
 