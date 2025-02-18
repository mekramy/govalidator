package funcs

import (
	"math/big"
	"mime/multipart"
	"net"
	"regexp"
	"strconv"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/inhies/go-bytesize"
)

// IsValidUsername checks if the username is valid (only letters, numbers, and underscores).
func IsValidUsername(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return re.MatchString(username)
}

// IsValidIranianPhone checks if the iranian phone number is valid.
func IsValidIranianPhone(phone string) bool {
	re := regexp.MustCompile(`^0[1-9][0-9]{9}$`)
	return re.MatchString(phone)
}

// IsValidIranianMobile checks if the iranian mobile number is valid.
func IsValidIranianMobile(mobile string) bool {
	re := regexp.MustCompile(`^09[0-9]{9}$`)
	return re.MatchString(mobile)
}

// IsValidIranianPostalCode checks if the iranian postal code is valid.
func IsValidIranianPostalCode(postalCode string) bool {
	re := regexp.MustCompile(`^[0-9]{10}$`)
	return re.MatchString(postalCode)
}

// IsValidIranianIDNumber checks if the Iranian ID (birth certificate) number is valid.
func IsValidIranianIDNumber(id string) bool {
	re := regexp.MustCompile(`^[0-9]{1,10}$`)
	return re.MatchString(id)
}

// IsValidIranianNationalCode checks if the Iranian National ID number is valid using the official checksum algorithm.
func IsValidIranianNationalCode(nationalCode string) bool {
	// National ID must be exactly 10 digits
	if len(nationalCode) != 10 {
		return false
	}

	// Check if it contains only digits
	re := regexp.MustCompile(`^[0-9]{10}$`)
	if !re.MatchString(nationalCode) {
		return false
	}

	// Checksum algorithm for Iranian National Code
	sum := 0
	for i := 0; i < 9; i++ {
		digit, _ := strconv.Atoi(string(nationalCode[i]))
		sum += digit * (10 - i)
	}

	remainder := sum % 11
	checkDigit, _ := strconv.Atoi(string(nationalCode[9]))

	if remainder < 2 {
		return checkDigit == remainder
	} else {
		return checkDigit == (11 - remainder)
	}
}

// IsValidIranianBankCard checks if the bank card number is valid using the Luhn algorithm.
func IsValidIranianBankCard(cardNumber string) bool {
	// Check if the card number is exactly 16 digits
	re := regexp.MustCompile(`^[0-9]{16}$`)
	if !re.MatchString(cardNumber) {
		return false
	}

	// Luhn algorithm for card validation
	sum := 0
	alternate := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(string(cardNumber[i]))
		if alternate {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		alternate = !alternate
	}

	return sum%10 == 0
}

// IsValidIranianIBAN checks if the Iranian IBAN (International Bank Account Number) is valid with or without the "IR" prefix.
func IsValidIranianIBAN(iban string) bool {
	// If it doesn't have "IR" at the beginning, add it
	if !strings.HasPrefix(iban, "IR") {
		iban = "IR" + iban
	}

	// Check if the IBAN is in the correct format (IR followed by 24 digits)
	re := regexp.MustCompile(`^IR[0-9]{24}$`)
	if !re.MatchString(iban) {
		return false
	}

	// Convert IBAN to a numeric format for MOD 97 check
	ibanNumeric := iban[2:] + "1827" // IR -> 1827

	// Convert the numeric string to a big integer
	bigInt, success := new(big.Int).SetString(ibanNumeric, 10)
	if !success {
		return false
	}

	// Perform MOD 97 check
	return new(big.Int).Mod(bigInt, big.NewInt(97)).Cmp(big.NewInt(1)) == 0
}

// IsValidIP checks if the given string is a valid IP address (IPv4 or IPv6).
func IsValidIP(ip string) bool {
	// Try to parse the IP address using net package
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil
}

// IsValidIPPort checks if the given IP:Port string is valid.
func IsValidIPPort(ipPort string) bool {
	// Split the IP:Port string by ":"
	parts := strings.Split(ipPort, ":")
	if len(parts) != 2 {
		// Should be exactly two parts: IP and Port
		return false
	}

	// Validate the IP part
	ip := parts[0]
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		// If the IP is invalid, return false
		return false
	}

	// Validate the Port part
	port := parts[1]
	portNum, err := strconv.Atoi(port)
	if err != nil || portNum < 1 || portNum > 65535 {
		// Port should be a number between 1 and 65535
		return false
	}

	// If both IP and Port are valid, return true
	return true
}

// IsValidFileSize checks if the file size is within the given min and max size.
func IsValidFileSize(file *multipart.FileHeader, min string, max string) (bool, error) {
	// Get the file size
	fileSize := file.Size

	// Convert min and max to bytes using go-bytesize
	minSize, err := bytesize.Parse(min)
	if err != nil {
		return false, err
	}
	maxSize, err := bytesize.Parse(max)
	if err != nil {
		return false, err
	}

	// Check if the file size is within the min and max size range
	return fileSize >= int64(minSize) && fileSize <= int64(maxSize), nil
}

// IsValidFileType checks if the MIME type of the file matches any of the provided valid MIME types
func IsValidFileType(file *multipart.FileHeader, mimes ...string) (bool, error) {
	// Open the file to read its content and determine MIME type
	f, err := file.Open()
	if err != nil {
		return false, err
	}
	defer f.Close()

	// Get MIME type using mimetype package
	mime, err := mimetype.DetectReader(f)
	if err != nil {
		return false, err
	}

	// Compare detected MIME type with allowed MIME types
	for _, mimeType := range mimes {
		if strings.ToLower(mime.String()) == strings.ToLower(mimeType) {
			return true, nil
		}
	}

	// If MIME type doesn't match any of the allowed types, return false
	return false, nil
}
