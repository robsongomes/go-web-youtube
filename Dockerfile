FROM golang:1.19 as build-site
ENV CGO_ENABLED 0

COPY . /app

WORKDIR /app

RUN go build -ldflags "-X main.env=prd" -o site *.go

FROM alpine:3.14
COPY --from=build-site /app/site /app/site

WORKDIR /app
CMD [ "./site" ]