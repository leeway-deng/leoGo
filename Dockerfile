# leoGo
# VERSION 0.0.1
# docker login daocloud.io
# docker pull daocloud.io/library/golang:1.8;docker images
# docker rmi -f leo-leogo;docker build -t leo-leogo .

#Base image
FROM daocloud.io/library/golang:1.8

MAINTAINER dengliwei dengliwei@le.com

# Expose the application on port 8081
EXPOSE 8081

RUN mkdir -p /app/go/leoGo/config
RUN mkdir -p /data/leoGo

COPY dist/leoGo-amd64-1.0.0/leoGo /app/go/leoGo
COPY config /app/go/leoGo/config

WORKDIR /app/go
VOLUME /data/leoGo

#ENTRYPOINT /app/go/leoGo/leoGo
#RUN ["chmod", "+x", "/app/go/leoGo/leoGo"]
#CMD [ "/bin/bash", "-c", "ls -ll /app/go/leoGo" ]
#CMD [ "/bin/bash", "-c", "go version" ]
CMD [ "/bin/bash", "-c", "cd /app/go/leoGo;./leoGo" ]
