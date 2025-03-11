FROM debian:stable-slim
COPY go-server /bin/go-server
CMD ["/bin/go-server"]
ENV PORT=8080
