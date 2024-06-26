FROM alpine:latest
RUN mkdir /app

WORKDIR /app
COPY . .



# Run the server executable
CMD [ "/app/myapp" ]
