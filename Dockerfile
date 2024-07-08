FROM --platform=linux/amd64 golang:1.22

ENV HOME /home/local/
WORKDIR /home/local/src/

RUN apt-get update && apt-get upgrade -y

COPY */go.mod */go.sum ./
