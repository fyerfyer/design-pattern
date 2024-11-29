package interface_

import "fmt"

type Image struct {
	Data []byte
	Name string
}

// interface for image storage.
type ImageStore interface {
	Upload(image Image, bucketName string) string
	Download(url string) Image
}

// AliyunImageStore implementation
type AliyunImageStore struct{}

func (a *AliyunImageStore) Upload(image Image, bucketName string) string {
	a.createBucketIfNotExisting(bucketName)
	accessToken := a.generateAccessToken()
	fmt.Println("Uploading to Aliyun with access token:", accessToken)
	return fmt.Sprintf("https://aliyun.com/%s/%s", bucketName, image.Name)
}

func (a *AliyunImageStore) Download(url string) Image {
	accessToken := a.generateAccessToken()
	fmt.Println("Downloading from Aliyun with access token:", accessToken)
	return Image{Data: []byte("image data"), Name: "downloaded_image"}
}

func (a *AliyunImageStore) createBucketIfNotExisting(bucketName string) {
	fmt.Println("Creating bucket:", bucketName)
}

func (a *AliyunImageStore) generateAccessToken() string {
	return "aliyun_access_token"
}

// PrivateImageStore implementation
type PrivateImageStore struct{}

func (p *PrivateImageStore) Upload(image Image, bucketName string) string {
	p.createBucketIfNotExisting(bucketName)
	fmt.Println("Uploading to private cloud")
	return fmt.Sprintf("https://privatecloud.com/%s/%s", bucketName, image.Name)
}

func (p *PrivateImageStore) Download(url string) Image {
	fmt.Println("Downloading from private cloud")
	return Image{Data: []byte("image data"), Name: "downloaded_image"}
}

func (p *PrivateImageStore) createBucketIfNotExisting(bucketName string) {
	fmt.Println("Creating bucket:", bucketName)
}

// use the ImageProcessJob struct to do dependency injection
type ImageProcessingJob struct {
	ImageStore ImageStore
	BucketName string
}

func (job *ImageProcessingJob) Process(image Image) {
	fmt.Println("Processing image:", image.Name)
	url := job.ImageStore.Upload(image, job.BucketName)
	fmt.Println("Image uploaded to URL:", url)
}
