FROM alpine
CMD [ "cp certs/domain.crt /usr/local/share/ca-certificates/myregistrydomain.com.crt", "update-ca-certificates" ]
COPY dist/servercheck /bin/
EXPOSE 8005
ENTRYPOINT [ "/bin/servercheck" ]
