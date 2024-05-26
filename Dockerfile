FROM golang:1.22
WORKDIR /app
COPY . .
RUN go mod download
RUN make build
EXPOSE 8080
CMD ["make", "run-bin"]
