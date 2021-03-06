package templates

func init() {
	registerTemplate(PortfolioSheetType, portfolioTemplate{})
}

type portfolioTemplate struct{}

func (p portfolioTemplate) ID() SheetType {
	return PortfolioSheetType
}

func (p portfolioTemplate) RowRange() string {
	return "A2:X7"
}

func (p portfolioTemplate) Values() [][]interface{} {
	return [][]interface{}{
		{"", "", "Total PNL", "0"},
		{"", "", "Total Worth", "0"},
		{"", "", "Asset Pair", "USDT", "", "", "", "", "", "", "", "", "", "", "", "", "Historical"},
		{},
		{},
		{"Date Bought", "Size", "Coin", "Asset Pair", "Avg. Entry", "Bought", "Amount", "Current Price", "Current Value", "PNL Value", "PNL %", "Target (Optional)", "Location", "Staked", "Notes", "", "Coin", "Asset Pair", "Bought For", "Sold", "Sold %", "PNL Value", "Notes"},
	}
}
