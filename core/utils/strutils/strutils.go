package strutils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"go-admin/core/global"
	"go-admin/core/utils/idgen"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// HidePartStr
/**
 * @Description: 字符串中间替换为*
 * @param value query string true "原始字符串"
 * @param n query int true "替换长度"
 * @return string
 */
func HidePartStr(value string, n int) string {
	if value == "" {
		return ""
	}
	startIndex := len(value)/2 - n/2
	replaceSymbol := "*"
	var builder strings.Builder
	for i, v := range value {
		if i >= startIndex-1 && i < startIndex+n {
			builder.WriteString(replaceSymbol)
		} else {
			builder.WriteString(string(v))
		}
	}
	return builder.String()
}

// IsNum 是否为整数
func IsNum(d decimal.Decimal) bool {
	if strings.Contains(d.String(), ".") {
		return false
	}
	return true
}

// GenerateValidateCode
// @Description: 随机生成6位数字验证码（crypto/rand，不可预测）
// @return string
func GenerateValidateCode() string {
	// 拒绝采样消除模偏差：只接受 [0, ceiling) 的值，保证 0-999999 均匀
	const max = 1000000
	const ceiling = (1 << 32) / max * max
	var n uint32
	for {
		b := make([]byte, 4)
		if _, err := rand.Read(b); err != nil {
			// crypto/rand 失败为致命错误（熵源不可用），直接暴露
			panic("GenerateValidateCode: crypto/rand failed: " + err.Error())
		}
		n = binary.BigEndian.Uint32(b)
		if n < ceiling {
			break
		}
	}
	return fmt.Sprintf("%06d", n%max)
}

// VerifyEmailFormat
// @Description: 检测邮箱格式
// @param email query string true "邮箱"
// @return bool
func VerifyEmailFormat(email string) bool {
	if email == "" {
		return false
	}
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func IsEmail(email string) bool {
	if email == "" {
		return false
	}
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	b := reg.MatchString(email)
	return b
}

// VersionOrdinal
// @Description:
// @param version query string true "版本号"
// @return string
// @return error
func VersionOrdinal(version string) (string, error) {
	// ISO/IEC 14651:2011
	const maxByte = 1<<8 - 1
	vo := make([]byte, 0, len(version)+8)
	j := -1
	for i := 0; i < len(version); i++ {
		b := version[i]
		if '0' > b || b > '9' {
			vo = append(vo, b)
			j = -1
			continue
		}
		if j == -1 {
			vo = append(vo, 0x00)
			j = len(vo) - 1
		}
		if vo[j] == 1 && vo[j+1] == '0' {
			vo[j+1] = b
			continue
		}
		if vo[j]+1 > maxByte {
			return "", errors.New("VersionOrdinal: invalid version")
		}
		vo = append(vo, b)
		vo[j]++
	}
	return string(vo), nil
}

type Address []byte

func IsMobile(mobile string) bool {
	//result, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, mobile)
	//涉及到各国，因此，只要判断长度和是否纯数字
	if len(mobile) < 5 {
		return false
	}
	result, err := regexp.MatchString(`^[-+]?(([0-9]+)([.]([0-9]+))?|([.]([0-9]+))?)$`, mobile)
	if err != nil {
		return false
	}
	if !result {
		return false
	}
	return true
}

func Hmac(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func IsStringEmpty(str string) bool {
	return strings.Trim(str, " ") == ""
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

func Base64ToImage(imageBase64 string) ([]byte, error) {
	image, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func GetDirFiles(dir string) ([]string, error) {
	dirList, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	filesRet := make([]string, 0)

	for _, file := range dirList {
		if file.IsDir() {
			files, err := GetDirFiles(dir + string(os.PathSeparator) + file.Name())
			if err != nil {
				return nil, err
			}

			filesRet = append(filesRet, files...)
		} else {
			filesRet = append(filesRet, dir+string(os.PathSeparator)+file.Name())
		}
	}

	return filesRet, nil
}

func GetCurrentTimeStamp() int64 {
	return time.Now().UnixNano() / 1e6
}

// slice去重
func RemoveRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

func CompareHashAndPassword(e string, p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		return false
	}
	return true
}

// GenerateMsgIDFromContext 生成msgID
func GenerateMsgIDFromContext(c *gin.Context) string {
	requestId := c.GetHeader(global.TrafficKey)
	if requestId == "" {
		requestId = idgen.UUID()
		c.Header(global.TrafficKey, requestId)
	}
	return requestId
}

func IntToString(e int) string {
	return strconv.Itoa(e)
}

func UIntToString(e uint) string {
	return strconv.Itoa(int(e))
}

func Int64ToString(e int64) string {
	return strconv.FormatInt(e, 10)
}

func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n // TODO +0.5 是为了四舍五入，如果不希望这样去掉这个
}

func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}

func StringToInt64(e string) (int64, error) {
	return strconv.ParseInt(e, 10, 64)
}

func GetCurrentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func StructToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

// InterfaceToInt64 辅助函数
func InterfaceToInt64(value interface{}) int64 {
	switch v := value.(type) {
	case float64:
		return int64(v)
	case string:
		var result int64
		fmt.Sscanf(v, "%d", &result)
		return result
	default:
		return 0
	}
}

func InterfaceToFloat(value interface{}) float64 {
	switch v := value.(type) {
	case float64:
		return v
	case string:
		var result float64
		fmt.Sscanf(v, "%f", &result)
		return result
	default:
		return 0
	}
}

func InterfaceToString(value interface{}) string {
	if str, ok := value.(string); ok {
		return str
	}
	return ""
}

func InterfaceToBool(value interface{}) bool {
	if v, ok := value.(bool); ok {
		return v
	}
	return false
}

// GetStringFromMap 安全的获取嵌套map值的辅助函数
func GetStringFromMap(m map[string]interface{}, keys ...string) string {
	if len(keys) == 0 {
		return ""
	}

	current := m
	for i, key := range keys {
		if i == len(keys)-1 {
			// 最后一个key，返回字符串
			if val, ok := current[key].(string); ok {
				return val
			}
			return ""
		}

		// 中间key，继续深入
		if next, ok := current[key].(map[string]interface{}); ok {
			current = next
		} else {
			return ""
		}
	}
	return ""
}
