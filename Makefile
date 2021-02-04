OSQUERY_PKG_PATH = ./pkg/osquery

all: gen

gen:
	mkdir -p ${OSQUERY_PKG_PATH}/gen
	thrift --gen go:package_prefix=github.com/mlukasik-dev/osquery-server/pkg/osquery/gen \
		-out ${OSQUERY_PKG_PATH}/gen \
		${OSQUERY_PKG_PATH}/osquery.thrift
	rm -rf ${OSQUERY_PKG_PATH}/gen/osquery/extension-remote \
		${OSQUERY_PKG_PATH}/gen/osquery/extension_manager-remote
	gofmt -w ${OSQUERY_PKG_PATH}/gen

.PHONY: gen