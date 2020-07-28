package main

import (
	"fmt"
	"time"
)

func Date2Time() {
	fmt.Println(">> Date2Time")
	defer fmt.Println("<< Date2Time")

	// 一定要以Go诞生的时间为基准
	// 2006年1月2号，MST时区，下午3:04分为基准
	const dateFormat = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(dateFormat, "May 20, 2020 at 0:00am (UTC)")
	fmt.Println(t)

	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2020-May-20")
	fmt.Println(t)

	t, _ = time.Parse("01/02/2006", "05/20/2020")
	fmt.Println(t)
}

func Time2Date() {
	fmt.Println(">> Time2Date")
	defer fmt.Println("<< Time2Date")

	tm := time.Now()
	fmt.Println(tm.Format("2006-01-02 03:04:05 PM"))
	fmt.Println(tm.Format("2006-1-2 03:04:05 PM"))
	fmt.Println(tm.Format("2006-Jan-02 03:04:05 PM"))
	fmt.Println(tm.Format("02/01/2006 03:04:05 PM"))
}

func Timestamp2Time() {
	fmt.Println(">> Timestamp2Time")
	defer fmt.Println("<< Timestamp2Time")

	ts := int64(1595900001)
	tm := time.Unix(ts, 0)
	fmt.Println(tm)
}

func Time2Timestamp() {
	fmt.Println(">> Time2Timestamp")
	defer fmt.Println("<< Time2Timestamp")

	tm := time.Now()
	ts := tm.Unix()
	fmt.Println(ts)
}

func main() {
	Date2Time()
	Time2Date()
	Timestamp2Time()
	Time2Timestamp()
}

/*
// 运行结果：
>> Date2Time
2020-05-20 00:00:00 +0000 UTC
2020-05-20 00:00:00 +0000 UTC
2020-05-20 00:00:00 +0000 UTC
<< Date2Time
>> Time2Date
2020-07-28 09:35:46 AM
2020-7-28 09:35:46 AM
2020-Jul-28 09:35:46 AM
28/07/2020 09:35:46 AM
<< Time2Date
>> Timestamp2Time
2020-07-28 09:33:21 +0800 CST
<< Timestamp2Time
>> Time2Timestamp
1595900146
<< Time2Timestamp
*/
