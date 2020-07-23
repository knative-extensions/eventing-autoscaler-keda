FROM alpine:latest
RUN apk --no-cache add ca-certificates
USER 1001
COPY build/autoscaler-keda /app/autoscaler-keda
EXPOSE 18000 19000
ENTRYPOINT ["/app/autoscaler-keda"]