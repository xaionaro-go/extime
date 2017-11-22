package extime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Time time.Time

func ParseTime(layout, value string) (Time, error) {
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
func (t *Time) Scan(src interface{}) (err error) {
	switch srcTyped := src.(type) {
	case time.Time:
		*t = Time(srcTyped)
	case []uint8:
		*t, err = ParseTime("2006-01-02 15:04:05", string(srcTyped))
	default:
		err = fmt.Errorf("don't know how to covert %T (\"%v\") to extime.Time", src, src)
	}
	return
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
func (t Time) Date() Date {
	return Date(t)
}
func (t Time) AddDate(years int, months int, days int) Time {
	return Time(time.Time(t).AddDate(years, months, days))
}

type Date time.Time

func ParseDate(layout, value string) (Date, error) {
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
func (t *Date) Scan(src interface{}) (err error) {
	switch srcTyped := src.(type) {
	case time.Time:
		*t = Date(srcTyped)
	case []uint8:
		var tTmp Time
		tTmp, err = ParseTime("2006-01-02", string(srcTyped))
		if err != nil {
			tTmp, err = ParseTime("2006-01-02 15:04:05", string(srcTyped))
		}
		*t = Date(tTmp)
	default:
		err = fmt.Errorf("don't know how to covert %T (\"%v\") to extime.Time", src, src)
	}
	return
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
func (t Date) Time() Time {
	return Time(t)
}
func (t Date) AddDate(years int, months int, days int) Date {
	return Date(time.Time(t).AddDate(years, months, days))
}

func Now() Time {
	return Time(time.Now())
}
func NowDate() Date {
	return Date(time.Now())
}
