FROM golang:1.19-alpine

WORKDIR /app

COPY *.go .
COPY go.mod .
COPY go.sum .
COPY issue.html .
COPY db db

RUN go mod download
RUN go build -o /issue-tracker

EXPOSE 8080

CMD [ "/issue-tracker" ]
