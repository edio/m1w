FROM alpine:latest
ADD bin/m1ws /m1w
EXPOSE 80
ENTRYPOINT ["/m1w"]
