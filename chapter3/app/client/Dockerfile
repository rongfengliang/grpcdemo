# build stage
FROM golang:1.9-alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ENV D=/go/src/github.com/rongfengliang/grpclient
ADD . $D
RUN cd $D && go build -o grpclient && cp grpclient /tmp/

FROM alpine:latest
WORKDIR /app
COPY --from=build-env /tmp/grpclient /app/grpclient
CMD ["./grpclient"]
