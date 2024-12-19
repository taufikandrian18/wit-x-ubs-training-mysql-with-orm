package utility

import "gitlab.com/wit-id/go-orm-mysql/common/constants"

func PaymentGateway(pg string) bool {
	switch pg {
	case constants.PaymentGatewayFlip:
		return true
	case constants.PaymentGatewayXendit:
		return true
	}

	return false
}
