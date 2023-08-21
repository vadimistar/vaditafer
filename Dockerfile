FROM --platform=amd64 golang:alpine as builder
WORKDIR /app
COPY go.mod go.sum  /app/
RUN go mod download
COPY . .
RUN go build -buildvcs=false -v -ldflags="-X 'main.Version=$VERSION'" -o app

CMD /app/app
