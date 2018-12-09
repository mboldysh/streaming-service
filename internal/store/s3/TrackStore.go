package s3

import (
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
	"github.com/mboldysh/streaming-service/internal/model"
	"github.com/mboldysh/streaming-service/internal/store"
	"github.com/mboldysh/streaming-service/pkg/bucket"
)

type trackStore struct {
	bucketName string
	client     *s3.S3
	uploader   *s3manager.Uploader
	expireTime time.Duration
}

//NewTrackStore initialize a new TrackStore
func NewTrackStore(client *s3.S3, uploader *s3manager.Uploader, bucketName string, expireTime time.Duration) (store.TrackStore, error) {

	_, err := bucket.CreateIfDoesntExists(client, bucketName)

	if err != nil {
		return nil, err
	}

	return &trackStore{
		bucketName: bucketName,
		client:     client,
		uploader:   uploader,
		expireTime: expireTime,
	}, nil
}

//Upload uploads file to bucket with specified key
func (s *trackStore) Upload(key string, file io.Reader) error {
	_, err := s.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(key),
		Body:   file,
	})

	return err
}

//FindAll finds all objects in a bucket using pagination
func (s *trackStore) FindAll(userID string) ([]model.Track, error) {
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

//GetPresignedURL returns presignedURL for specified key
func (s *trackStore) GetPresignedURL(key, trackName string) (*model.PresignedTrack, error) {
	ok, err := bucket.ObjectExists(s.client, s.bucketName, key)

	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, fmt.Errorf("Object with key %s does not exists", key)
	}

	req := s.client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(key),
	})

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

//DeleteObject deletes object by key
func (s *trackStore) DeleteObject(key string) error {
	req := s.client.DeleteObjectRequest(&s3.DeleteObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(key),
	})

	_, err := req.Send()
	
	if err != nil {
		return err
	}

	return nil
}
