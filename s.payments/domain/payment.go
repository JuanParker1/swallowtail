package domain

import "time"

// Payment ...
type Payment struct {
	UserID        string    `db:"user_id"`
	TransactionID string    `db:"transaction_id"`
	Timestamp     time.Time `db:"payment_timestamp"`
	AmountInUSDT  float64   `db:"amount_in_usdt"`
	AuditNote     string    `db:"audit_note"`
}
