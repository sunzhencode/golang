package main

import (
	"fmt"
	"time"
)

var (
	loc *time.Location
)

const (
	TIME_FMT = "2006-01-02 15:04:05"
)

func init() {
	loc, _ = time.LoadLocation("Asia/Shanghai")
}

//把字符串1998-10-01 08:10:00解析成time.Time，再格式化成字符串199810010810
func timeFmt() {
	t, _ := time.ParseInLocation(TIME_FMT, "1998-10-01 08:10:00", loc)
	fmt.Println(t)
	tf := t.Format("200601021504")
	fmt.Println(tf)
	ta, _ := time.Parse(TIME_FMT, TIME_FMT)
	fmt.Println(int(ta.Weekday())) //2006-01-02是星期一
	fmt.Println(time.UnixDate)     //自带的时间格式

}

//输出未来4个周六日期（不考虑法定假日）
func printSix() {

}

//把一个目录下的所有.txt文件合一个大的.txt文件，再对这个大文件进行压缩
func sumFile() {

}

func main() {
	timeFmt()
}
