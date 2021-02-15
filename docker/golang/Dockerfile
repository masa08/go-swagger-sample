FROM golang:1.15.7-alpine

RUN apk add --update --no-cache \
  git \
  tar \
  postgresql-client

RUN go get -u github.com/oxequa/realize

RUN wget https://github.com/golang-migrate/migrate/releases/download/v4.12.2/migrate.linux-amd64.tar.gz && \
  tar -xvzf migrate.linux-amd64.tar.gz && \
  mv migrate.linux-amd64 bin/migrate && \
  rm migrate.linux-amd64.tar.gz

RUN wget https://taskfile.dev/install.sh && sh install.sh && rm install.sh

EXPOSE 80
CMD ["realize", "start"]