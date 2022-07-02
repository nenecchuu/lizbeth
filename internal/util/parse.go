package util

import (
	"strconv"
	"time"

	errlib "github.com/nenecchuu/arcana/err"
)

func ParseStringToInt64(value string, fieldName string) (int64, errlib.Error) {
	parsedVal, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, NewInvalidFieldFormatErr(fieldName)
	}

	return parsedVal, nil
}

func ParseStringToTime(value string, timeFormat string, fieldName string) (time.Time, errlib.Error) {
	parsedVal, err := time.Parse(timeFormat, value)
	if err != nil {
		return time.Time{}, NewInvalidFieldFormatErr(fieldName)
	}

	return parsedVal, nil
}
