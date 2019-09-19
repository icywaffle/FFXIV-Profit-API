FROM alpine:latest

RUN mkdir -p /go/src/marketboard-backend
WORKDIR /go/src
# We need bash to run our sh file.
RUN apk add --no-cache bash

COPY backendbin /go/src/marketboard-backend

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

EXPOSE 9000