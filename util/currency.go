package util

const (
	USD = "USD"
	EUR = "EUR"
	GBP = "GBP"
	CAD = "CAD"
)

// IsSupportedCurrency checks if the currency is supported.
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, GBP, CAD:
		return true
	}
	return false
}
