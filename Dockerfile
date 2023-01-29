FROM golang:1.19 AS build
RUN go install github.com/go-task/task/v3/cmd/task@latest

WORKDIR /app
COPY . .
RUN task build

FROM alpine:latest AS app
RUN apk --no-cache --update add ca-certificates
COPY --from=build /app/plagiarism-backend /
ENTRYPOINT ["/plagiarism-backend"]
