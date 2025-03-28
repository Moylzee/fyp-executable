package utils

import (
	"bytes"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var sess *session.Session
var svc *s3.S3

func init() {
    // Initialize the AWS session
    var err error
    sess, err = session.NewSession(&aws.Config{
        Region: aws.String("eu-north-1"), // Specify your region
    })
    if err != nil {
        log.Fatal("Unable to initialize AWS session: ", err)
    }
    svc = s3.New(sess)
}

func UploadFileToS3(fileName string, content []byte, bucket string) error {
    // Create the PutObject request
    input := &s3.PutObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(fileName),
        Body:   bytes.NewReader(content),
    }

    _, err := svc.PutObject(input)
    if err != nil {
        return fmt.Errorf("unable to upload file to S3: %v", err)
    }
    return nil
}