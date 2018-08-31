package s3

import (
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
	"github.com/mboldysh/streaming-service/internal/model"
	"github.com/mboldysh/streaming-service/internal/store/s3/bucket"
)

type TrackStore struct {
	bucketName string
	client     *s3.S3
	uploader   *s3manager.Uploader
	expireTime time.Duration
}

func NewTrackStore(client *s3.S3, uploader *s3manager.Uploader, bucketName string, expireTime time.Duration) (*TrackStore, error) {

	_, err := bukcet.CreateIfDoesntExists(client, bucketName)

	if err != nil {
		return nil, err
	}

	return &TrackStore{
		bucketName: bucketName,
		client:     client,
		uploader:   uploader,
		expireTime: expireTime,
	}, nil
}

func (s *TrackStore) Upload(key string, file io.Reader) error {
	_, err := s.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(key),
		Body:   file,
	})

	return err
}

//FindAll objects in a bucket using pagination
func (s *TrackStore) FindAll(userID string) ([]model.Track, error) {
	var trackList []model.Track

	req := s.client.ListObjectsV2Request(&s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucketName),
		Prefix: aws.String(userID),
	})

	p := req.Paginate()

	for p.Next() {
		page := p.CurrentPage()
		for _, obj := range page.Contents {
			t := model.Track{
				Name: aws.StringValue(obj.Key),
				Size: aws.Int64Value(obj.Size),
			}
			trackList = append(trackList, t)
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return trackList, nil
}

func (s *TrackStore) GetPresignedURL(key, trackName string) (*model.PresignedTrack, error) {
	req := s.client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(key),
	})

	_, err := req.Send()

	if err != nil {
		return nil, err
	}

	url, header, err := req.PresignRequest(s.expireTime)

	if err != nil {
		return nil, err
	}

	return &model.PresignedTrack{
		Name:   trackName,
		URL:    url,
		Header: header,
	}, nil
}
