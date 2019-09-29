FROM library/golang

## Godep for vendoring
RUN  curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

ENV APP_DIR $GOPATH/src/github.com/beego/samples/todo
RUN mkdir -p $APP_DIR

# Set the entrypoint
ENTRYPOINT (cd $APP_DIR && ./todo)
ADD . $APP_DIR

# Compile the binary and statically link
RUN cd $APP_DIR && dep ensure && CGO_ENABLED=0 go build -ldflags '-d -w -s'

EXPOSE 8080

