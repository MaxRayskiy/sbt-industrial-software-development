FROM golang:1.19-alpine

WORKDIR /app

COPY *.go go.mod go.sum issue.html ./
COPY db ./db

RUN go mod download
RUN go build -o /issue-tracker

EXPOSE 8080

CMD [ "/issue-tracker" ]
