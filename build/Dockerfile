
FROM golang:1.22-alpine AS build
LABEL authors="rafaeltomelin"

WORKDIR /app
RUN apk update && apk add --no-cache git

# Create nonroot user
RUN adduser --disabled-password -u 10001 appuser

ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV CGO_ENABLED=0

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

WORKDIR /app/
RUN go build -ldflags="-w -s" -o /app/api cmd/*

##
## Deploy
FROM scratch

WORKDIR /app

COPY --from=build /app/api /app
COPY --from=build /app/cmd/.config/config.yaml /app/cmd/.config/

EXPOSE 8443
EXPOSE 8082
EXPOSE 50051

# copy userus from build
COPY --from=build /etc/passwd /etc/passwd
USER appuser

ENTRYPOINT ["/app/api"]