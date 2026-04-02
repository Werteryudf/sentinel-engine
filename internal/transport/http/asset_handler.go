package http

import(
	
	"net/http"
	"github.com/Werteryudf/sentinel-engine/internal/repository"
	"encoding/json"
)


type AssetHandler struct {
    repo *repository.AssetRepository
}


func NewAssetHandler(repo *repository.AssetRepository) (*AssetHandler){
	return &AssetHandler{repo: repo}
}


func (h *AssetHandler) GetAll(w http.ResponseWriter, r *http.Request) {
    assets, err := h.repo.GetAll()
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(assets)
	
}