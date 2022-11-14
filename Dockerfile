# Base image
# FROM node:lts-bullseye
FROM node

MAINTAINER Bluesentinelsec <bluesentinel@protonmail.com>

RUN apt-get update -y

# install build dependencies
RUN apt-get install -y -q \
    build-essential \
    tar \
    zip \
    unzip \
    wget \
    git \
    make \
    awscli \
    python3-pip \
    python3-venv

# install Go
ENV PATH "/usr/local/go/bin:$PATH"
ENV GOPATH "/opt/go/"
ENV PATH "$PATH:$GOPATH/bin"
RUN arch=$(arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/) && wget https://go.dev/dl/go1.19.linux-${arch}.tar.gz 
RUN arch=$(arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/) && tar -xvf go1.19.linux-${arch}.tar.gz
RUN mv go /usr/local/

# install AWS CDK
RUN npm install -g aws-cdk
RUN cdk --version
RUN python3 -m pip install aws-cdk-lib

# container entry point
CMD ["/bin/bash"]
