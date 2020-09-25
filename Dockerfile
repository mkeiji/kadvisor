FROM ubuntu
WORKDIR /myapp

COPY kadvisor /myapp
COPY client/dist /myapp/client/dist

CMD ["./kadvisor"]

###########################################################
#FROM golang
#WORKDIR /myapp

#COPY . .

##RUN apt-get update -yqq \
    ##&& apt-get install -yqq --no-install-recommends \
    ##nodejs npm \
    ##&& apt-get -q clean \
    ##&& rm -rf /var/lib/apt/lists \
    ##&& npm install -g nx

#RUN cd client/ && npm install && nx build \
    #&& cd .. && go build

#CMD ["./kadvisor"]
