package echohttp

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"gitlab.com/wit-id/go-orm-mysql/common/constants"
	"gitlab.com/wit-id/go-orm-mysql/common/httpservice"
	"gitlab.com/wit-id/go-orm-mysql/toolkit/config"
)

func handleEchoError(_ config.KVStore) echo.HTTPErrorHandler {
	return func(err error, ctx echo.Context) {
		var echoError *echo.HTTPError

		// if *echo.HTTPError, let echokit middleware handles it
		if errors.As(err, &echoError) {
			return
		}

		statusCode := http.StatusInternalServerError
		// message := "mohon maaf, terjadi kesalahan pada server"
		var message interface{}

		message = ErrInternalServerErrror

		switch {
		case errors.Is(err, httpservice.ErrBadRequest):
			statusCode = http.StatusBadRequest
			message = ErrBadRequest

		case errors.Is(err, httpservice.ErrInvalidAppKey):
			statusCode = http.StatusUnauthorized
			message = ErrInvalidAppKey

		case errors.Is(err, httpservice.ErrUnknownSource):
			statusCode = http.StatusUnprocessableEntity
			message = ErrUnknownSource

		case errors.Is(err, httpservice.ErrMissingHeaderData):
			statusCode = http.StatusBadRequest
			message = ErrMissingHeaderData

		case errors.Is(err, httpservice.ErrInvalidToken):
			statusCode = http.StatusUnauthorized
			message = ErrInvalidToken

		case errors.Is(err, httpservice.ErrUserAlreadyCheckIn):
			statusCode = http.StatusBadRequest
			message = ErrUserAlreadyCheckIn

		case errors.Is(err, httpservice.ErrAttemptIsReachedMax):
			statusCode = http.StatusBadRequest
			message = ErrUserModuleAttemptReachedMax

		case errors.Is(err, httpservice.ErrAttemptIsAlreadyRecorded):
			statusCode = http.StatusBadRequest
			message = ErrUserModuleAttemptIsTaken

		case errors.Is(err, httpservice.ErrAssesmentCantUpdate):
			statusCode = http.StatusBadRequest
			message = ErrEditAssesmentWhichAlreadyTaken

		case errors.Is(err, httpservice.ErrUnauthorizedTokenData):
			statusCode = http.StatusUnauthorized
			message = ErrUnauthorizedTokenData

		case errors.Is(err, httpservice.ErrInvalidOTP):
			statusCode = http.StatusUnauthorized
			message = ErrInvalidOTP

		case errors.Is(err, httpservice.ErrInvalidOTPToken):
			statusCode = http.StatusUnauthorized
			message = ErrInvalidOTPToken

		case errors.Is(err, httpservice.ErrInvalidPhoneNumberOTP):
			statusCode = http.StatusUnauthorized
			message = ErrInvalidPhoneNumberOTP

		case errors.Is(err, httpservice.ErrPasswordNotMatch):
			statusCode = http.StatusUnauthorized
			message = ErrPasswordNotMatch

		case errors.Is(err, httpservice.ErrConfirmPasswordNotMatch):
			statusCode = http.StatusBadRequest
			message = ErrPasswordNotMatch

		case errors.Is(err, httpservice.ErrNoResultData):
			statusCode = http.StatusNotFound
			message = ErrNoResultData

		case errors.Is(err, httpservice.ErrUserAlreadyRegistered):
			statusCode = http.StatusConflict
			message = ErrUserAlreadyRegistered

		case errors.Is(err, httpservice.ErrUserNotFound):
			statusCode = http.StatusUnauthorized
			message = ErrUserNotFound

		case errors.Is(err, httpservice.ErrUnauthorizedUser):
			statusCode = http.StatusUnauthorized
			message = ErrUnauthorizedUser

		case errors.Is(err, httpservice.ErrInActiveUser):
			statusCode = http.StatusUnauthorized
			message = ErrInactiveUser

		case errors.Is(err, httpservice.SuccessChangePassword):
			statusCode = http.StatusOK
			message = SuccessChangedPassword

		case errors.Is(err, httpservice.ErrRoleNotFound):
			statusCode = http.StatusConflict
			message = ErrRoleNotFound

		case errors.Is(err, httpservice.ErrInvalidromotionCode):
			statusCode = http.StatusForbidden
			message = ErrInvalidPromotionCode

		case errors.Is(err, httpservice.ErrInsufficientQuantityVoucher):
			statusCode = http.StatusConflict
			message = ErrInsufficientQuantityVoucher

		case errors.Is(err, httpservice.ErrVoucherIsNotActive):
			statusCode = http.StatusConflict
			message = ErrVoucherIsNotActive

		case errors.Is(err, httpservice.ErrVoucherIsExpired):
			statusCode = http.StatusConflict
			message = ErrVoucherIsExpired

		case errors.Is(err, httpservice.ErrInvalidPaymentID):
			statusCode = http.StatusConflict
			message = ErrInvalidPaymentID

		case errors.Is(err, httpservice.ErrBadRequestWithMessage):
			statusCode = http.StatusBadRequest

			langError := strings.Split(err.Error(), "|")
			message = constants.ErrorResponse{
				ID: strings.ReplaceAll(langError[1], ": error with message", ""),
				EN: langError[0],
			}
		case errors.Is(err, httpservice.ErrNoResultDataWithMessage):
			statusCode = http.StatusNotFound

			langError := strings.Split(err.Error(), "|")
			message = constants.ErrorResponse{
				ID: strings.ReplaceAll(langError[1], ": no result data with message", ""),
				EN: langError[0],
			}
		case errors.Is(err, httpservice.ErrNIKAlreadyExist):
			statusCode = http.StatusConflict
			message = ErrNIKAlreadyExist
		case errors.Is(err, httpservice.ErrIDCardAlreadyExist):
			statusCode = http.StatusConflict
			message = ErrIDCardAlreadyExist
		case errors.Is(err, httpservice.ErrNPWPAlreadyExist):
			statusCode = http.StatusConflict
			message = ErrNPWPAlreadyExist
		case errors.Is(err, httpservice.ErrEmailAlreadyExist):
			statusCode = http.StatusConflict
			message = ErrEmailAlreadyExist
		case errors.Is(err, httpservice.ErrPhoneNumberAlreadyExist):
			statusCode = http.StatusConflict
			message = ErrPhoneNumberAlreadyExist
		case errors.Is(err, httpservice.ErrUnknownSourceWithMessage):
			statusCode = http.StatusUnprocessableEntity
		case errors.Is(err, httpservice.ErrEmailNotFound):
			statusCode = http.StatusBadRequest
			message = ErrEmailNotFound

			//reservation
		case errors.Is(err, httpservice.ErrArrivalDateNil):
			statusCode = http.StatusBadRequest
			message = ErrArrivalDateNil

		case errors.Is(err, httpservice.ErrArrivalDateFormat):
			statusCode = http.StatusBadRequest
			message = ErrArrivalDateFormat

		case errors.Is(err, httpservice.ErrArrivalDatePast):
			statusCode = http.StatusBadRequest
			message = ErrArrivalDatePast

		case errors.Is(err, httpservice.ErrDepartureDateNil):
			statusCode = http.StatusBadRequest
			message = ErrDepartureDateNil

		case errors.Is(err, httpservice.ErrDepartureDateFormat):
			statusCode = http.StatusBadRequest
			message = ErrDepartureDateFormat

		case errors.Is(err, httpservice.ErrDepartureDateAfter):
			statusCode = http.StatusBadRequest
			message = ErrDepartureDateAfter

		case errors.Is(err, httpservice.ErrAddressNil):
			statusCode = http.StatusBadRequest
			message = ErrAddressNil

		case errors.Is(err, httpservice.ErrEmailNil):
			statusCode = http.StatusBadRequest
			message = ErrEmailNil

		case errors.Is(err, httpservice.ErrEmailFormat):
			statusCode = http.StatusBadRequest
			message = ErrEmailFormat

		case errors.Is(err, httpservice.ErrFirstnameNil):
			statusCode = http.StatusBadRequest
			message = ErrFirstnameNil

		case errors.Is(err, httpservice.ErrLastnameNil):
			statusCode = http.StatusBadRequest
			message = ErrLastnameNil

		case errors.Is(err, httpservice.ErrTelephoneNil):
			statusCode = http.StatusBadRequest
			message = ErrTelephoneNil

		case errors.Is(err, httpservice.ErrTelephoneFormat):
			statusCode = http.StatusBadRequest
			message = ErrTelephoneFormat

		case errors.Is(err, httpservice.ErrRoomOutOfStock):
			statusCode = http.StatusBadRequest
			message = ErrRoomOutOfStock

		case errors.Is(err, httpservice.ErrSalutationsNil):
			statusCode = http.StatusBadRequest
			message = ErrSalutationsNil

		case errors.Is(err, httpservice.ErrPropertyNotFound):
			statusCode = http.StatusBadRequest
			message = ErrPropertyNotFound

		case errors.Is(err, httpservice.ErrDateToNil):
			statusCode = http.StatusBadRequest
			message = ErrDateToNil

		case errors.Is(err, httpservice.ErrSameDate):
			statusCode = http.StatusBadRequest
			message = ErrSameDate

		case errors.Is(err, httpservice.ErrDateToFalse):
			statusCode = http.StatusBadRequest
			message = ErrDateToFalse

		case errors.Is(err, httpservice.ErrDateRangeTooLong):
			statusCode = http.StatusBadRequest
			message = ErrDateRangeTooLong

		case errors.Is(err, httpservice.ErrDateFromPast):
			statusCode = http.StatusBadRequest
			message = ErrDateFromPast

		case errors.Is(err, httpservice.ErrCountryNil):
			statusCode = http.StatusBadRequest
			message = ErrCountryNil

		case errors.Is(err, httpservice.ErrCountryNull):
			statusCode = http.StatusBadRequest
			message = ErrCountryNull

		case errors.Is(err, httpservice.ErrAdultMax):
			statusCode = http.StatusBadRequest
			message = ErrAdultMax

		case errors.Is(err, httpservice.ErrAdultMin):
			statusCode = http.StatusBadRequest
			message = ErrAdultMin

		case errors.Is(err, httpservice.ErrChildMax):
			statusCode = http.StatusBadRequest
			message = ErrChildMax

		case errors.Is(err, httpservice.ErrChildrenMax):
			statusCode = http.StatusBadRequest
			message = ErrChildrenMax

		case errors.Is(err, httpservice.ErrPropertyNil):
			statusCode = http.StatusBadRequest
			message = ErrPropertyNil

		case errors.Is(err, httpservice.ErrPropertyStaahNil):
			statusCode = http.StatusBadRequest
			message = ErrPropertyStaahNil

		case errors.Is(err, httpservice.ErrBookingDetailNil):
			statusCode = http.StatusBadRequest
			message = ErrBookingDetailNil

		case errors.Is(err, httpservice.ErrRoomNil):
			statusCode = http.StatusBadRequest
			message = ErrRoomNil

		case errors.Is(err, httpservice.ErrRatesNil):
			statusCode = http.StatusBadRequest
			message = ErrRatesNil

		case errors.Is(err, httpservice.ErrRoomNotFound):
			statusCode = http.StatusBadRequest
			message = ErrRoomNotFound

		case errors.Is(err, httpservice.ErrRateNotFound):
			statusCode = http.StatusBadRequest
			message = ErrRateNotFound

		case errors.Is(err, httpservice.ErrRatesNull):
			statusCode = http.StatusBadRequest
			message = ErrRatesNull

			langError := strings.Split(err.Error(), "|")
			message = constants.ErrorResponse{
				ID: strings.ReplaceAll(langError[1], ": unknown error with message", ""),
				EN: langError[0],
			}
		case errors.Is(err, httpservice.ErrPowerproServerDown):
			statusCode = http.StatusInternalServerError
			message = ErrInternalServerErrror
		}

		_ = ctx.JSON(statusCode, echo.NewHTTPError(statusCode, message))
	}
}
