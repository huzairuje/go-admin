package utils

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go-admin/infrastructure/session"
	models "go-admin/modules/primitive/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	CorrelationID  string = "X-Correlation-ID"
	ErrorLogFormat        = "got err: %v, context: %s - %s"
)

func IsValidSanitizeSQL(queryParam string) bool {
	regexQueryParam := regexp.MustCompile(`^[\w ]+$`)
	return regexQueryParam.MatchString(queryParam)
}

func IpExtractor(c echo.Context) string {
	realIP := c.RealIP()
	if realIP != "" {
		return realIP
	}
	return strings.Split(c.Request().RemoteAddr, ":")[0]
}

func StringToUnixTimestamp(input time.Duration, unit string) int64 {
	var output int64
	switch unit {
	case "minutes":
		output = time.Now().Add(input * time.Minute).Unix()
	case "hours":
		output = time.Now().Add(input * time.Hour).Unix()
	case "days":
		output = time.Now().Add(input * time.Hour * 24).Unix()
	default:
		output = time.Now().Add(5 * time.Hour).Unix()
	}
	return output
}

func StringUnitToDuration(input string) time.Duration {
	durationMapping := map[string]time.Duration{
		"second": time.Second,
		"minute": time.Minute,
		"hour":   time.Hour,
		"day":    24 * time.Hour,
		"week":   7 * 24 * time.Hour,
		"month":  30 * 24 * time.Hour,
		"year":   365 * 24 * time.Hour,
	}

	switch input {
	case "second":
		return durationMapping["second"]
	case "minute":
		return durationMapping["minute"]
	case "hour":
		return durationMapping["hour"]
	case "week":
		return durationMapping["week"]
	case "month":
		return durationMapping["month"]
	case "year":
		return durationMapping["year"]
	default:
		return time.Second
	}
}

func IsErrorContain(err error, sliceError []error) bool {
	for _, e := range sliceError {
		if errors.Is(e, err) {
			return true
		}
	}
	return false
}

func UniqueInt64(intSlice []int64) (uniqueUint64 []int64, lenUniqueUint int64) {
	keys := make(map[int64]bool)
	var list []int64
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list, int64(len(list))
}

func GenerateTimeNow() string {
	//fetching current time
	loc, _ := time.LoadLocation("Asia/Jakarta")
	currentTime := time.Now().In(loc).Format(time.RFC3339)
	//differnce between pastdate and current date
	return currentTime
}

func GetUserInfoFromContext(ctx echo.Context, db *gorm.DB) (userModel models.User, err error) {
	result, err := session.Manager.Get(ctx, session.SessionId)
	if err != nil {
		panic(err)
	}
	userInfo := result.(session.UserInfo)
	err = db.Model(models.User{}).Preload(clause.Associations).Where("user_id =?", userInfo.UserID).First(&userModel).Error
	return
}

func ItemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)
	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}
	return false
}

func GetAmountFromRupiah(rupiah string) int64 {
	re := regexp.MustCompile("[0-9]+")
	getNumber := re.FindAllString(rupiah, -1)
	amountString := strings.Join(getNumber, "")
	amount, err := strconv.ParseInt(amountString, 10, 64)
	if err != nil {
		return 0
	}
	return amount
}

func UploadImageV2(src string) (string, error) {
	if src == "" {
		return "", errors.New("image base64 is null")
	}
	decode, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}
	ext := http.DetectContentType(decode)
	if ext == "text/plain; charset=utf-8" {
		return "", errors.New("cannot get extension image")
	}
	res := strings.Split(ext, "image/")

	fName := uuid.New().String() + "." + res[1]
	r := bytes.NewReader(decode)
	dir, _ := os.Getwd()
	fileLocation := filepath.Join(dir, "assets/upload/avatars", fName)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, r); err != nil {
		return "", err
	}
	return "/assets/upload/avatars/" + fName, err
}

func ConfigTime(tm int) (date string) {
	t := time.Now()
	convert := t.AddDate(0, 0, -tm)
	date = convert.Format("2006-01-02")
	return date
}

func StrToUint64(str string) (uint64, error) {
	i, err := strconv.ParseInt(str, 10, 64)
	return uint64(i), err
}

func Contains(elems []string, elem string) bool {
	for _, e := range elems {
		if elem == e {
			return true
		}
	}
	return false
}
