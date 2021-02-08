FROM ubuntu:latest

COPY build/bin/ /usr/local/bin/
RUN chmod 555 /usr/local/bin/github-actions-exporter

EXPOSE 9870
ENTRYPOINT ["./usr/local/bin/github-actions-exporter"]
