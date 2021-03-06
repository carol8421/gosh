// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The time package provides functionality for measuring and
// displaying time.
package time

import (
	"os";
)

// Seconds reports the number of seconds since the Unix epoch,
// January 1, 1970 00:00:00 UTC.
func Seconds() int64 {
	sec, _, err := os.Time();
	if err != nil {
		panic("time: os.Time: ", err.String());
	}
	return sec;
}

// Nanoseconds reports the number of nanoseconds since the Unix epoch,
// January 1, 1970 00:00:00 UTC.
func Nanoseconds() int64 {
	sec, nsec, err := os.Time();
	if err != nil {
		panic("time: os.Time: ", err.String());
	}
	return sec*1e9 + nsec;
}

// Days of the week.
const (
	Sunday	= iota;
	Monday;
	Tuesday;
	Wednesday;
	Thursday;
	Friday;
	Saturday;
)

// Time is the struct representing a parsed time value.
type Time struct {
	Year			int64;	// 2008 is 2008
	Month, Day		int;	// Sep-17 is 9, 17
	Hour, Minute, Second	int;	// 10:43:12 is 10, 43, 12
	Weekday			int;	// Sunday, Monday, ...
	ZoneOffset		int;	// seconds east of UTC
	Zone			string;
}

var nonleapyear = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var leapyear = []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func months(year int64) []int {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return leapyear;
	}
	return nonleapyear;
}

const (
	secondsPerDay	= 24*60*60;
	daysPer400Years	= 365*400 + 97;
	daysPer100Years	= 365*100 + 24;
	daysPer4Years	= 365*4 + 1;
	days1970To2001	= 31*365 + 8;
)

// SecondsToUTC converts sec, in number of seconds since the Unix epoch,
// into a parsed Time value in the UTC time zone.
func SecondsToUTC(sec int64) *Time {
	t := new(Time);

	// Split into time and day.
	day := sec / secondsPerDay;
	sec -= day * secondsPerDay;
	if sec < 0 {
		day--;
		sec += secondsPerDay;
	}

	// Time
	t.Hour = int(sec/3600);
	t.Minute = int((sec/60)%60);
	t.Second = int(sec%60);

	// Day 0 = January 1, 1970 was a Thursday
	t.Weekday = int((day+Thursday)%7);
	if t.Weekday < 0 {
		t.Weekday += 7;
	}

	// Change day from 0 = 1970 to 0 = 2001,
	// to make leap year calculations easier
	// (2001 begins 4-, 100-, and 400-year cycles ending in a leap year.)
	day -= days1970To2001;

	year := int64(2001);
	if day < 0 {
		// Go back enough 400 year cycles to make day positive.
		n := -day / daysPer400Years + 1;
		year -= 400*n;
		day += daysPer400Years * n;
	} else {
		// Cut off 400 year cycles.
		n := day / daysPer400Years;
		year += 400*n;
		day -= daysPer400Years * n;
	}

	// Cut off 100-year cycles
	n := day / daysPer100Years;
	year += 100*n;
	day -= daysPer100Years * n;

	// Cut off 4-year cycles
	n = day / daysPer4Years;
	year += 4*n;
	day -= daysPer4Years * n;

	// Cut off non-leap years.
	n = day/365;
	year += n;
	day -= 365*n;

	t.Year = year;

	// If someone ever needs yearday,
	// tyearday = day (+1?)

	months := months(year);
	var m int;
	yday := int(day);
	for m = 0; m < 12 && yday >= months[m]; m++ {
		yday -= months[m];
	}
	t.Month = m+1;
	t.Day = yday+1;
	t.Zone = "UTC";

	return t;
}

// UTC returns the current time as a parsed Time value in the UTC time zone.
func UTC() *Time {
	return SecondsToUTC(Seconds());
}

// SecondsToLocalTime converts sec, in number of seconds since the Unix epoch,
// into a parsed Time value in the local time zone.
func SecondsToLocalTime(sec int64) *Time {
	z, offset := lookupTimezone(sec);
	t := SecondsToUTC(sec+int64(offset));
	t.Zone = z;
	t.ZoneOffset = offset;
	return t;
}

// LocalTime returns the current time as a parsed Time value in the local time zone.
func LocalTime() *Time {
	return SecondsToLocalTime(Seconds());
}

