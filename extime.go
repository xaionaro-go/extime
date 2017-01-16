package extime

import (
	"database/sql/driver"
	"time"
)

type Time time.Time

func ParseTime (layout, value string) (Time, error) {
	time, err := time.Parse(layout, value)
	return Time(time), err
}

func (t Time) Format(fmt string) string {
	return (time.Time)(t).Format(fmt)
}
func (t Time) Unix() int64 {
	return (time.Time)(t).Unix()
}
func (t Time) UnixNano() int64 {
	return (time.Time)(t).UnixNano()
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
func (t Time) IsInFuture() bool {
	return t.UnixNano() > time.Now().UnixNano()
}

type Date time.Time

func ParseDate (layout, value string) (Date, error) {
	time, err := time.Parse(layout, value)
	return Date(time), err
}

func (t Date) Format(fmt string) string {
	return (time.Time)(t).Format(fmt)
}
func (t Date) Unix() int64 {
	return (time.Time)(t).Unix()
}
func (t Date) UnixNano() int64 {
	return (time.Time)(t).UnixNano()
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
func (t Date) IsInFuture() bool {
	return t.UnixNano() > time.Now().UnixNano()
}

