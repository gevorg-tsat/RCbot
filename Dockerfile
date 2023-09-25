FROM golang:1.21


RUN mkdir /code
WORKDIR /code

COPY . /code/