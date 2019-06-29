FROM ubuntu:18.04

ENV GOROOT /goroot
ENV GOPATH /gopath
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin
ENV GOVERSION 1.12.6
ENV APPDIR /app

RUN apt-get update
RUN apt-get install -y  --no-install-recommends net-tools supervisor curl ca-certificates git && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    rm -rf /tmp/*

# Create go directories, install golang
RUN mkdir $GOROOT && mkdir $GOPATH && mkdir $APPDIR && curl -fsSL "https://golang.org/dl/go$GOVERSION.linux-amd64.tar.gz" | tar xvzf - -C /goroot --strip-components=1

COPY ./app/main.go /app/main.go
COPY supervisord.conf /etc/supervisor/supervisord.conf

RUN cd $APPDIR && go build -gcflags "all=-N -l" && go get -u github.com/go-delve/delve/cmd/dlv

EXPOSE 3000

CMD ["/usr/bin/supervisord","-c","/etc/supervisor/supervisord.conf"]