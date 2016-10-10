package models

import (
	"math"
	"strconv"
	"strings"
	"time"
)

type DateInfo struct {
	DateId    time.Time
	MonthId   int
	MonthDesc string
	QtrId     int
	QtrDesc   string
	Year      int
}

type SortDirection struct {
	Field string
	Dir   string
}

func GetDateInfo(t time.Time) DateInfo {
	di := DateInfo{}

	year := t.Year()
	month := int(t.Month())

	monthid := strconv.Itoa(year) + LeftPad2Len(strconv.Itoa(month), "0", 2)
	monthdesc := t.Month().String() + " " + strconv.Itoa(year)

	qtr := 0
	if month%3 > 0 {
		qtr = int(math.Ceil(float64(month / 3)))
		qtr = qtr + 1
	} else {
		qtr = month / 3
	}

	qtrid := strconv.Itoa(year) + LeftPad2Len(strconv.Itoa(qtr), "0", 2)
	qtrdesc := "Q" + strconv.Itoa(qtr) + " " + strconv.Itoa(year)

	di.DateId = t.UTC()
	di.Year = year
	di.MonthDesc = monthdesc
	di.MonthId, _ = strconv.Atoi(monthid)
	di.QtrDesc = qtrdesc
	di.QtrId, _ = strconv.Atoi(qtrid)

	return di
}

func LeftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}
