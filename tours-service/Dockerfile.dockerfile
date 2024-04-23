FROM golang:alpine as build-stage
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o tours-service

FROM alpine
COPY --from=build-stage app/tours-service /usr/bin
EXPOSE 8082
ENTRYPOINT [ "tours-service" ]