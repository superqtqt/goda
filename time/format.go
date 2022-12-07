package time

import "strings"

var replace = [][]string{
	[]string{"yyyy", "2006"},
	[]string{"yy", "06"},
	[]string{"MMMM", "January"},
	[]string{"MMM", "Jan"},
	[]string{"MM", "01"},
	[]string{"M", "1"},
	[]string{"dd", "02"},
	[]string{"d", "02"},
	[]string{"HH", "15"},
	[]string{"hh", "03"},
	[]string{"mm", "04"},
	[]string{"m", "04"},
	[]string{"ss", "05"},
	[]string{"s", "05"},
	[]string{"cccc", "Monday"},
	[]string{"cccc", "Mon"},
}

// FormatParse
//  @Description:
//  @param ymdFormat
//  @return string
func FormatParse(ymdFormat string) string {
	rs := ymdFormat
	for i := range replace {
		r := replace[i]
		rs = strings.ReplaceAll(rs, r[0], r[1])
	}
	return rs
}
