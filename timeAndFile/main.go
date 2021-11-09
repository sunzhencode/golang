package main

import (
	"bufio"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
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
	today := time.Now()
	sevenday := time.Duration(time.Hour * 24 * 7)
	t := today.Add(sevenday)
	fmt.Println(t.Format(TIME_FMT))
}

func writeAbc(outpath, line string) {
	//打开写操作的文件
	fout, err := os.OpenFile(outpath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fout.Close()
	writer := bufio.NewWriter(fout)
	writer.WriteString(line)
	writer.WriteString("\n")
	writer.Flush()
}

func readAbc(path string) (s string) {

	//打开读操作的文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		s = fmt.Sprintf("%s%s", s, line)
		//line = strings.TrimRight(line, "\n")
		//fmt.Println(line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("Read file error!", err)
			}
		}
	}
	return
	//writer.WriteString("\n")
	//writer.Flush()
	// 按行读取的方法
	//for {
	//	if line, _,err := reader.ReadLine(); err != nil {
	//		if err == io.EOF {
	//			break
	//		}
	//	} else {
	//		fmt.Println(string(line))
	//	}
	//}
}

func compress(filename, comfile string) {
	fin, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fout, err := os.OpenFile(comfile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	bs := make([]byte, 10)
	writer := zlib.NewWriter(fout) //压缩写入
	for {
		n, err := fin.Read(bs)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
			}
		} else {
			writer.Write(bs[:n])
		}
	}
	writer.Close()
	fout.Close()
	fin.Close()
	fin, err = os.Open(comfile)
	if err != nil {
		fmt.Println(err)
		return
	}
	stat, _ := fin.Stat()
	fmt.Printf("压缩后文件大小 %dB\n", stat.Size())

	reader, err := zlib.NewReader(fin) //解压
	io.Copy(os.Stdout, reader)         //把一个流拷贝到另外一个流
	reader.Close()
	fin.Close()
}

//把一个目录下的所有.txt文件合一个大的.txt文件，再对这个大文件进行压缩
func sumFile() {
	outfile := "abc.txt"
	os.Remove(outfile)
	if fileinfos, err := ioutil.ReadDir("."); err == nil {
		for _, fileinfo := range fileinfos {
			filename := fileinfo.Name()
			if strings.HasSuffix(filename, ".txt") {
				outline := readAbc(filename)
				writeAbc(outfile, outline)
			}
		}
	}
	compress(outfile, "abc.zlib")
}

func main() {
	printSix()
}
