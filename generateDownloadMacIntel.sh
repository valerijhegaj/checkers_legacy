#on intel macbook

go build -o checkers

MAC_BIN_PATH="chekers.app/Contents/MacOS"

mkdir $MAC_BIN_PATH/saves
tar -czf download/macos-amd64/checkers.tar.gz checkers $MAC_BIN_PATH/saves
