FROM golang:1.17.8-alpine3.15

RUN mkdir APIWithGin

WORKDIR /APIWithGin

COPY . .

RUN export GO111MODULE=on
RUN cd /APIWithGin
RUN go get github.com/gin-gonic/gin
RUN go build -o main.exe

EXPOSE 8080

CMD [ "/APIWithGin/main.exe" ]