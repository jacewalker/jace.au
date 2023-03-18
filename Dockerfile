FROM golang:latest 
RUN mkdir /app 
WORKDIR /app
RUN git clone https://github.com/jacewalker/jace.au
WORKDIR /app/jace.au
RUN go install
RUN go build -o main .
EXPOSE 8080
CMD ["/app/jace.au/main"]

# docker build -t jaceau:latest . 
# docker run -it -p 80:8080 jaceau:latest