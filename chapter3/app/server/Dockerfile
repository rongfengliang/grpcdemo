# build stage
FROM golang:1.9-alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ENV D=/go/src/github.com/rongfengliang/grpserver
ADD . $D
RUN cd $D && go build -o grpserver && cp grpserver /tmp/

FROM debian:jessie-slim
WORKDIR /app
EXPOSE 9080
COPY --from=build-env /tmp/grpserver /app/grpserver
CMD ["./grpserver"]
