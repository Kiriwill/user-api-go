FROM golang:1.18
WORKDIR /src/verifymydesafio
COPY . .
RUN mkdir -p /build/bin && CGO_ENABLED=0 GOOS=linux go build -o ./build/bin/ ./cmd/...
RUN ls -lhtr ./build/bin
