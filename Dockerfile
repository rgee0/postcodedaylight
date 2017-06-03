FROM golang:1.7.5-alpine
MAINTAINER docker@technologee.co.uk
ENTRYPOINT []

WORKDIR /go/src/github.com/rgee0/postcodedaylight
COPY . /go/src/github.com/rgee0/postcodedaylight

RUN go install

ADD https://github.com/alexellis/faas/releases/download/0.5.6-alpha/fwatchdog /usr/bin
RUN chmod +x /usr/bin/fwatchdog

ENV fprocess "/go/bin/postcodedaylight"

CMD [ "/usr/bin/fwatchdog"]