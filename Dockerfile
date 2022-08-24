FROM debian:10-slim
WORKDIR /app
ADD ./scheduler-go /app

CMD ["/app/scheduler-go"]