FROM daocloud.io/golang
maintainer kaesa.li@daocloud.io
WORKDIR /gopath/app
ENV GOPATH /gopath/app
RUN apt-get update
ADD . /gopath/app/
RUN go install redis_cluster_test
RUN rm -fr /gopath/app/src
ENV TZ Asia/Shanghai
EXPOSE 8787
CMD ["/gopath/app/bin/redis_cluster_test"]