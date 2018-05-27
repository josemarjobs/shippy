FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
ADD user-cli /app/user-cli

CMD [ "./user-cli" ]