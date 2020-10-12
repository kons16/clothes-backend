package service

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
	"os"
)

// ランダムな文字列の生成
func MakeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error...")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}

// base64 のデータを s3 にアップロードして URL を返す
func UploadS3(imageBase64 string) string {
	if err := godotenv.Load(); err != nil {
		return ""
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(os.Getenv("S3AK"), os.Getenv("S3SK"), ""),
		Region:      aws.String("ap-northeast-1"),
	}))

	uploader := s3manager.NewUploader(sess)

	// TODO: base64先頭の image/jpeg などを取り除く(正規表現でやったほうがいい)
	imageBase64 = imageBase64[23:]

	data, _ := base64.StdEncoding.DecodeString(imageBase64)
	wb := new(bytes.Buffer)
	wb.Write(data)

	name, _ := MakeRandomStr(10)
	fileType := "image/jpeg"

	res, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("S3_KEY")),
		Key:         aws.String("clothes/" + name + ".jpeg"),
		Body:        wb,
		ContentType: &fileType,
	})

	if err != nil {
		fmt.Println(res)
	}

	url := os.Getenv("S3_URL") + name + ".jpeg"

	return url
}
