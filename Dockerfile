FROM debian:10-slim
WORKDIR /app
ADD ./scheduler-go /app
ADD ./demodata /app/demodata

CMD ["/app/scheduler-go"]