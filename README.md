# go-whisper.cpp
Ready-to-go development environment for a writing a [Go](https://go.dev/) program using the [whisper.cpp Go bindings](https://github.com/ggml-org/whisper.cpp/tree/master/bindings/go), all running inside of a [Docker](https://www.docker.com/) container.

## Requirements
- [Go](https://go.dev/learn/)
- [Docker](https://www.docker.com/get-started/)

## Example Usage
```
docker build -t go-whisper.cpp .
docker run go-whisper.cpp
```

By default, the `base.en` whisper model will be downloaded during `docker build` and included in the built image.  To change this, use the `model` build-arg:

```
docker build -t go-whisper.cpp --build-arg model=large-v3-turbo .
docker run go-whisper.cpp
```

[List of available models.](https://huggingface.co/ggerganov/whisper.cpp)

Happy coding! 🙏😎
