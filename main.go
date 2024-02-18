package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
)

func main() {
	ctx := context.TODO()
	cfg := loadAws()
	rek := rekognition.NewFromConfig(cfg)
	//b64 := loadBase64Image()
	file, err := os.ReadFile("./img.png")
	output, err := rek.DetectText(ctx, &rekognition.DetectTextInput{
		Image: &types.Image{
			Bytes: file,
		},
	})
	if err != nil {
		panic(err)
	}

	for _, text := range output.TextDetections {
		fmt.Println(*text.DetectedText)
	}
}

func loadAws() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	return cfg
}
