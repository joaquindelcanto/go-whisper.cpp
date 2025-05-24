FROM ubuntu:noble-20250415.1

# Name of the Whisper model to be downloaded and used by this image
# https://huggingface.co/ggerganov/whisper.cpp
ARG model=base.en

# Install required tools
RUN apt-get update && apt-get install -y \
    build-essential \
    cmake \
    git \
    gcc-aarch64-linux-gnu \
    g++-aarch64-linux-gnu \
    && rm -rf /var/lib/apt/lists/*

# Download and install golang
ADD https://go.dev/dl/go1.24.3.linux-arm64.tar.gz .
RUN tar -C /usr/local -xzf go1.24.3.linux-arm64.tar.gz
RUN rm go1.24.3.linux-arm64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

# Download stable version (1.7.5) of whisper.cpp
ADD https://github.com/ggml-org/whisper.cpp/archive/refs/tags/v1.7.5.tar.gz .
RUN tar -C /usr/local -xzf v1.7.5.tar.gz
RUN rm v1.7.5.tar.gz

# Setting W for convenience - only used here in Dockerfile
ENV W=/usr/local/whisper.cpp-1.7.5

# Test and build golang bindings for whisper.cpp
WORKDIR ${W}/bindings/go
RUN make test
RUN make clean
RUN make whisper

# Set environment variables needed to run go code with go bindings for whisper.cpp

# whisper.h ggml.h
ENV C_INCLUDE_PATH=${W}/include:${W}/ggml/include

# libwhisper.a ggml
ENV LIBRARY_PATH=${W}/build_go/src:${W}/build_go/ggml/src

# only needed for darwin
# ENV GGML_METAL_PATH_RESOURCES=${W}

# Download Whisper model
ENV MODEL_DIR=/usr/local/whisper_models
WORKDIR ${MODEL_DIR}
ADD https://huggingface.co/ggerganov/whisper.cpp/resolve/main/ggml-${model}.bin .
ENV MODEL_PATH=${MODEL_DIR}/ggml-${model}.bin

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go test -v
CMD [ "go", "run", "." ]