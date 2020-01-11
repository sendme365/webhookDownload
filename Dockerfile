FROM golang:alpine as builder

WORKDIR /root
COPY web-download.go .
RUN go build -o web-download .

FROM python:3.6-alpine as prod
LABEL maintainer="Ted" 
RUN apk --no-cache add ffmpeg && pip3 install --no-cache-dir you-get && mkdir /download && rm -rf /var/cache/apk/* && rm -rf .cache/pip
WORKDIR /root
COPY --from=0 /root/web-download .
EXPOSE 3000 

CMD ["./web-download"]
