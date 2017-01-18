// Package goymd convert golang style time format and other general time format.
//
// Example
// -------
// fmt.Println(goymd.YmdStyle("2006-01-02T15:04:05Z07:00")) // => %Y-%m-%dT%H:%M:%S%:z
// fmt.Println(goymd.GoStyle("%Y/%m/%d %H:%M:%S %Z")) // => 2006/01/02 15:04:05 MST
//
// Supported golang style formats are suggested by following examples.
// ( see https://golang.org/pkg/time/#pkg-constants )
// -----------------------------------------
// ANSIC       = "Mon Jan _2 15:04:05 2006"
// UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
// RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
// RFC822      = "02 Jan 06 15:04 MST"
// RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
// RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
// RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
// RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
// RFC3339     = "2006-01-02T15:04:05Z07:00"
// RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
// Kitchen     = "3:04PM"
// Stamp       = "Jan _2 15:04:05"
// StampMilli  = "Jan _2 15:04:05.000"
// StampMicro  = "Jan _2 15:04:05.000000"
// StampNano   = "Jan _2 15:04:05.000000000"
package goymd

import "strings"

// YmdStyle convert golang style to general style
func YmdStyle(s string) string {
	for _, m := range mapping {
		s = strings.Replace(s, m.g, m.y, -1)
	}
	return s
}

// GoStyle convert general style to golang style
func GoStyle(s string) string {
	for _, m := range mapping {
		s = strings.Replace(s, m.y, m.g, -1)
	}
	return s
}

type m struct {
	g string
	y string
}

var mapping = []m{
	{g: "2006", y: "%Y"}, // year
	{g: "06", y: "%y"},
	{g: "01", y: "%m"}, // month
	{g: "Jan", y: "%b"},
	{g: "02", y: "%d"}, // date
	{g: "_2", y: "%e"},
	{g: "15", y: "%H"}, // hours
	{g: "3", y: "%I"},
	{g: "04", y: "%M"},         // minutes
	{g: "05", y: "%S"},         // seconds
	{g: "999999999", y: "%9N"}, // nano seconds
	{g: "000000000", y: "%N"},
	{g: "000000", y: "%6N"},
	{g: "000", y: "%3N"},
	{g: "Monday", y: "%A"}, // days of week
	{g: "Mon", y: "%a"},
	{g: "PM", y: "%p"},      // AM/PM
	{g: "MST", y: "%Z"},     // timezone
	{g: "Z07:00", y: "%:z"}, // offset
	{g: "-0700", y: "%z"},
}
