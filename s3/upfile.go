package s3

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

// S3PutObjectAPI defines the interface for the PutObject function.
// We use this interface to test the function using a mocked service.
type S3PutObjectAPI interface {
	PutObject(ctx context.Context,
		params *s3.PutObjectInput,
		optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}
type S3 struct {
	AppId     string
	AppSecret string
}

func NewS3() S3 {
	println("aws_config.access_key", viper.GetString("aws_config.access_key"))

	return S3{
		AppId:     viper.GetString("aws_config.access_key"),
		AppSecret: viper.GetString("aws_config.secret_key"),
	}
}

func (s S3) putFile(c context.Context, api S3PutObjectAPI, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return api.PutObject(c, input)
}

//RemoteUp 目前只试用小文件
func (s S3) RemoteUp(url string, bucket Bucket, name string) (uri string, err error) {
	if name == "" {
		name = uuid.New().String()
	}
	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		err = errors.New("received non 200 response code")
		return
	}
	fmt.Println("filenames:", name)
	var client = s.client()
	pix, err := ioutil.ReadAll(response.Body)
	var contentType = http.DetectContentType(pix)
	input := &s3.PutObjectInput{
		Bucket:      aws.String(string(bucket)),
		Key:         aws.String(name),
		ContentType: aws.String(contentType),
		Body:        bytes.NewReader(pix),
	}

	_, err = client.PutObject(context.TODO(), input)

	if err != nil {
		fmt.Println("Got error uploading file-err:", err)
	}
	uri = bucket.Uri(name)
	return
}

func (s S3) UpFile(filename string, bucket Bucket, key string) (uri string, err error) {

	var (
		//file *os.File
		buf []byte
	)
	buf, err = ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Unable to ReadFile file " + filename)
		return
	}
	var contentType = http.DetectContentType(buf)

	//file, err = os.Open(filename)
	//if err != nil {
	//	fmt.Println("Unable to open file " + filename)
	//	return
	//}
	//defer file.Close()

	input := &s3.PutObjectInput{
		Bucket:      aws.String(string(bucket)),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
		Body:        bytes.NewReader(buf),
	}
	var client = s.client()
	_, err = client.PutObject(context.TODO(), input)
	if err != nil {
		fmt.Println("Got error uploading file-err:", err)
	}
	uri = BucketDomain(bucket) + "/" + key
	return
}

func (s S3) client() (client *s3.Client) {
	var (
		cfg aws.Config
		err error
	)
	cfg, err = config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	cfg.Region = "us-east-1"
	var credentials = credentials.NewStaticCredentialsProvider(s.AppId, s.AppSecret, "")
	cfg.Credentials = credentials
	return s3.NewFromConfig(cfg)

}

//PreSignUri 临时上传uri
func (s S3) PreSignUri(bucket Bucket, name, fileType string) (preUri string, viewUrl string, err error) {
	if name == "" {
		name = "pre_upload/" + uuid.New().String()
	}
	println("filenames:", name)

	input := &s3.PutObjectInput{
		Bucket:      aws.String(string(bucket)),
		Key:         aws.String(name),
		ContentType: aws.String(fileType),
		//ContentMD5:  aws.String(md5),
		//Expires: aws.Time(time.Now().Add(time.Minute * 10)),
	}
	resp, err := s3.NewPresignClient(s.client()).PresignPutObject(context.TODO(), input)
	if err != nil {
		fmt.Println("Got an error retrieving pre-signed object:")
		fmt.Println(err)
		return
	}
	preUri = resp.URL
	viewUrl = bucket.Uri(name)
	return
}
