FROM ubuntu:20.04

RUN apt-get update && apt-get install -y --no-install-recommends apt-utils

RUN apt-get install golang -y

RUN apt-get install -y --no-install-recommends python3 python3-pip python3-setuptools python3-dev
RUN python3 -m pip install wheel
RUN python3 -m pip install six
RUN python3 -m pip install absl-py
RUN python3 -m pip install wrapt
RUN python3 -m pip install termcolor

RUN python3 -m pip install grpcio
RUN python3 -m pip install numpy
RUN python3 -m pip install tensorflow
RUN python3 -m pip install keras

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"


COPY . "$GOPATH/src/http-to-keras-api"
WORKDIR "$GOPATH/src/http-to-keras-api"
RUN go get .


RUN go build -o main .

ENTRYPOINT ["./main"]


