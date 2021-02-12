package docker

var dockerTemplate = `FROM golang:1.16-rc-alpine as builder
ENV GO111MODULE onENV GOPROXY https://proxy.golang.org/
# Installing nodejs
RUN apk add --update nodejs-current curl bash build-base

# Installing Yarn
RUN curl -o- -L https://yarnpkg.com/install.sh | bash
ENV PATH="$PATH:/root/.yarn/bin:/root/.config/yarn/global/node_modules"

# Installing ox
RUN go install github.com/wawandco/oxpecker/cmd/ox@latest
WORKDIR /{{.Name}}
ADD . .

# Building the application binary in bin/app 
RUN ox build --static -o bin/app

FROM alpine
# Binaries
COPY --from=builder /{{.Name}}/bin/* /bin/

# For migrations use 
# CMD pop migrate up; app 
CMD app`
