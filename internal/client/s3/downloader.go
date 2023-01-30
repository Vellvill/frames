package s3

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/Vellvill/frames/internal/config"
	"github.com/Vellvill/frames/internal/logger"
)

type Downloader interface {
	Download(ctx context.Context, key awsKey) (AWSObject, error)
}

type downloader struct {
	bucket *string

	*s3manager.Downloader
}

func NewDownloader() Downloader {
	s := session.Must(session.NewSession(&aws.Config{
		Region:     aws.String("msk"),
		MaxRetries: aws.Int(config.GetValue("s3_max_retries").Int()),
	}))
	s3manager.NewDownloader(s)
	return downloader{
		aws.String(config.GetValue("s3_bucket").String()),
		s3manager.NewDownloader(s),
	}
}

func (s downloader) Download(ctx context.Context, key awsKey) (AWSObject, error) {

	parentID, sequence, err := key.getParentIDAndSequenceFromKey()
	if err != nil {
		return AWSObject{}, err
	}

	res := AWSObject{
		ParentID: parentID,
		Sequence: sequence,
	}

	output, err := s.DownloadWithContext(ctx, aws.NewWriteAtBuffer(res.Obj), &s3.GetObjectInput{
		Bucket: s.bucket,
		Key:    aws.String(string(key)),
	})

	logger.Logger.Debugf("%+v", output)

	return res, err
}
