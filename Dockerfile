FROM --platform=linux/amd64 alpine:3.19

WORKDIR /workspace

RUN apk add go curl unzip bash sudo nodejs npm vim

ENV GOPATH=/root/go
ENV PATH=$PATH:$GOPATH/bin

# install terraform:
RUN curl -O https://releases.hashicorp.com/terraform/1.7.0/terraform_1.7.0_linux_amd64.zip && \
    unzip terraform_1.7.0_linux_amd64.zip && \
    mv terraform /usr/local/bin/ && \
    rm terraform_1.7.0_linux_amd64.zip