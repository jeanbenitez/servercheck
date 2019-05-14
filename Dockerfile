FROM alpine

COPY dist/servercheck /bin/

EXPOSE 5001

ENTRYPOINT [ "/bin/servercheck" ]