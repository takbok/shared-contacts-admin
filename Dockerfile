FROM ubuntu:14.04

ENV GOPATH=/go
ENV GOLANG=https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz
ENV SCA_SRC=$GOPATH/src/github.com/litriv/shared-contacts-admin/
ENV GAE_SDK_URL=https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_386-1.9.35.zip
ENV GAE_SDK=/usr/local/go_appengine
ENV PATH=/usr/local/go/bin:$GAE_SDK:$PATH

RUN mkdir -p "$SCA_SRC"
COPY *.go "$SCA_SRC"
COPY app.yaml "$SCA_SRC"

RUN apt-get update && apt-get install -y \
    curl \
    unzip \
    python \
    git \
 && curl "$GAE_SDK_URL" > /tmp/gaesdk.zip  \
 && unzip /tmp/gaesdk.zip -d /usr/local \
 && curl "$GOLANG" | tar -C /usr/local -xz \
 && go get golang.org/x/oauth2 \
 && go get google.golang.org/appengine \
 && go get google.golang.org/cloud/compute/metadata 

WORKDIR $SCA_SRC

CMD goapp serve
