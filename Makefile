PROJECTNAME=hello

BUILD_DIR=build

PATH_DARWIN_ARM64=$(BUILD_DIR)/macos-arm64
PATH_DARWIN_AMD64=$(BUILD_DIR)/macos-amd64
PATH_LINUX_AMD64=$(BUILD_DIR)/linux-amd64
PATH_WINDOWS_AMD64=$(BUILD_DIR)/windows-amd64

go-darwin-arm64:
    mkdir -p $(PATH_DARWIN_ARM64)
    env GOOS=darwin GOARCH=arm64 go build -o $(PATH_DARWIN_ARM64)/$(PROJECTNAME)

go-darwin-amd64:
    mkdir -p $(PATH_DARWIN_AMD64)
    env GOOS=darwin GOARCH=amd64 go build -o $(PATH_DARWIN_AMD64)/$(PROJECTNAME)

go-linux-amd64:
    mkdir -p $(PATH_LINUX_AMD64)
    fyne-cross linux -arch=amd64 -name $(PROJECTNAME)
    mv fyne-cross/bin/linux-amd64/$(PROJECTNAME) $(PATH_LINUX_AMD64)/$(PROJECTNAME)
    rm -rf fyne-cross

go-windows-amd64:
    mkdir -p $(PATH_WINDOWS_AMD64)
    fyne-cross windows -arch=amd64 -name $(PROJECTNAME)
    mv fyne-cross/bin/windows-amd64/$(PROJECTNAME).exe $(PATH_WINDOWS_AMD64)/$(PROJECTNAME).exe
    rm -rf fyne-cross

pack-darwin-arm64:
    mkdir -p $(PATH_DARWIN_ARM64)/tmp
    cd $(PATH_DARWIN_ARM64)/$(PROJECTNAME) $(PATH_DARWIN_ARM64)/tmp/$(PROJECTNAME)
    mv $(PATH_DARWIN_ARM64)/tmp $(PATH_DARWIN_ARM64)/$(PROJECTNAME)
    tar -czf download/macos-arm64/checkers.tar.gz $(PATH_DARWIN_ARM64)/$(PROJECTNAME)

pack-darwin-amd64:
    mkdir -p $(PATH_DARWIN_AMD64)
    env GOOS=darwin GOARCH=amd64 go build -o $(PATH_DARWIN_AMD64)/$(PROJECTNAME)

pack-linux-amd64:
    mkdir -p $(PATH_LINUX_AMD64)
    fyne-cross linux -arch=amd64 -name $(PROJECTNAME)
    mv fyne-cross/bin/linux-amd64/$(PROJECTNAME) $(PATH_LINUX_AMD64)/$(PROJECTNAME)
    rm -rf fyne-cross

pack-windows-amd64:
    mkdir -p $(PATH_WINDOWS_AMD64)
    fyne-cross windows -arch=amd64 -name $(PROJECTNAME)
    mv fyne-cross/bin/windows-amd64/$(PROJECTNAME) $(PATH_WINDOWS_AMD64)/$(PROJECTNAME)
    rm -rf fyne-cross

clean:
    rm -rf $(PATH_DARWIN_ARM64) $(PATH_DARWIN_AMD64) $(PATH_LINUX_AMD64) $(PATH_WINDOWS_AMD64)

help:
    echo "go-darwin-arm64 - compile for macos arm"
    echo "go-darwin-amd64 - compile for macos intel"
    echo "go-linux-amd64 - compile for linux x86-64"
    echo "go-darwin-arm64 - compile for macos arm"
    echo "compile-on-mac-m1 - compile on macos arm for macos arm, linux x86-64, macos arm"
    echo "pack-darwin-arm64 - compile for macos arm"
    echo "pack-darwin-amd64 - compile for macos intel"
    echo "pack-linux-amd64 - compile for linux x86-64"
    echo "pack-darwin-arm64 - compile for macos arm"

all: help

default: help