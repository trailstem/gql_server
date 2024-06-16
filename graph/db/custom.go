package db

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// CustomTime is a wrapper around time.Time to implement the driver.Valuer interface
type CustomTime struct {
	time.Time
}

// Value implements the driver.Valuer interface
func (ct CustomTime) Value() (driver.Value, error) {
	if ct.IsZero() {
		return nil, nil
	}
	return ct.Time, nil
}

// Scan implements the sql.Scanner interface
func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		*ct = CustomTime{Time: time.Time{}}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*ct = CustomTime{Time: v}
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into CustomTime", value)
	}
}
