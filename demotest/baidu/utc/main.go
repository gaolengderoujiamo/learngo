package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	year, mon, day := now.UTC().Date()
	hour, min, sec := now.UTC().Clock()
	zone, _ := now.UTC().Zone()
	fmt.Printf("UTC 时间是 %d-%d-%d %02d:%02d:%02d %s\n",
		year, mon, day, hour, min, sec, zone)

	year, mon, day = now.Date()
	hour, min, sec = now.Clock()
	zone, _ = now.Zone()
	fmt.Printf("本地时间是 %d-%d-%d %02d:%02d:%02d %s\n",
		year, mon, day, hour, min, sec, zone)

	Bdate := "2014-06-24 14:30" //时间字符串

	t, _ := time.ParseInLocation("2006-01-02 15:04", Bdate, time.Local) //t被转为本地时间的time.Time
	fmt.Println(t)

	t, _ = time.Parse("2006-01-02 15:04", Bdate) //t被转为UTC时间的time.Time
	fmt.Println(t)

	t, _ = time.Parse(time.RFC3339, "2016-08-22T08:23:42Z")
	fmt.Println(t.In(time.Local))
}
