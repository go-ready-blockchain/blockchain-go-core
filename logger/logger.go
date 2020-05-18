package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	FileName string
	NodeName string
)

// Logger is the strutcure of the log to be appended
type Logger struct {
	File      string    `json:"File"`
	Function  string    `json:"Function"`
	Line      int       `json:"Line"`
	Timestamp time.Time `json:"Timestamp"`
	Message   string    `json:"Message"`
	Node 			string 		`json:"Node"`
}

// CreateFile is used to create a new file
func CreateFile(name string) bool {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Error creating file")
		return false
	}
	defer file.Close()
	return true
}

// WriteToFile is to write data into a file
func WriteToFile(name string, body string) bool {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error opening file")
		return false
	}
	defer file.Close()

	var valueLogger = Logger{}
	valueLogger.Timestamp = time.Now().UTC()

	pc, callingFile, callingLine, _ := runtime.Caller(1)
	caller := runtime.FuncForPC(pc)

	valueLogger.File = callingFile
	valueLogger.Line = callingLine
	valueLogger.Message = body
	valueLogger.Function = caller.Name()
	valueLogger.Node = NodeName

	value, _ := json.Marshal(valueLogger)
	finalLog := string(value)
	fmt.Println(finalLog)
	
	finalLog += "\n"
	if _, err := file.WriteString(finalLog); err != nil {
		fmt.Println("Error writing file")
		return false
	}

	return true
}

// UploadToS3Bucket uploads the file to specified S3 bucket
func UploadToS3Bucket(name string, dir string) bool {
	file, err := os.Open(name)

	if err != nil {
		fmt.Println("Error opening file")
		return false
	}
	defer file.Close()

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), DisableSSL: aws.Bool(true),
	})

	svc := s3manager.NewUploader(sess)

	result, err := svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String("go-ready-blockchain"),
		Key:    aws.String(dir + "/" + filepath.Base(name)),
		Body:   file,
	})

	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	fmt.Println(result)

	return true
}
