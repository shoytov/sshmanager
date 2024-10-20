
#!/bin/bash

PROGRAM_NAME="sshm"

echo "Компиляция для Linux..."
GOOS=linux GOARCH=amd64 go build -o ${PROGRAM_NAME}_linux_amd64 .

echo "Компиляция для macOS Intel..."
GOOS=darwin GOARCH=amd64 go build -o ${PROGRAM_NAME}_darwin_amd64 .

echo "Компиляция для macOS M1..."
GOOS=darwin GOARCH=arm64 go build -o ${PROGRAM_NAME}_darwin_arm64 .

echo "Компиляция завершена."
