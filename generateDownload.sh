#on m1 macbook
fyne package
fyne-cross windows -arch=amd64
fyne-cross linux -arch=amd64


WIN_BIN_PATH="fyne-cross/bin/windows-amd64"
LINUX_BIN_PATH="fyne-cross/bin/linux-amd64"
MAC_BIN_PATH="chekers.app/Contents/MacOS"

mkdir $WIN_BIN_PATH/saves
zip download/windows-amd64/checkers.zip $WIN_BIN_PATH/checkers.exe $WIN_BIN_PATH/saves

mkdir $LINUX_BIN_PATH/saves
tar -czf download/linux-amd64/checkers.tar.gz $LINUX_BIN_PATH/checkers $LINUX_BIN_PATH/saves

mkdir $MAC_BIN_PATH/saves
tar -czf download/macos-arm64/checkers.tar.gz $MAC_BIN_PATH/chekers $MAC_BIN_PATH/saves
