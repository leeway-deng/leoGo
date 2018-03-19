# leoGo
# VERSION 0.0.1
# docker login daocloud.io
# docker pull daocloud.io/library/golang:1.8;docker images
# docker build -t leo-leogo .

#Base image
FROM daocloud.io/library/golang:1.8

MAINTAINER dengliwei dengliwei@le.com

# Expose the application on port 8081
EXPOSE 8081

RUN mkdir -p /app/go
ADD dist/leoGo-amd64-1.0.0 /app/go/leoGo
ADD config /app/go/leoGo/config

WORKDIR /app
ENTRYPOINT /app/go/leoGo

VOLUME /data/leoGo

CMD [ "/bin/bash","-c", "cd /app/go/leoGo;nohup ./leoGo &" ]
