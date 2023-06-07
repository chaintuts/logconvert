// This file contains code that converts log files to CSV format
//
// Author: Josh McIntyre
//
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)


// Convert an individual file to CSV to the output dir
func convert_file (log_path string, output_dir string) {

	office_bin := "C:/Program Files/LibreOffice/program/soffice.exe"
	
	fmt.Printf("Coverting file %s to output_dir %s\n", log_path, output_dir)
	_, err := exec.Command(office_bin, "--headless", "--convert-to", "csv",  log_path, "--outdir", output_dir).Output()
    if err != nil {
		fmt.Printf("ERROR file %s to dir %s\n", log_path, output_dir)
		fmt.Println(err)
    } else {
		fmt.Printf("SUCCESS file %s to output dir %s\n", log_path, output_dir)
	}
}

// Handle conversion of log files
func find_and_convert(source_dir string, output_dir string) {

	// Find all log files in the source directory
	files, err := ioutil.ReadDir(source_dir)
	if err != nil {
		fmt.Printf("Unable to read from dir %s\n", source_dir)
	}
	
	for _, file := range files {
		if !file.IsDir() {
			log_path := path.Join(source_dir, file.Name())
			convert_file(log_path, output_dir)
		}
	}
}

// The main entry point for the program
func main() {

	// Fetch filenames from command line arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage ./logconvert <source dir> <output dir>")
		return 
	}

	args := os.Args[1:]
	source_dir := args[0]
	output_dir := args[1]
	
	find_and_convert(source_dir, output_dir)

}