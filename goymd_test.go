package goymd

import "testing"

type testdata map[string]string

var data = []testdata{
	{"g": "Mon Jan _2 15:04:05 2006", "y": "%a %b %e %H:%M:%S %Y"},
	{"g": "Mon Jan _2 15:04:05 MST 2006", "y": "%a %b %e %H:%M:%S %Z %Y"},
	{"g": "Mon Jan 02 15:04:05 -0700 2006", "y": "%a %b %d %H:%M:%S %z %Y"},
	{"g": "02 Jan 06 15:04 MST", "y": "%d %b %y %H:%M %Z"},
	{"g": "02 Jan 06 15:04 -0700", "y": "%d %b %y %H:%M %z"},
	{"g": "Monday, 02-Jan-06 15:04:05 MST", "y": "%A, %d-%b-%y %H:%M:%S %Z"},
	{"g": "Mon, 02 Jan 2006 15:04:05 MST", "y": "%a, %d %b %Y %H:%M:%S %Z"},
	{"g": "Mon, 02 Jan 2006 15:04:05 -0700", "y": "%a, %d %b %Y %H:%M:%S %z"},
	{"g": "2006-01-02T15:04:05Z07:00", "y": "%Y-%m-%dT%H:%M:%S%:z"},
	{"g": "2006-01-02T15:04:05.999999999Z07:00", "y": "%Y-%m-%dT%H:%M:%S.%9N%:z"},
	{"g": "3:04PM", "y": "%I:%M%p"},
	{"g": "Jan _2 15:04:05", "y": "%b %e %H:%M:%S"},
	{"g": "Jan _2 15:04:05.000", "y": "%b %e %H:%M:%S.%3N"},
	{"g": "Jan _2 15:04:05.000000", "y": "%b %e %H:%M:%S.%6N"},
	{"g": "Jan _2 15:04:05.000000000", "y": "%b %e %H:%M:%S.%N"},
}

func TestYmdStyle(t *testing.T) {
	for i, d := range data {
		actual := YmdStyle(d["g"])
		if actual != d["y"] {
			t.Errorf("[%d] expected %v but actual %v", i, d["y"], actual)
		}
	}
}

func TestGoStyle(t *testing.T) {
	for i, d := range data {
		actual := GoStyle(d["y"])
		if actual != d["g"] {
			t.Errorf("[%d] expected %v but actual %v", i, d["g"], actual)
		}
	}
}
