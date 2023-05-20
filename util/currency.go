package util

// constant for all ssupported currencies

const (
	USD = "USD"
	EUR = "EUR"
	NGN = "NGN"
	CAD = "CAD"
)

func IsSupporteCurrency(currency string) bool {
	switch currency {
	case USD, EUR, NGN, CAD:
		return true
	}
	return false
}