FROM golang:latest 
RUN mkdir /app 
WORKDIR /app
ADD . /app/
# RUN git clone https://github.com/jacewalker/jace.au
WORKDIR /app/
RUN go install
RUN go build -o main .
EXPOSE 8080
CMD ["/app/main"]

# docker build -t jaceau:latest . --no-cache
# docker run --name jace.au -itd --network=proxy -p :8080 jaceau:latest