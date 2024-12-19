package payload

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	defaultLimit      = 10
	defaultOrderValue = "created_at DESC"
)

type CommonLanguagePayload struct {
	EN string `json:"en"`
	ID string `json:"id"`
}

func limitWithDefault(limit int32) int32 {
	if limit <= 0 {
		return defaultLimit
	}

	return limit
}

func makeOffset(limit, offset int32) int32 {

	if offset == 0 {
		return (1 * limit) - limit
	} else {
		return (offset * limit) - limit
	}
}

func makeOrderParam(orderBy, sort string) string {
	if orderBy == "" || sort == "" {
		return defaultOrderValue
	}

	return fmt.Sprintf(strings.ToLower("%s %s"), orderBy, sort)
}

func queryStringLike(param string) string {
	return "%" + param + "%"
}

func isNumberAndDot(str string) bool {
	re := regexp.MustCompile(`^[0-9.]+$`)
	return re.MatchString(str)
}

func containsAlphabet(str string) bool {
	re := regexp.MustCompile(`[a-zA-Z]`)
	return re.MatchString(str)
}

type CommonSubResponsePayload struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Pagination struct {
	Limit  int32  `json:"limit" valid:"required"`
	Offset int32  `json:"page" valid:"required"`
	Order  string `json:"order" valid:"required"`
	Sort   string `json:"sort" valid:"required"` // ASC, DESC
}

type ResponseData struct {
	Code      string      `json:"code"`
	Status    string      `json:"status"`
	Data      interface{} `json:"data"`
	MessageEn string      `json:"message_en"`
	MessageID string      `json:"message_id"`
}

type CommonPayload struct {
	GuID string `json:"guid" param:"guid"`
	Name string `json:"name" param:"name"`
}
