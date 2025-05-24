package main

import (
	"fmt"
	"io"
	"testing"

	// Package imports
	whisper "github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
)

func TestWhisper(t *testing.T) {

	var context whisper.Context
	{
		// Load the model
		model := loadModel()
		defer model.Close()

		// Create a new context
		var err error
		context, err = model.NewContext()
		if err != nil {
			t.Fatalf("Failed to create new context: %v", err)
			return
		}
	}

	// Set the parameters
	context.SetLanguage("en")
	context.SetTranslate(false)
	context.SetThreads(4)
	context.SetMaxSegmentLength(20)

	fmt.Printf("\n%s\n", context.SystemInfo())

	// Load audio file, get samples
	const wavFilePath = "./testdata/jfk.wav"
	fmt.Printf("Loading %q\n", wavFilePath)
	samples, err := loadSamplesFromWavFile(wavFilePath)
	if err != nil {
		t.Fatalf("Failed to load samples from wav file: %v", err)
		return
	}

	// Process the data
	_testWhisperProcessing(t, context, samples)
}

func _testWhisperProcessing(t *testing.T, context whisper.Context, samples []float32) {
	// Process the data
	fmt.Print("  ...processing")
	context.ResetTimings()
	if err := context.Process(samples, nil, nil, nil); err != nil {
		t.Fatalf("Failed to process samples: %v", err)
		return
	}

	context.PrintTimings()

	// Print out the results
	for {
		segment, err := context.NextSegment()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Fatalf("Failed to get next segment: %v", err)
			return
		}
		fmt.Printf("[%6s->%6s] %s\n", segment.Start, segment.End, segment.Text)
	}
}
