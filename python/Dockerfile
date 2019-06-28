FROM ubuntu:18.04
RUN apt-get update
RUN apt-get install -y  --no-install-recommends net-tools supervisor python3 python3-pip vim python3-setuptools && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    rm -rf /tmp/*

RUN pip3 install ptvsd


COPY supervisord.conf /etc/supervisor/supervisord.conf
COPY ./opt/app /opt/app

EXPOSE 3000

CMD ["/usr/bin/supervisord","-c","/etc/supervisor/supervisord.conf"]