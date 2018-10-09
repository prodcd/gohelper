package helper

import (
	"bytes"
	"encoding/binary"
	"github.com/pkg/errors"
	"math"
	"strconv"
	"time"
)

func BytesToUint8(b []byte) uint8 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint8
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}
func BytesToInt32(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}
func DateTimeStringToBytes(dt string) []byte {
	t, _ := time.Parse("2006-01-02 15:04:05", dt)

	dateYear := uint32(t.Year() * 10000)
	dateMonth := uint32(t.Month() * 100)
	dateDay := uint32(t.Day())
	date := dateYear + dateMonth + dateDay
	//fmt.Println(date)
	timeHour := uint32(t.Hour() * 10000000)
	timeMinute := uint32(t.Minute() * 100000)
	timeSecond := uint32(t.Second() * 1000)
	timee := timeHour + timeMinute + timeSecond
	//fmt.Println(timee)
	var buf bytes.Buffer
	buf.Write(Uint32ToBytes(date))
	buf.Write(Uint32ToBytes(timee))
	//fmt.Println(buf.Bytes())
	return buf.Bytes()
}
func BytesToUint32(b []byte) uint32 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return uint32(x)
}

func BytesToInt16(b []byte) uint16 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint16
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return uint16(x)
}
func BytesToUint16(b []byte) uint16 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint16
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return uint16(x)
}

func Float32ToBytes(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, bits)

	return bytes
}

func Uint32ToBytes(u uint32) []byte {
	var b_buf *bytes.Buffer
	b_buf = bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.BigEndian, u)
	return b_buf.Bytes()
}
func Int32ToBytes(x int32) []byte {
	b_buf := bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.BigEndian, x)
	return b_buf.Bytes()
}
func IntToBytes(x int) []byte {
	return Int32ToBytes(int32(x))
}
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
func Int32StringToBytes(x string) []byte{
	outerr,_:=strconv.Atoi(x)
	b_buf := bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.BigEndian, outerr)
	return b_buf.Bytes()
}
func Int64ToBytes(x int64) []byte {
	b_buf := bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.BigEndian, x)
	return b_buf.Bytes()
}
func BytesToFloat32(bytes []byte) float32 {
	bits := binary.BigEndian.Uint32(bytes)

	return math.Float32frombits(bits)
}

func Float64ToBytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, bits)

	return bytes
}

func BytesToFloat64(bytes []byte) float64 {
	bits := binary.BigEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

func Uint32ToDateString(i uint32) string {
	s := strconv.Itoa(int(i))
	l := len(s)
	year := s[l-8 : l-4]
	month := s[l-4 : l-2]
	day := s[l-2 : l]
	return year + "-" + month + "-" + day
}
func Uint32ToTimeString(i uint32) string {
	s := strconv.Itoa(int(i))
	l := len(s)
	hour := s[l-9 : l-7]
	min := s[l-7 : l-5]
	sec := s[l-5 : l-3]
	return hour + ":" + min + ":" + sec
}
func StringToBytes(s string) []byte {
	return []byte(s)
}
func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
func BytesToInt64(b []byte) int64 {
	bytesBuffer := bytes.NewBuffer(b)
	var x int64
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}
func BytesToDateTimeString(b []byte) (string, error) {
	bytesBuffer := bytes.NewBuffer(b[0:4])
	var x uint32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	s := strconv.Itoa(int(x))
	l := len(s)
	year := s[l-8 : l-4]
	month := s[l-4 : l-2]
	day := s[l-2 : l]
	bytesBuffer2 := bytes.NewBuffer(b[4:8])
	var y uint32
	binary.Read(bytesBuffer2, binary.BigEndian, &y)
	s2 := strconv.Itoa(int(y+1000000000))
	l2 := len(s2)
	hour := s2[l2-9 : l2-7]
	min := s2[l2-7 : l2-5]
	sec := s2[l2-5 : l2-3]
	//判断格式错误
	montherr, err := strconv.Atoi(month)
	if montherr > 12 || err != nil {
		return "", errors.New("22003")
	}
	dayerr, err := strconv.Atoi(day)
	if (montherr == 1 && dayerr > 31) || (montherr == 2 && dayerr > 29) || (montherr == 3 && dayerr > 31) || (montherr == 4 && dayerr > 30) || (montherr == 5 && dayerr > 31) || (montherr == 6 && dayerr > 30) || (montherr == 7 && dayerr > 31) || (montherr == 8 && dayerr > 31) || (montherr == 9 && dayerr > 30) || (montherr == 10 && dayerr > 31) || (montherr == 11 && dayerr > 30) || (montherr == 12 && dayerr > 31) || err != nil {
		return "", errors.New("22003")
	}
	//TODO 闰年验证
	hourerr, err := strconv.Atoi(hour)
	if hourerr > 23 {
		return "", errors.New("22004")
	}
	minerr, err := strconv.Atoi(min)
	if minerr > 59 {
		return "", errors.New("22004")
	}
	secerr, err := strconv.Atoi(sec)
	if secerr > 59 {
		return "", errors.New("22004")
	}
	return year + "-" + month + "-" + day + " " + hour + ":" + min + ":" + sec, nil
}
func Join(s []string, separator string) string {
	r := ""
	j := len(s)-1
	for i, v := range s {
		if i != j {
			r += v + separator
		} else {
			r += v
		}
	}
	return r
}
func BytesToString(b []byte) string {
	return string(b)
}
func Uint32Join(s []uint32, separator string) string {
	r := ""
	j := len(s) - 1
	for i, v := range s {
		if i != j {
			r += Uint32ToString(v) + separator
		} else {
			r += Uint32ToString(v)
		}
	}
	return r
}
func Uint32ToString(u uint32) string {
	return strconv.FormatUint(uint64(u), 10)
}
