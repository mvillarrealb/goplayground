FROM golang:1.15.8-alpine3.13
RUN mkdir -p /var/lib/app
WORKDIR /var/lib/app
COPY *.go .
CMD ["go", "run", "."]