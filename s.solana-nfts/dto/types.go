package dto

// GetVendorPriceStatisticsByCollectionIDRequest ...
type GetVendorPriceStatisticsByCollectionIDRequest struct {
	CollectionID string
}

// VendorPriceStatistic ...
type VendorPriceStatistic struct {
	Name          string  `json:"name"`
	ID            string  `json:"id"`
	Price         float64 `json:"price"`
	LastSoldPrice float64 `json:"last_sold_price"`
	IsForSale     bool    `json:"is_for_sale"`
}

// GetVendorPriceStatisticsByCollectionIDResponse ...
type GetVendorPriceStatisticsByCollectionIDResponse struct {
	Stats       []*VendorPriceStatistic `json:"stats"`
	TotalListed int                     `json:"total_listed"`
}

// SolanartPriceStatistic ...
type SolanartPriceStatistic struct {
	ID            int     `json:"id,omitempty"`
	TokenAdd      string  `json:"token_add,omitempty"`
	Price         float64 `json:"price"`
	ForSale       int     `json:"for_sale,omitempty"`
	LinkImage     string  `json:"link_img,omitempty"`
	Name          string  `json:"name,omitempty"`
	EscrowAdd     string  `json:"escrowAdd,omitempty"`
	SellerAddress string  `json:"seller_address,omitempty"`
	Attributes    string  `json:"attributes,omitempty"`
	Skin          string  `json:"skin,omitempty"`
	Type          string  `json:"type,omitempty"`
	Ranking       float64 `json:"ranking,omitempty"`
	LastSoldPrice float64 `json:"lastSoldPrice,omitempty"`
}

// GetSolanartPriceStatisticsByCollectionIDResponse ...
type GetSolanartPriceStatisticsByCollectionIDResponse []*SolanartPriceStatistic

// Results used to generalize the response across vendors.
func (s GetSolanartPriceStatisticsByCollectionIDResponse) Results() GetSolanartPriceStatisticsByCollectionIDResponse {
	return s
}

// MagicEdenStatistic ...
type MagicEdenStatistic struct {
	ID    string `json:"title,omitempty"`
	Price string `json:"price"`
}

// GetMagicEdenPriceStatisticsByCollectionIDResponse ...
type GetMagicEdenPriceStatisticsByCollectionIDResponse struct {
	Results []*MagicEdenStatistic `json:"results"`
}
