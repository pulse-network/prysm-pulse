FROM ubuntu:18.04 as builder

RUN apt update

# install basic build dependencies
RUN apt install -y cmake git libssl-dev libgmp-dev libtinfo5 unzip curl python g++

# install protoc at specific required version
RUN PROTOC_ZIP=protoc-3.14.0-linux-x86_64.zip && \
    curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/$PROTOC_ZIP && \
    unzip -o $PROTOC_ZIP -d /usr/local bin/protoc && \
    unzip -o $PROTOC_ZIP -d /usr/local 'include/*' && \
    rm -f $PROTOC_ZIP

# install bazelisk
RUN curl -o bazelisk -L https://github.com/bazelbuild/bazelisk/releases/download/v1.14.0/bazelisk-linux-amd64 && \
    chmod +x bazelisk && \
    mv bazelisk /usr/local/bin && \
    ln -s /usr/local/bin/bazelisk /usr/local/bin/bazel && \
    bazelisk --version

# create app placeholder
RUN mkdir /app
WORKDIR /app

# prevent git warnings throughout build process
RUN git config --global --add safe.directory /app

# perform the build then print output directories
CMD bazel build //cmd/beacon-chain:image_bundle --config=release && \
    bazel build //cmd/validator:image_bundle --config=release && \
    bazel build cmd/beacon-chain:image_bundle.tar && \
    bazel build cmd/validator:image_bundle.tar && \
    ls -l | grep -o bazel-.* | sed 's/\/root\/.cache/.bazel-in-docker-cache/g' | awk '{print "ln -sf " $3 " " $1}' > .bazel-in-docker-outputs && \
    chmod 666 .bazel-in-docker-outputs
