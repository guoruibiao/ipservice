FROM ubuntu
WORKDIR /app
ADD . /app
EXPOSE 8080
CMD ["/app/ipservice"]
