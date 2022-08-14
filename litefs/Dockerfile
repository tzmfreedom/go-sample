FROM ubuntu

ENV TZ=Asia/Tokyo
RUN apt update
RUN apt install -y curl git sqlite fuse3 libfuse-dev consul wget gpg lsb-release vim unzip
RUN curl -LO https://go.dev/dl/go1.18.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.18.4.linux-amd64.tar.gz

RUN curl -O https://releases.hashicorp.com/consul/1.12.3/consul_1.12.3_linux_amd64.zip
RUN unzip consul_1.12.3_linux_amd64.zip -d /usr/local/bin/
# RUN wget -O- https://apt.releases.hashicorp.com/gpg | gpg --dearmor | tee /usr/share/keyrings/hashicorp-archive-keyring.gpg
# RUN echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/hashicorp.list
# RUN apt update && apt install -y consul

ENV GOROOT=/usr/local/go
ENV GOPATH=$HOME/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

ENV ARCH="amd64"
ENV VERSION="0.1.0"

WORKDIR /tmp
RUN wget "https://github.com/superfly/litefs/releases/download/v${VERSION}/litefs-v${VERSION}-linux-${ARCH}.tar.gz" && \
    tar xvf "litefs-v${VERSION}-linux-${ARCH}.tar.gz" && \
    mv litefs /usr/local/bin

# RUN git clone https://github.com/superfly/litefs
#
# WORKDIR /root/litefs
# RUN go build -ldflags "-s -w -X 'main.Version=latest' -extldflags '-static'" -tags osusergo,netgo,sqlite_omit_load_extension -o /usr/local/bin/litefs ./cmd/litefs

RUN mkdir /tmp/x /tmp/y
RUN apt install -y make gcc
WORKDIR /app

CMD make consul
