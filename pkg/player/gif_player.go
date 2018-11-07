package player

import (
	"fmt"
	"github.com/qeesung/asciiplayer/pkg/decoder"
	"github.com/qeesung/image2ascii/convert"
	"log"
	"os"
	"time"
)

// GifTerminalPlayer is terminal player that implement the Player interface
type GifTerminalPlayer struct {
	decoder   decoder.Decoder
	converter *convert.ImageConverter
}

// NewGifTerminalPlayer create a new gif terminal player
func NewGifTerminalPlayer() Player {
	return &GifTerminalPlayer{
		decoder:   decoder.NewGifDeCoder(),
		converter: convert.NewImageConverter(),
	}
}

// Play decode the gif file content then play it in the terminal
func (terminalPlayer *GifTerminalPlayer) Play(filename string, playOptions *PlayOptions) {
	// decode the file first
	frames, err := terminalPlayer.decoder.DecodeFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	convertOptions := convert.DefaultOptions
	convertOptions.StretchedScreen = playOptions.StretchedScreen
	convertOptions.Reversed = playOptions.Reversed
	convertOptions.FitScreen = playOptions.FitScreen
	convertOptions.Ratio = playOptions.Ratio
	convertOptions.FixedWidth = playOptions.FixedWidth
	convertOptions.FixedHeight = playOptions.FixedHeight
	convertOptions.Colored = playOptions.Colored
	delay := playOptions.Delay

	stdout := os.Stdout
	for {
		for _, frame := range frames {
			asciiImageStr := terminalPlayer.converter.Image2ASCIIString(frame, &convertOptions)
			fmt.Fprint(stdout, asciiImageStr)
			time.Sleep(delay)
			fmt.Fprint(stdout, clearScreen)
		}
	}
}
