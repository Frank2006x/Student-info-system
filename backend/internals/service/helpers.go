package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func optionalDate(value string) (interface{}, error) {
	if strings.TrimSpace(value) == "" {
		return nil, nil
	}
	return requiredDate(value)
}

func requiredDate(value string) (pgtype.Date, error) {
	parsed, err := time.Parse("2006-01-02", value)
	if err != nil {
		return pgtype.Date{}, fmt.Errorf("date must use YYYY-MM-DD")
	}
	return pgtype.Date{Time: parsed, Valid: true}, nil
}

func optionalString(value string) interface{} {
	if strings.TrimSpace(value) == "" {
		return nil
	}
	return value
}

func nullableID(value *int64) pgtype.Int8 {
	if value == nil {
		return pgtype.Int8{}
	}
	return pgtype.Int8{Int64: *value, Valid: true}
}
