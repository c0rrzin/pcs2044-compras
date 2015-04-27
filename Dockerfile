# Pull base image.
FROM dockerfile/ubuntu

# Install Go
RUN \
  mkdir -p /goroot && \
  curl https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz | tar xvzf - -C /goroot --strip-components=1

# Set environment variables.
ENV GOROOT /goroot
ENV GOPATH /gopath
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH

# Define working directory.
WORKDIR /gopath/src/c0rrzin/github.com/compras

#add dependencies
RUN \
  go get github.com/c0rrzin/router && \
  go get github.com/jinzhu/gorm && \
  go get github.com/mattn/go-sqlite3

#copy source files
RUN mkdir -p /gopath/src/c0rrzin/github.com/compras
COPY . /gopath/src/c0rrzin/github.com/compras/

#compile and run
RUN go build

#add permission to execute
RUN chmod +x compras

#run
CMD ["./compras"]

#expose port
EXPOSE 8080
