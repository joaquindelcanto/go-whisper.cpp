package main

import (
	"fmt"
	"os"

	"github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

func loadSamplesFromWavFile(filePath string) (samples []float32, err error) {
	// Open the WAV file
	fh, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	// Decode the WAV file - load the full buffer
	dec := wav.NewDecoder(fh)
	var buf *audio.IntBuffer
	if buf, err = dec.FullPCMBuffer(); err != nil {
		return nil, err
	} else if dec.SampleRate != whisper.SampleRate {
		return nil, fmt.Errorf("unsupported sample rate: %d", dec.SampleRate)
	} else if dec.NumChans != 1 {
		return nil, fmt.Errorf("unsupported number of channels: %d", dec.NumChans)
	} else {
		samples = buf.AsFloat32Buffer().Data
		return samples, nil
	}
}
