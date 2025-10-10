FROM alpine:latest
WORKDIR /app
COPY . .
RUN chmod +x eden-ark
EXPOSE 9090
CMD ["./eden-ark"]