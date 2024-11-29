package main

import interface_ "interface-oriented/interface"

func main() {
	job := interface_.ImageProcessingJob{
		ImageStore: &interface_.PrivateImageStore{},
		BucketName: "ai_images_bucket",
	}

	image := interface_.Image{Data: []byte("image data"), Name: "example.jpg"}
	job.Process(image)
}
