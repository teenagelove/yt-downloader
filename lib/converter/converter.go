package converter

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
)

func Converter(fileName string) {
	err := ffmpeg.Input(fileName).Output(fileName + ".mp3").
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		log.Println(err)
	}
}
