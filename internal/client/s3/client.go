package s3

import (
	"bytes"
	"context"
	"github.com/Vellvill/frames/internal/config"
	"github.com/Vellvill/frames/internal/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Client interface {
	Put(ctx context.Context, object AWSObject) error
}

type client struct {
	storage *s3.S3
	bucket  *string
}

func New() Client {
	s := session.Must(session.NewSession(&aws.Config{
		Region:     aws.String("msk"),
		MaxRetries: aws.Int(config.GetValue("s3_max_retries").Int()),
	}))

	return client{s3.New(s), aws.String(config.GetValue("s3_bucket").String())}
}

func (s client) Put(ctx context.Context, object AWSObject) error {
	output, err := s.storage.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: s.bucket,
		Key:    aws.String(object.Key),
		Body:   bytes.NewReader(object.Obj),
	})

	logger.Logger.Debug(output.String())

	return err
}

func (s client) Get(ctx context.Context, key string) (AWSObject, error) {
	output, err := s.storage.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: s.bucket,
		Key:    aws.String(key),
	})
	if err != nil {

	}

	logger.Logger.Debug(output.String())

	buf := make([]byte, *output.ContentLength)

	_, err = output.Body.Read(buf)

	return AWSObject{
		Key: key,
		Obj: buf,
	}, nil
}
