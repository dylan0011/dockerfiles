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


RUN chmod +x bin/mqnamesrv

CMD cd ${ROCKETMQ_HOME}/bin && export JAVA_OPT=" -Duser.home=/opt" && sh mqnamesrv


EXPOSE 9876
VOLUME /opt/logs \
		/opt/store
