package util

import (
	"strconv"
	"time"
)

const DefaultTimeFormat = "2006-01-02 15:04:05"

// GetHourTime 获取日期范围
func GetHourTime(timeType, start, end string) (startTime, endTime int64) {
	switch timeType {
	case "1": // 今天
		dateNow := time.Now()
		startTime = time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 0, 0, 0, 0, dateNow.Location()).Unix()
		endTime = time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 23, 59, 59, 0, dateNow.Location()).Unix()
	case "2": // 昨天
		dateNow := time.Now().AddDate(0, 0, -1)
		startTime = time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 0, 0, 0, 0, dateNow.Location()).Unix()
		endTime = time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 23, 59, 59, 0, dateNow.Location()).Unix()
	case "3": // 前天
		dateNow := time.Now().AddDate(0, 0, -2)
		startTime = time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 0, 0, 0, 0, dateNow.Location()).Unix()
		endTime = time.Date(dateNow.Year(), dateNow.Month(), dateNow.Day(), 23, 59, 59, 0, dateNow.Location()).Unix()
	case "4": // 指定日期范围
		dateNow := time.Now()
		startUnix, _ := time.ParseInLocation("2006-01-02 15:04:05", start, dateNow.Location())
		endUnix, _ := time.ParseInLocation("2006-01-02 15:04:05", end, dateNow.Location())
		startTime = time.Date(startUnix.Year(), startUnix.Month(), startUnix.Day(), 0, 0, 0, 0, startUnix.Location()).Unix()
		endTime = time.Date(endUnix.Year(), endUnix.Month(), endUnix.Day(), 23, 59, 59, 0, endUnix.Location()).Unix()
	}

	return
}

// GetLastDateFormat 获取昨天的时间
// GetLastDateFormat format "20060102" | "2006-01-02 15:04:05" | ""
func GetLastDateFormat(format string) string {
	startTime, _ := GetHourTime("2", "", "")
	if format == "" {
		format = DefaultTimeFormat
	}
	return time.Unix(startTime, 0).Format(format)
}

// GetNowTimeString 获取当前时间的字符串
func GetNowTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetNowDateFormat(format string) string {
	if format == "" {
		format = DefaultTimeFormat
	}
	return time.Now().Format(format)
}

// GetLastMonthStartEndDate 默认获取当前月份，如果往前的话就 -1
func GetLastMonthStartEndDate(m int) (time.Time, time.Time) {
	now := time.Now()
	lastMonthFirstDay := now.AddDate(0, m, -now.Day()+1)
	lastMonthStart := time.Date(lastMonthFirstDay.Year(), lastMonthFirstDay.Month(), lastMonthFirstDay.Day(), 0, 0, 0, 0, now.Location())
	lastMonthEndDay := lastMonthFirstDay.AddDate(0, 1, -1)
	lastMonthEnd := time.Date(lastMonthEndDay.Year(), lastMonthEndDay.Month(), lastMonthEndDay.Day(), 23, 59, 59, 0, now.Location())
	return lastMonthStart, lastMonthEnd
}

func GetLastTimeSting(year, month, day int, format string) string {
	now := time.Now()
	LastTime := now.AddDate(year, month, day)

	if format == "" {
		format = DefaultTimeFormat
	}

	return LastTime.Format(format)
}

// GetNowTimeUnix 【mod】 "s" , "ms" , "ns" "ws"
func GetNowTimeUnix(mod string) (UnixStr string, UnixInt int64) {
	switch mod {
	case "s":
		// 秒级别
		UnixInt = time.Now().Unix()
	case "ms":
		// 毫秒级别
		UnixInt = time.Now().UnixMilli()
	case "ns":
		// 纳秒级别
		UnixInt = time.Now().UnixNano()
	case "ws":
		// 微秒级别
		UnixInt = time.Now().UnixMicro()
	default:
		// 默认是S级别
		UnixInt = time.Now().Unix()
	}
	UnixStr = strconv.FormatInt(UnixInt, 10)
	return
}
