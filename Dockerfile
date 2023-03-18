FROM golang:latest 
RUN mkdir /app 
# ADD . /app/ 
WORKDIR /app
RUN git clone https://github.com/jacewalker/jace.au
WORKDIR /app/jace.au
RUN go install
RUN go build -o main .
EXPOSE 8080
CMD ["/app/jace.au/main"]