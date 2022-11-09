FROM golang:1.18
WORKDIR /src/afyadigital/<PROJECT_NAME>
COPY . .
RUN mkdir -p /build/bin && CGO_ENABLED=0 GOOS=linux go build -o ./build/bin/ ./cmd/...
RUN ls -lhtr ./build/bin


FROM alpine:latest
# download certificate
RUN apk --no-cache add ca-certificates tzdata
COPY --from=0 /src/afyadigital/<PROJECT NAME>/build/bin/ .
CMD ["./<PROJECT_NAME>"]
