package cmd

import (
	"bufio"
	"io"
	"math/rand"
	"strings"
	"sync"
	"time"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

var (
	random  *rand.Rand
	getRand sync.Once

	getLocation sync.Once

	now    time.Time
	getNow sync.Once
)

func GetRand() *rand.Rand {
	getRand.Do(func() {
		random = rand.New(rand.NewSource(time.Now().UnixNano()))
	})
	return random
}

func GetLocation() *time.Location {
	getLocation.Do(func() {
		loc, _ := time.LoadLocation(GetFlags().Location)
		time.Local = loc
	})
	return time.Local
}

func Now() time.Time {
	getNow.Do(func() {
		timeString := GetFlags().Now
		if len(timeString) < 1 {
			now = time.Now()
		} else {
			t, _ := time.ParseInLocation("2006-01-02 15:04:05.999999999", timeString, GetLocation())
			now = t
		}
	})
	return now
}

func GetReader(r io.Reader, enc Encoding) io.Reader {
	if enc == SJIS {
		return transform.NewReader(r, japanese.ShiftJIS.NewDecoder())
	}
	return bufio.NewReader(r)
}

func UnescapeString(s string) string {
	s = strings.Replace(s, "\\a", "\a", -1)
	s = strings.Replace(s, "\\b", "\b", -1)
	s = strings.Replace(s, "\\f", "\f", -1)
	s = strings.Replace(s, "\\n", "\n", -1)
	s = strings.Replace(s, "\\r", "\r", -1)
	s = strings.Replace(s, "\\t", "\t", -1)
	s = strings.Replace(s, "\\v", "\v", -1)
	s = strings.Replace(s, "\\\"", "\"", -1)
	s = strings.Replace(s, "\\\\", "\\", -1)
	return s
}
