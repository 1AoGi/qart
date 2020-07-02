FROM library/golang

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

# Setup for proxy
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.io,direct

ENV APP_DIR $GOPATH/src/qart
RUN mkdir -p $APP_DIR

# Set the entrypoint
ADD . $APP_DIR
WORKDIR $APP_DIR
ENTRYPOINT ($APP_DIR/qart)
CMD ["--prod"]

# Compile the binary and statically link
RUN cd $APP_DIR && CGO_ENABLED=0 go build -ldflags '-d -w -s'

EXPOSE 8080
