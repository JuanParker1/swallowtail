package domain

import "time"

// VenueAccount holds metadata for a given exchange.
type VenueAccount struct {
	VenueAccountID string    `db:"venue_account_id"`
	VenueID        string    `db:"venue_id"`
	APIKey         string    `db:"api_key"`
	SecretKey      string    `db:"secret_key"`
	SubAccount     string    `db:"subaccount"`
	UserID         string    `db:"user_id"`
	Created        time.Time `db:"created"`
	Updated        time.Time `db:"updated"`
	IsActive       bool      `db:"is_active"`
	AccountAlias   string    `db:"account_alias"`
}
