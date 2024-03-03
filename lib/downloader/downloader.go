package downloader

import (
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

func Donwload(link string) (outputFile string, err error) {
	videoID := link
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	fileName := "./storage/" + video.Title
	file, err := os.Create(fileName + ".mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}

	return fileName, err
}
