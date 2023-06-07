## General
____________

### Author
* Josh McIntyre

### Website
* jmcintyre.net

### Overview
* LogConvert is a simple ODS to CSV converter script

## Development
________________

### Git Workflow
* master for releases (merge development)
* development for bugfixes and new features

### Building
* make build
Build the application - wraps `go build`
* make clean
Clean the build directory

### Features
* Find all ODS files in a directory
* Convert using LibreOffice headless
* Convert to a specified output directory

### Requirements
* Requires Go language build tools
* Requires LibreOffice to be installed

### Platforms
* Windows
* MacOSX
* Linux

## Usage
____________

### Command Line Usage
* Run `./logconvert <source_dir> <output_dir>`