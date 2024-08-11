REQ_GEN_DEPS := curl protoc protoc-gen-go protoc-gen-go-vtproto

GEN_REPO := https://raw.githubusercontent.com/prometheus/prometheus/main/prompb
GEN_PROTO := remote.proto types.proto
GEN_MODEL_DIR := promrw/model

.ONESHELL:
.PHONY: generate

generate: ## Genertes Golang code based on Proto files from Prometheus repo
	$(foreach bin,$(REQ_GEN_DEPS),$(if $(shell which $(bin)),,$(error "Please install '$(bin)'")))

	TMPDIR=$$(mktemp --directory)
	trap "rm -r $${TMPDIR}" EXIT

	$(foreach proto,$(GEN_PROTO),curl -sSfL "$(GEN_REPO)/$(proto)" -o "$${TMPDIR}/$(proto)";)

	sed -i -E -e '/import "gogoproto/d' \
	-e 's/ \[\(gogoproto\.nullable\) = false\]//g' \
	-e 's/package = "prompb"/package = "\.\/;model"/g' \
	$${TMPDIR}/*.proto

	protoc --proto_path="$${TMPDIR}" \
	--go_out="$(GEN_MODEL_DIR)" \
	--go-vtproto_out="$(GEN_MODEL_DIR)" --go-vtproto_opt="features=size+marshal" \
	$(GEN_PROTO)
