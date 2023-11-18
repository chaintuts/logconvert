# This file contains a make script for the LogConvert application
#
# Author: Josh McIntyre
#

# This block defines makefile variables
SRC_FILES=src/*.go
BUILD_DIR=bin/logconvert

# This rule builds the application
# Here we are simply wrapping Go's build tool
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR) $(SRC_FILES)

# This rule cleans the build directory
clean: $(BUILD_DIR)
	rm -r $(BUILD_DIR)/*
	rmdir $(BUILD_DIR)
