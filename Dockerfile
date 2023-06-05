FROM golang:1.19

RUN apt-get update && apt-get install -y wget 

RUN wget https://dist.ipfs.io/go-ipfs/v0.10.0/go-ipfs_v0.10.0_linux-amd64.tar.gz

RUN tar xvfz go-ipfs_v0.10.0_linux-amd64.tar.gz 
RUN mv go-ipfs/ipfs /usr/local/bin/ipfs
RUN rm -rf go-ipfs_v0.10.0_linux-amd64.tar.g go-ipfs
RUN ipfs init
RUN ipfs --version

WORKDIR /daemonpackage

COPY . .

RUN go build -o main .

CMD ["./main"]

ENV IPFS_PATH C:/User/gayat

