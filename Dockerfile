FROM ubuntu:19.04
WORKDIR /myapp

COPY db /myapp/db
COPY kadvisor /myapp
COPY .env /myapp
COPY client/build /myapp/client/build

CMD ["./kadvisor"]