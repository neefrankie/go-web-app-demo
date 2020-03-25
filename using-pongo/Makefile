build_dir := build
BINARY := pongo

BUILD := `date +%FT%T%z`

LDFLAGS := -ldflags "-w -s -X main.build=${BUILD}"

.PHONY: build linux deploy attack report lastcommit mkbuild clean
build :
	go build $(LDFLAGS) -o $(build_dir)/$(BINARY) -v .

run :
	./$(build_dir)/${BINARY}

linux :
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(build_dir)/linux/$(BINARY) -v .

clean :
	go clean -x
	rm build/* && rm rice-box.go
