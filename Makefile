DEFAULT: lint
BIN?=/usr/local/bin

BUF_VERSION?=0.41.0
BUF_BINARY_NAME?=buf

# should match the version installed by `brew install protobuf`
PROTOC_VERSION?=3.19.4

PROTOC_GEN_GO_VERSION?=1.26.0
PROTOC_GEN_GO_GRPC_VERSION?=1.1.0

# nodejs related
NVM_VERSION?=0.35.3
NVM_DIR?=${HOME}/.nvm
NVM_BINARY_NAME?=nvm
NODE_VERSION?=14

#if we are on OSX we need to convert the 'Darwin' string produced by uname to 'osx' for the protoc zip file download path to be correct
PROTOC_OS_COMMAND := case $$(uname -s) in Darwin*) printf osx;; *) printf $$(uname -s);; esac;
PROTOC_OS := $(shell ${PROTOC_OS_COMMAND})

remove_buf:
	rm -f ${BIN}/${BUF_BINARY_NAME}

get_buf: remove_buf
	curl -sSL \
		"https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/${BUF_BINARY_NAME}-$(shell uname -s)-$(shell uname -m)" \
		-o "${BIN}/${BUF_BINARY_NAME}" && \
	chmod +x "${BIN}/${BUF_BINARY_NAME}"

get_protoc:
	curl -sSL \
		"https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-${PROTOC_OS}-$(shell uname -m).zip" \
		-o /tmp/proto.zip
	unzip -o /tmp/proto.zip -d /usr/local bin/protoc
	unzip -o /tmp/proto.zip -d /usr/local 'include/*'
	rm -f /tmp/proto.zip

get_go_tools:
	@echo Installing protobuf golang and gRPC drivers
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v${PROTOC_GEN_GO_VERSION}
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v${PROTOC_GEN_GO_GRPC_VERSION}


get_nodejs_tools:
	# curl -sL https://deb.nodesource.com/setup_${NODE_VERSION}.x | bash -
	# apt install nodejs -y
	# node --version
	# npm i -g --production grpc_tools_node_protoc_ts
	# npm i -g --production grpc-tools --unsafe-perm
	# npm i --production grpc_tools_node_protoc_ts
	# npm i --production grpc-tools --unsafe-perm
	npm i

get_nodejs_tools_nvm:
# Install this if nvm is required
	# curl -o- "https://raw.githubusercontent.com/nvm-sh/nvm/v${NVM_VERSION}/install.sh" | bash
	# [ -s "${NVM_DIR}/nvm.sh" ] && \. "${NVM_DIR}/nvm.sh"
	# [ -s "${NVM_DIR}/bash_completion" ] && \. "${NVM_DIR}/bash_completion"
	# chmod +x ${NVM_DIR}/nvm.sh
	# make nvm_install
	# make nvm_use
	# npm i -g --production grpc_tools_node_protoc_ts
	# npm i -g --production grpc-tools

	npm i

nvm_install:
	. ${NVM_DIR}/nvm.sh && nvm install ${NODE_VERSION}

nvm_use:
	. ${NVM_DIR}/nvm.sh && nvm use ${NODE_VERSION}

generate: lint
	rm -rf build
	buf generate protos/ -o build

lint:
	# buf breaking --against 'https://github.com/nwosuudoka/ssprotos#branch=master' protos/
	buf lint protos/
