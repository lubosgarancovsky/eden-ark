FROM alpine:latest
WORKDIR /app
COPY . .
RUN chmod +x eden-ark
EXPOSE 50002
CMD ["./eden-ark"]