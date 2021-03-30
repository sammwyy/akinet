# Initial variables
$BIN_NAME="akinet"

# Compile for windows
$GOOS="linux"
echo "Compiling binary for linux..."
go build -o ./out/akinet-linux ./src/main.go

# Compile for windows
$GOOS="windows"
echo "Compiling binary for windows..."
go build -o ./out/akinet-win32.exe ./src/main.go

# Show a message when the process end
echo "Task completed."