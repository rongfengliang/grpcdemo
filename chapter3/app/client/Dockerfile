# build stage
FROM golang:1.9-alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ENV D=/go/src/github.com/rongfengliang/grpcclient
ADD . $D
RUN cd $D && go build -o grpcclient && cp grpcclient /tmp/

FROM alpine:latest
WORKDIR /app
COPY --from=build-env /tmp/grpcclient /app/grpcclient
CMD ["./grpcclient"]
