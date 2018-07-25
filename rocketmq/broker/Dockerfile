# Start from a Java image.
FROM java:8

ARG version

# Rocketmq version
ENV ROCKETMQ_VERSION 4.1.0-incubating

# Rocketmq home
ENV ROCKETMQ_HOME  /opt/rocketmq-${ROCKETMQ_VERSION}



WORKDIR  ${ROCKETMQ_HOME}

RUN mkdir -p \
		/opt/logs \
	    /opt/store

RUN curl https://dist.apache.org/repos/dist/release/rocketmq/${ROCKETMQ_VERSION}/rocketmq-all-${ROCKETMQ_VERSION}-bin-release.zip -o rocketmq.zip \
          && unzip rocketmq.zip \
          && mv rocketmq-all*/* . \
          && rmdir rocketmq-all*  \
          && rm rocketmq.zip


RUN chmod +x bin/mqbroker

CMD cd ${ROCKETMQ_HOME}/bin && export JAVA_OPT=" -Duser.home=/opt" && sh mqbroker -n namesrv:9876

EXPOSE 10909 10911
VOLUME /opt/logs \
		/opt/store