// Seconds returns the number of seconds since January 1, 1970 represented by the
// parsed Time value.
func (t *Time) Seconds() int64 {
	// First, accumulate days since January 1, 2001.
	// Using 2001 instead of 1970 makes the leap-year
	// handling easier (see SecondsToUTC), because
	// it is at the beginning of the 4-, 100-, and 400-year cycles.
	day := int64(0);

	// Rewrite year to be >= 2001.
	year := t.Year;
	if year < 2001 {
		n := (2001-year)/400 + 1;
		year += 400*n;
		day -= daysPer400Years * n;
	}

	// Add in days from 400-year cycles.
	n := (year-2001)/400;
	year -= 400*n;
	day += daysPer400Years * n;

	// Add in 100-year cycles.
	n = (year-2001)/100;
	year -= 100*n;
	day += daysPer100Years * n;

	// Add in 4-year cycles.
	n = (year-2001)/4;
	year -= 4*n;
	day += daysPer4Years * n;

	// Add in non-leap years.
	n = year-2001;
	day += 365*n;

	// Add in days this year.
	months := months(t.Year);
	for m := 0; m < t.Month - 1; m++ {
		day += int64(months[m]);
	}
	day += int64(t.Day - 1);

	// Convert days to seconds since January 1, 2001.
	sec := day * secondsPerDay;

	// Add in time elapsed today.
	sec += int64(t.Hour)*3600;
	sec += int64(t.Minute)*60;
	sec += int64(t.Second);

	// Convert from seconds since 2001 to seconds since 1970.
	sec += days1970To2001 * secondsPerDay;

	// Account for local time zone.
	sec -= int64(t.ZoneOffset);
	return sec;
}

var longDayNames = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

var shortDayNames = []string{
	"Sun",
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
}

var shortMonthNames = []string{
	"---",
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

func copy(dst []byte, s string) {
	for i := 0; i < len(s); i++ {
		dst[i] = s[i];
	}
}

func decimal(dst []byte, n int) {
	if n < 0 {
		n = 0;
	}
	for i := len(dst)-1; i >= 0; i-- {
		dst[i] = byte(n%10 + '0');
		n /= 10;
	}
}

func addString(buf []byte, bp int, s string) int {
	n := len(s);
	copy(buf[bp : bp+n], s);
	return bp+n;
}

// Just enough of strftime to implement the date formats below.
// Not exported.
func format(t *Time, fmt string) string {
	buf := make([]byte, 128);
	bp := 0;

	for i := 0; i < len(fmt); i++ {
		if fmt[i] == '%' {
			i++;
			switch fmt[i] {
			case 'A':	// %A full weekday name
				bp = addString(buf, bp, longDayNames[t.Weekday]);
			case 'a':	// %a abbreviated weekday name
				bp = addString(buf, bp, shortDayNames[t.Weekday]);
			case 'b':	// %b abbreviated month name
				bp = addString(buf, bp, shortMonthNames[t.Month]);
			case 'd':	// %d day of month (01-31)
				decimal(buf[bp : bp+2], t.Day);
				bp += 2;
			case 'e':	// %e day of month ( 1-31)
				if t.Day >= 10 {
					decimal(buf[bp : bp+2], t.Day);
				} else {
					buf[bp] = ' ';
					buf[bp+1] = byte(t.Day + '0');
				}
				bp += 2;
			case 'H':	// %H hour 00-23
				decimal(buf[bp : bp+2], t.Hour);
				bp += 2;
			case 'M':	// %M minute 00-59
				decimal(buf[bp : bp+2], t.Minute);
				bp += 2;
			case 'S':	// %S second 00-59
				decimal(buf[bp : bp+2], t.Second);
				bp += 2;
			case 'Y':	// %Y year 2008
				decimal(buf[bp : bp+4], int(t.Year));
				bp += 4;
			case 'y':	// %y year 08
				decimal(buf[bp : bp+2], int(t.Year % 100));
				bp += 2;
			case 'Z':
				bp = addString(buf, bp, t.Zone);
			default:
				buf[bp] = '%';
				buf[bp+1] = fmt[i];
				bp += 2;
			}
		} else {
			buf[bp] = fmt[i];
			bp++;
		}
	}
	return string(buf[0:bp]);
}

// Asctime formats the parsed time value in the style of
// ANSI C asctime: Sun Nov  6 08:49:37 1994
func (t *Time) Asctime() string {
	return format(t, "%a %b %e %H:%M:%S %Y");
}

// RFC850 formats the parsed time value in the style of
// RFC 850: Sunday, 06-Nov-94 08:49:37 UTC
func (t *Time) RFC850() string {
	return format(t, "%A, %d-%b-%y %H:%M:%S %Z");
}

// RFC1123 formats the parsed time value in the style of
// RFC 1123: Sun, 06 Nov 1994 08:49:37 UTC
func (t *Time) RFC1123() string {
	return format(t, "%a, %d %b %Y %H:%M:%S %Z");
}

// String formats the parsed time value in the style of
// date(1) - Sun Nov  6 08:49:37 UTC 1994
func (t *Time) String() string {
	return format(t, "%a %b %e %H:%M:%S %Z %Y");
}
