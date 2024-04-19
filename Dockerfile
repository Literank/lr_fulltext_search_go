# alpine linux
FROM alpine:3.19

ENV APP_BIN=lr_ft_books
ARG SERVER_DIR=/home/.server
WORKDIR $SERVER_DIR
COPY ./bin/${APP_BIN} .

ENV GIN_MODE=release

CMD ./${APP_BIN}
