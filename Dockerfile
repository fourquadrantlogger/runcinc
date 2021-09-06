FROM golang
ENV GOPROXY=https://goproxy.cn
WORKDIR /home/runcic
COPY . /home/runcic
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
FROM centos
COPY --from=0 /home/runcic/runcic* /usr/bin/
RUN yum install -y podman &&yum clean all &&rm -rf /tmp/
RUN mkdir /image /cic /cic/up /cic/work &&chmod +x /usr/bin/runcic.sh