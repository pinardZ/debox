FROM alpine:latest

ADD debox /

RUN chmod 777 /debox

# 通过先设置一个环境变量，然后在容器运行时传入环境变量具体的值，达到外部指定参数运行的效果

ENV PARAMS=""

ENTRYPOINT ["./debox"]