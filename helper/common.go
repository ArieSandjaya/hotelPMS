package helper

import (
	"sort"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func GetUUID() string {
	uuidWithHyphen := uuid.New()
	return uuidWithHyphen.String()
}

func GetRoman(number int) string {
	// create a denary_number:roman_symbol map
	romanMap := map[int]string{
		1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X", 40: "XL", 50: "L",
		90: "XC", 100: "C", 400: "CD", 500: "D", 900: "CM", 1000: "M",
	}
	// create a slice of slices
	rows := len(romanMap)
	matrix := make([][]string, rows)
	// create a slice of the map keys
	var key_slice []int
	// range of a map returns key, value pair
	// value is not needed so use blank identifier _
	for k := range romanMap {
		key_slice = append(key_slice, k)
	}
	// sort the slice in place, highest number first (decending)
	sort.Sort(sort.Reverse(sort.IntSlice(key_slice)))
	// create a slice of romanMap content slices
	row := 0
	// range of a slice returns index, value pair
	// index not needed so use blank identifier _
	for _, key := range key_slice {
		// convert int key to string key
		skey := strconv.Itoa(key)
		matrix[row] = []string{skey, romanMap[key]}
		row++

	}
	result := ""
	for _, item := range matrix {
		// convert string to int
		den, err := strconv.Atoi(item[0])
		if err != nil {
			panic(err)
		}
		sym := item[1]
		for number >= den {
			result += sym
			number -= den
		}
	}
	return result
}

func SetStatus(s string) string {
	switch s {
	case "DRAFT_ST":
		return "DRAFT"
	case "PROCESS_ST":
		return "PROCESSED"
	case "FINISH_ST":
		return "FINISHED"
	case "DELETE_ST":
		return "DELETED"
	}
	return "UNKNOWN"
}
