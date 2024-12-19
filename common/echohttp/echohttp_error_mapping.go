package echohttp

import "gitlab.com/wit-id/go-orm-mysql/common/constants"

var (
	ErrInternalServerErrror = constants.ErrorResponse{
		ID: "Terjadi Kesalahan Pada Server",
		EN: "Internal Server Error",
	}

	ErrBadRequest = constants.ErrorResponse{
		ID: "Payload Permintaan Tidak Sesuai",
		EN: "Bad Request Payload",
	}

	ErrInvalidAppKey = constants.ErrorResponse{
		ID: "APP Key Tidak Sesuai",
		EN: "Invalid APP Key",
	}

	ErrUnknownSource = constants.ErrorResponse{
		ID: "Error Tidak Diketahui",
		EN: "Unknown Error",
	}

	ErrMissingHeaderData = constants.ErrorResponse{
		ID: "Data Header Tidak Ada",
		EN: "Missing Header Data",
	}

	ErrInvalidToken = constants.ErrorResponse{
		ID: "Token Tidak Valid",
		EN: "Invalid Token",
	}

	ErrUserAlreadyCheckIn = constants.ErrorResponse{
		ID: "user sudah check in hari ini",
		EN: "this user is checked in already",
	}

	ErrUserModuleAttemptReachedMax = constants.ErrorResponse{
		ID: "user untuk module attempt ini sudah melampaui batas",
		EN: "user for this module attempt reached max",
	}

	ErrUserModuleAttemptIsTaken = constants.ErrorResponse{
		ID: "user untuk module attempt ini sudah tersimpan di database",
		EN: "user for this module attempt is taken already",
	}

	ErrEditAssesmentWhichAlreadyTaken = constants.ErrorResponse{
		ID: "user tidak dapat mengubah assesment ketika assesment sudah di gunakan",
		EN: "user cant update this assesment while this assesment already taken",
	}

	ErrUnauthorizedTokenData = constants.ErrorResponse{
		ID: "Data Token Tidak Sah",
		EN: "Unauthorized token data",
	}

	ErrInvalidOTP = constants.ErrorResponse{
		ID: "OTP Tidak Valid",
		EN: "Invalid OTP",
	}

	ErrInvalidOTPToken = constants.ErrorResponse{
		ID: "OTP Token Tidak Valid",
		EN: "Invalid OTP Token",
	}

	ErrInvalidPhoneNumberOTP = constants.ErrorResponse{
		ID: "Nomor Telefon Anda Tidak Valid Untuk OTP ini",
		EN: "Your Phone Number Is Invalid For This OTP",
	}

	ErrPasswordNotMatch = constants.ErrorResponse{
		ID: "Kata Sandi Tidak Cocok",
		EN: "Password Not Match",
	}

	ErrConfirmPasswordNotMatch = constants.ErrorResponse{
		ID: "Konfirmasi Kata Sandi Tidak Cocok",
		EN: "Confirm Password Not Match",
	}

	SuccessChangedPassword = constants.ErrorResponse{
		ID: "Kata Sandi Berhasil Diganti",
		EN: "Successfully Changed Password",
	}

	ErrNoResultData = constants.ErrorResponse{
		ID: "Tidak Ada Data Hasil",
		EN: "No Result Data",
	}

	ErrUserAlreadyRegistered = constants.ErrorResponse{
		ID: "Email atau Nomor Telefon Sudah Terdaftar",
		EN: "Email or Phone Number is Already Registered",
	}

	ErrUserNotFound = constants.ErrorResponse{
		ID: "User Tidak Ditemukan",
		EN: "User Not Found",
	}

	ErrUnauthorizedUser = constants.ErrorResponse{
		ID: "User Tidak Sah",
		EN: "Unauthorized User",
	}

	ErrInactiveUser = constants.ErrorResponse{
		ID: "User Tidak Aktif",
		EN: "User not Active",
	}

	ErrRoleNotFound = constants.ErrorResponse{
		ID: "Role Tidak Ditemukan",
		EN: "Role not Found",
	}

	ErrInvalidPromotionCode = constants.ErrorResponse{
		ID: "Kode Promosi Tidak Valid",
		EN: "Invalid Promotion Code",
	}

	ErrInsufficientQuantityVoucher = constants.ErrorResponse{
		ID: "Kuantitas Voucher Tidak Mencukupi",
		EN: "Insufficient Quantities of Voucher",
	}

	ErrVoucherIsNotActive = constants.ErrorResponse{
		ID: "Voucher tidak aktif",
		EN: "Voucher Is not Active",
	}

	ErrVoucherIsExpired = constants.ErrorResponse{
		ID: "Voucher Sudah Kadaluarsa",
		EN: "Voucher is Expired",
	}

	ErrInvalidPaymentID = constants.ErrorResponse{
		ID: "ID Pembayaran Tidak Valid",
		EN: "Invalid Payment ID",
	}

	ErrNIKAlreadyExist = constants.ErrorResponse{
		ID: "Nomor NIK Sudah Terdaftar",
		EN: "NIK Number is Already Registered",
	}

	ErrIDCardAlreadyExist = constants.ErrorResponse{
		ID: "Nomor ID Card Sudah Terdaftar",
		EN: "ID Card Number is Already Registered",
	}

	ErrNPWPAlreadyExist = constants.ErrorResponse{
		ID: "Nomor NPWP Sudah Terdaftar",
		EN: "NPWP Number is Already Registered",
	}

	ErrEmailAlreadyExist = constants.ErrorResponse{
		ID: "Alamat Email Sudah Terdaftar",
		EN: "Email Address is Already Registered",
	}

	ErrPhoneNumberAlreadyExist = constants.ErrorResponse{
		ID: "Nomor Telepon Sudah Terdaftar",
		EN: "Phone Number is Already Registered",
	}
	ErrArrivalDateNil = constants.ErrorResponse{
		ID: "Pilih tanggal check-in untuk melanjutkan pemesanan Anda.",
		EN: "Please select a check-in date to proceed with your booking.",
	}

	//reservation
	ErrEmailNotFound = constants.ErrorResponse{
		ID: "Email tidak ditemukan.",
		EN: "Email not found.",
	}

	ErrArrivalDateFormat = constants.ErrorResponse{
		ID: "Format tanggal check-in salah.",
		EN: "The check-in date format is incorrect.",
	}

	ErrArrivalDatePast = constants.ErrorResponse{
		ID: "Tanggal check-in tidak boleh di masa lalu.",
		EN: "Check-in date cannot be in the past.",
	}

	ErrDepartureDateNil = constants.ErrorResponse{
		ID: "Pilih tanggal check-out untuk melanjutkan.",
		EN: "Please select a check-out date to continue.",
	}

	ErrDepartureDateFormat = constants.ErrorResponse{
		ID: "Format tanggal check-out salah.",
		EN: "The check-out date format is incorrect.",
	}

	ErrDepartureDateAfter = constants.ErrorResponse{
		ID: "Tanggal check-out harus setelah tanggal check-in.",
		EN: "The check-out date must be after the check-in date.",
	}

	ErrAddressNil = constants.ErrorResponse{
		ID: "Alamat tidak boleh kosong. Silakan masukkan alamat Anda untuk melanjutkan.",
		EN: "Address cannot be blank. Please enter your address to proceed.",
	}

	ErrEmailNil = constants.ErrorResponse{
		ID: "Alamat email tidak boleh kosong. Masukkan alamat email yang valid.",
		EN: "Email address cannot be empty. Please enter a valid email address.",
	}

	ErrEmailFormat = constants.ErrorResponse{
		ID: "Format email salah. Masukkan alamat email yang valid.",
		EN: "The email format is incorrect. Please enter a valid email address.",
	}

	ErrFirstnameNil = constants.ErrorResponse{
		ID: "Nama depan tidak boleh kosong. Masukkan nama depan Anda.",
		EN: "First name cannot be blank. Please enter your first name.",
	}

	ErrLastnameNil = constants.ErrorResponse{
		ID: "Nama belakang tidak boleh kosong. Masukkan nama belakang Anda.",
		EN: "Last name cannot be blank. Please enter your last name.",
	}

	ErrTelephoneNil = constants.ErrorResponse{
		ID: "Nomor telepon tidak boleh kosong. Masukkan nomor telepon Anda.",
		EN: "Telephone number cannot be blank. Please enter your telephone number.",
	}

	ErrTelephoneFormat = constants.ErrorResponse{
		ID: "Format nomor telepon salah. Masukkan nomor telepon yang valid.",
		EN: "The telephone number format is incorrect. Please enter a valid phone number.",
	}

	ErrRoomOutOfStock = constants.ErrorResponse{
		ID: "Mohon maaf, kamar yang anda pilih sudah habis",
		EN: "We apologize, the room you selected is out of stock.",
	}

	ErrSalutationsNil = constants.ErrorResponse{
		ID: "Silakan pilih salam untuk melanjutkan.",
		EN: "Please select a salutation to proceed.",
	}

	//List Room
	ErrPropertyNotFound = constants.ErrorResponse{
		ID: "Kami tidak bisa menemukan properti yang Anda cari. Silakan periksa detail properti atau coba cari properti lain.",
		EN: "We couldn't find the property you're looking for. Please check the property details or try searching for a different one.",
	}

	ErrDateToNil = constants.ErrorResponse{
		ID: "Pilih tanggal check-out untuk melanjutkan",
		EN: "Please select a check-out date to continue",
	}

	ErrSameDate = constants.ErrorResponse{
		ID: "Tanggal check-in dan check-out Anda tidak boleh sama. Silakan pilih tanggal yang berbeda.",
		EN: "Your check-in and check-out dates cannot be the same. Please select different dates.",
	}

	ErrDateToFalse = constants.ErrorResponse{
		ID: "Tanggal check-in harus sebelum tanggal check-out. Harap sesuaikan tanggal dengan yang sesuai.",
		EN: "The check-in date must be before the check-out date. Please adjust the dates accordingly.",
	}

	ErrDateRangeTooLong = constants.ErrorResponse{
		ID: "Pemesanan dapat dilakukan selama maksimal 31 hari. Mohon untuk mempersingkat durasi menginap Anda.",
		EN: "Bookings can be made for a maximum of 31 days. Please shorten your stay duration.",
	}

	ErrDateFromPast = constants.ErrorResponse{
		ID: "Anda tidak dapat memilih tanggal yang sudah lewat.",
		EN: "You cannot select a past date.",
	}

	ErrCountryNil = constants.ErrorResponse{
		ID: "Pilih negara untuk melanjutkan pemesanan Anda.",
		EN: "Please select a country to proceed with your booking.",
	}

	ErrCountryNull = constants.ErrorResponse{
		ID: "Pilih negara untuk melanjutkan pemesanan Anda.",
		EN: "Please select a country to proceed with your booking.",
	}

	ErrAdultMax = constants.ErrorResponse{
		ID: "Orang dewasa tidak boleh lebih dari 2.",
		EN: "Adult cannot be more than 2.",
	}

	ErrAdultMin = constants.ErrorResponse{
		ID: "adults should not be less than the number of rooms",
		EN: "orang dewasa tidak boleh kurang dari jumlah ruangan.",
	}

	ErrChildMax = constants.ErrorResponse{
		ID: "Anak tidak boleh lebih dari 1.",
		EN: "Child cannot be more than 1.",
	}

	ErrChildrenMax = constants.ErrorResponse{
		ID: "Anak tidak boleh lebih dari 1.",
		EN: "Child cannot be more than 1.",
	}

	ErrPropertyNil = constants.ErrorResponse{
		ID: "Pilih properti untuk melanjutkan pemesanan Anda.",
		EN: "Please select a property to proceed with your booking.",
	}

	ErrPropertyStaahNil = constants.ErrorResponse{
		ID: "Pilih properti staah untuk melanjutkan pemesanan Anda.",
		EN: "Please select a property staah to proceed with your booking.",
	}

	ErrBookingDetailNil = constants.ErrorResponse{
		ID: "Pilih detail booking untuk melanjutkan pemesanan Anda.",
		EN: "Please select a booking detail to proceed with your booking.",
	}

	ErrRoomNil = constants.ErrorResponse{
		ID: "Pilih room untuk melanjutkan pemesanan Anda.",
		EN: "Please select a room to proceed with your booking.",
	}

	ErrRatesNil = constants.ErrorResponse{
		ID: "Pilih rate id untuk melanjutkan pemesanan Anda.",
		EN: "Please select a rate id to proceed with your booking.",
	}

	ErrRoomNotFound = constants.ErrorResponse{
		ID: "Kami tidak bisa menemukan room id yang Anda cari. Silakan periksa room id atau coba cari properti lain.",
		EN: "We couldn't find the room id you're looking for. Please check the room id or try searching for a different one.",
	}

	ErrRateNotFound = constants.ErrorResponse{
		ID: "Kami tidak bisa menemukan rate id yang Anda cari. Silakan periksa rate id atau coba cari properti lain.",
		EN: "We couldn't find the rate id you're looking for. Please check the rate id or try searching for a different one.",
	}

	ErrRatesNull = constants.ErrorResponse{
		ID: "Pilih rate id untuk melanjutkan pemesanan Anda.",
		EN: "Please select a rate id to proceed with your booking.",
	}

	ErrPowerproServerDown = constants.ErrorResponse{
		ID: "Saat ini server Powerpro Sedang tidak bisa di akses, mohon di coba lagi beberapa saat",
		EN: "Couldn't access to powerpro server at the moment, please try again later",
	}
)
