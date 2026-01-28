FROM alpine:3.23.3
COPY gh-app-access-token /usr/local/bin/gh-app-access-token
ENTRYPOINT ["/usr/local/bin/gh-app-access-token"]
