FROM alpine AS prebuilder

FROM minio/minio
ENV TZ="Asia/Shanghai"
RUN echo $TZ | tee /etc/timezone

FROM apsl/thumbor




ENV TZ="Asia/Shanghai"
RUN echo $TZ | tee /etc/timezone


ENV MINIO_ACCESS_KEY="access"
ENV MINIO_SECRET_KEY="secret"
ENV AWS_ACCESS_KEY_ID=$MINIO_ACCESS_KEY
ENV AWS_SECRET_ACCESS_KEY=$MINIO_SECRET_KEY



