package converter

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
)

func Converter(fileName string) (outputFile string, err error) {
	outputFile = fileName + ".mp3"
	err = ffmpeg.Input(fileName).Output(outputFile).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		log.Println(err)
	}

	return outputFile, err
}
