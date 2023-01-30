package s3

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/Vellvill/frames/internal/config"
	"github.com/Vellvill/frames/internal/logger"
)

type Uploader interface {
	Upload(ctx context.Context, object AWSObject) error
}

type uploader struct {
	bucket *string

	*s3manager.Uploader
}

func NewUploader() Uploader {
	s := session.Must(session.NewSession(&aws.Config{
		Region:     aws.String("msk"),
		MaxRetries: aws.Int(config.GetValue("s3_max_retries").Int()),
	}))
	s3manager.NewUploader(s)
	return uploader{
		aws.String(config.GetValue("s3_bucket").String()),
		s3manager.NewUploader(s),
	}
}

func (s uploader) Upload(ctx context.Context, object AWSObject) error {
	output, err := s.UploadWithContext(ctx, &s3manager.UploadInput{
		Body:   bytes.NewReader(object.Obj),
		Bucket: s.bucket,
		Key:    aws.String(fmt.Sprintf(string(key), object.ParentID, object.Sequence)),
	})

	logger.Logger.Debugf("%+v", output)

	return err
}
