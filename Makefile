ARTIFACT = filehunt
SRC = main.go
OBJ = build/$(ARTIFACT)

buildL:
	@echo packing release......
	@goreleaser release --clean --skip=publish

