FROM golang:1.13.4 AS build
ENV CGO_ENABLED 0
COPY . /go/src/marketboard-backend
WORKDIR /go/src/marketboard-backend
RUN go get -d ./...

# Install revel framework
RUN go get -u github.com/revel/revel
RUN go get -u github.com/revel/cmd/revel
#build revel app
RUN GOOS=linux GOARCH=amd64 revel build marketboard-backend backendbin


# Stage 2
FROM alpine:3.10.3

RUN mkdir -p /go/src/marketboard-backend/backendbin
WORKDIR /go/src
# We need bash to run our sh file.
RUN apk add --no-cache bash

COPY --from=build /go/src/marketboard-backend/backendbin /go/src/marketboard-backend

# Give full permissions for the files to be read by the script
RUN chmod -vR 777 /go/src/marketboard-backend

# We need this since the mongoDB server keeps complaining about invalid ca-certificates.
# So we kinda just need to update it.
RUN apk update \
    && apk upgrade \
    && apk add --no-cache \
    ca-certificates \
    && update-ca-certificates 2>/dev/null || true

CMD ["bash", "/go/src/marketboard-backend/run.sh"]

EXPOSE 8080