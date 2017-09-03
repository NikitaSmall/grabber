BIN_DIR=bin
BINARY=grabber

BINARY_OUTPUT=${BIN_DIR}/${BINARY}

all:
	go build -o ${BINARY_OUTPUT}
