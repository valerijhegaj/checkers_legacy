#on intel macbook
fyne package

MAC_BIN_PATH="chekers.app/Contents/MacOS"

mkdir $MAC_BIN_PATH/saves
tar -czf download/macos-amd64/checkers.tar.gz $MAC_BIN_PATH/chekers $MAC_BIN_PATH/saves
