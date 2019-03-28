FROM api-base:4ff4afb as build-env

WORKDIR /go/src/api-server

COPY . ./

RUN go install

FROM alpine:3.9
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

WORKDIR /usr/local/bin
COPY --from=build-env /go/bin/api-server /usr/local/bin/api-server
COPY ./docker/entrypoint.sh ./

CMD [ "./entrypoint.sh" ]