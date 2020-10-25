FROM golang:1.14

ENV APIENV=PROD
ENV PORT=8080
EXPOSE 8080

WORKDIR /app/robovsdino
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go test -v
RUN go build -o ./out/robovsdino .

CMD ["./out/robovsdino"]