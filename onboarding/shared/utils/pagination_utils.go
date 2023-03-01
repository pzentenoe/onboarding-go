package utils

import (
	"strconv"
)

func ConvertPaginationValues(pageNumberStr, sizeStr string) (pageNumber int, size int, err error) {
	pageNumber, err = strconv.Atoi(pageNumberStr)
	if err != nil {
		return
	}
	size, err = strconv.Atoi(sizeStr)
	if err != nil {
		return
	}
	return
}
