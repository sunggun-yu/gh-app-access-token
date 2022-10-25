FROM alpine
COPY gh-app-access-token /usr/local/bin/gh-app-access-token
ENTRYPOINT ["/usr/local/bin/gh-app-access-token"]
