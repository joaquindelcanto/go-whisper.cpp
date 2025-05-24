package main

import (
	"os"

	"github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
)

const (
	ENV_MODEL_PATH = "MODEL_PATH"
)

func loadModel() whisper.Model {

	// Load the model
	model, err := whisper.New(os.Getenv(ENV_MODEL_PATH))
	if err != nil {
		panic(err)
	}
	return model
}
