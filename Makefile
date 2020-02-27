BAZEL           = bazel

clean:
	$(BAZEL) clean --expunge

.PHONY: build
build: gazelle
	$(BAZEL) build //...

.PHONY: build-tool
build-tool: gazelle
	$(BAZEL) build //tool/punk
	$(BAZEL) build //tool/punk-gen-bts
	$(BAZEL) build //tool/punk-gen-project
	$(BAZEL) build //tool/punk-gen-redis
	$(BAZEL) build //tool/punk-protoc
#
# .PHONY: build-all-platform-tool
# build-all-platform-tool: gazelle
# 	$(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:linux_386 //tool/punk
# 	$(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //tool/punk
# 	$(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:darwin_amd64 //tool/punk
#
# 	$(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:linux_386 //tool/punk-gen-bts
# 	$(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //tool/punk-gen-bts
#     $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:darwin_amd64 //tool/punk-gen-bts
#
#
# 	$(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:linux_386 //tool/punk-gen-project
# 	$(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //tool/punk-gen-project
#     $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:darwin_amd64 //tool/punk-gen-project
#
#
# 	$(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:linux_386 //tool/punk-gen-redis
# 	$(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //tool/punk-gen-redis
#     $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:darwin_amd64 //tool/punk-gen-redis

dep-ensure:
	go mod download
	go mod vendor

#gazelle-repos:
#	$(BAZEL) run //:gazelle -- update-repos -from_file=go.mod

gazelle:
	$(BAZEL) run //:gazelle

go-setup: dep-ensure gazelle
