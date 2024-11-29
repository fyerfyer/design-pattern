package main

func main() {
	job := ImageProcessingJob{
		ImageStore: &PrivateImageStore{},
		BucketName: "ai_images_bucket",
	}

	image := Image{Data: []byte("image data"), Name: "example.jpg"}
	job.Process(image)
}
