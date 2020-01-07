# FROM golang:alpine as builder
# RUN apk --no-cache add git
# WORKDIR /root
# COPY web-download.go .

# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o web-download .
# FROM alpine:latest as prod
# RUN apk --no-cache add ca-certificates curl ffmpeg python3 && pip3 install --no-cache-dir you-get 
# WORKDIR /root/
# RUN mkdir /root/data
# EXPOSE 8080
# COPY --from=0 /root/web-download .
# CMD ["./web-download"]



#########
FROM golang:alpine as builder
WORKDIR /root
COPY web-download.go .
RUN go build -o web-download .

FROM python:3.6-alpine as prod
RUN apk --no-cache add ffmpeg && pip3 install --no-cache-dir you-get && mkdir /download && rm -rf /var/cache/apk/* && rm -rf .cache/pip
WORKDIR /root
COPY --from=0 /root/web-download .
EXPOSE 8080

CMD ["./web-download"]