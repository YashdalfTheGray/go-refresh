.PHONY: all run clean
all: bin/go-refresh

# go list is the canonical utility to find go files
GOFILES := $(shell go list -f '{{ join .GoFiles "\n" }}' ./...)

# check if imgcat exists on the system
IMGCAT := $(shell command -v imgcat 2> /dev/null)

bin/go-refresh: .bin-stamp ${GOFILES}
	go build -o bin/go-refresh
	chmod +x bin/go-refresh

# directories do werid things in make, so we can use a stamp
.bin-stamp:
	mkdir -p bin
	touch .bin-stamp

# Use 'go run' so we don't have to worry about recompiling
run:
	go run .

clean:
	rm -rf bin
	rm .bin-stamp

# go program generates an image, this is an easy way to convert
# and display the .png
image:
	base64 --decode --input bin/temp.base64 --output bin/temp.png
ifndef IMGCAT
	$(echo "no imgcat found on system, skipping displaying image")
endif
	imgcat bin/temp.png
