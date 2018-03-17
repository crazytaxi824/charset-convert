package charset

import (
	"bytes"
	"errors"
	"io/ioutil"
	"unicode/utf16"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// GBKtoUTF8 字符转换
func GBKtoUTF8(gbkData []byte) (utf8Data []byte, err error) {
	reader := transform.NewReader(bytes.NewReader(gbkData), simplifiedchinese.GBK.NewDecoder())
	utf8Data, err = ioutil.ReadAll(reader)
	return
}

// UTF8toGBK 字符转换
func UTF8toGBK(utf8Data []byte) (gbkData []byte, err error) {
	reader := transform.NewReader(bytes.NewReader(utf8Data), simplifiedchinese.GBK.NewEncoder())
	gbkData, err = ioutil.ReadAll(reader)
	return
}

// GB18030toUTF8 字符转换
func GB18030toUTF8(gb18030Data []byte) (utf8Data []byte, err error) {
	reader := transform.NewReader(bytes.NewReader(gb18030Data), simplifiedchinese.GB18030.NewDecoder())
	utf8Data, err = ioutil.ReadAll(reader)
	return
}

// UTF8toGB18030 字符转换
func UTF8toGB18030(utf8Data []byte) (gb18030Data []byte, err error) {
	reader := transform.NewReader(bytes.NewReader(utf8Data), simplifiedchinese.GB18030.NewEncoder())
	gb18030Data, err = ioutil.ReadAll(reader)
	return
}

// GB2312toUTF8 字符转换
func GB2312toUTF8(gb2312Data []byte) (utf8Data []byte, err error) {
	reader := transform.NewReader(bytes.NewReader(gb2312Data), simplifiedchinese.HZGB2312.NewDecoder())
	utf8Data, err = ioutil.ReadAll(reader)
	return
}

// UTF8toGB2312 字符转换
func UTF8toGB2312(utf8Data []byte) (gb2312Data []byte, err error) {
	reader := transform.NewReader(bytes.NewReader(utf8Data), simplifiedchinese.HZGB2312.NewEncoder())
	gb2312Data, err = ioutil.ReadAll(reader)
	return
}

// UTF16toUTF8 字符转换
func UTF16toUTF8(b []byte) ([]byte, error) {
	if len(b)%2 != 0 {
		return nil, errors.New("[]byte 字符长度必须为双数, must have even length byte slice")
	}

	u16s := make([]uint16, 1)
	result := &bytes.Buffer{}
	b8buf := make([]byte, 4)

	lb := len(b)
	for i := 0; i < lb; i += 2 {
		u16s[0] = uint16(b[i]) + (uint16(b[i+1]) << 8)
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		result.Write(b8buf[:n])
	}
	return result.Bytes(), nil
}
