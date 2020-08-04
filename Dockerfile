FROM alpine:latest
RUN apk --no-cache add ca-certificates
USER 1001
COPY build/eventing-autoscaler-keda /app/eventing-autoscaler-keda
EXPOSE 18000 19000
ENTRYPOINT ["/app/eventing-autoscaler-keda"]