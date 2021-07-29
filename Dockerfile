FROM golang:latest as builder

RUN apt-get update
RUN apt-get install -y gcc
RUN apt-get install -y xorg-dev
RUN apt-get install -y libgl1-mesa-dev

ENV CGO_CFLAGS_ALLOW=".*"
ENV CGO_CXXFLAGS_ALLOW=".*"
ENV CGO_LDFLAGS_ALLOW=".*"


ENV GOPATH /go
RUN mkdir -p "$GOPATH/src/app/water_jug_riddle" "$GOPATH/bin"

ADD . ${GOPATH}/src/app/water_jug_riddle

WORKDIR ${GOPATH}/src/app/water_jug_riddle

RUN go get -u fyne.io/fyne

COPY go.mod go.sum ./
RUN go mod vendor
RUN go build -o try_me

FROM scratch as export
COPY --from=builder /go/src/app/water_jug_riddle .