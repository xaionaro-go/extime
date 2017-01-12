package time

import (
	"database/sql/driver"
	"time"
)

const (
	Nanosecond  time.Duration = time.Nanosecond
	Microsecond               = time.Microsecond
	Millisecond               = time.Millisecond
	Second                    = time.Second
	Minute                    = time.Minute
	Hour                      = time.Hour
)

type Time time.Time

func (t Time) Format(fmt string) string {
	return (time.Time)(t).Format(fmt)
}
func (t Time) Unix() int64 {
	return (time.Time)(t).Unix()
}
func (t *Time) Scan(src interface{}) error {
	*t = Time(src.(time.Time))
	return nil
}
func (t Time) String() string {
	return t.Format("2006-01-02 15:04:05")
}
func (t Time) Value() (driver.Value, error) {
	return []byte(t.String()), nil
}
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte("\"" + t.String() + "\""), nil
}

type Date time.Time

func (t Date) Format(fmt string) string {
	return (time.Time)(t).Format(fmt)
}
func (t Date) Unix() int64 {
	return (time.Time)(t).Unix()
}
func (t *Date) Scan(src interface{}) error {
	*t = Date(src.(time.Time))
	return nil
}
func (t Date) String() string {
	return t.Format("2006-01-02")
}
func (t Date) Value() (driver.Value, error) {
	return []byte(t.String()), nil
}
func (t Date) MarshalJSON() ([]byte, error) {
	return []byte("\"" + t.String() + "\""), nil
}
