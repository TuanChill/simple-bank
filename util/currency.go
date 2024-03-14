package util

// Constraints for all supported currencies
const (
	USD = "USD"
	CAD = "CAD"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}

	return false
}
