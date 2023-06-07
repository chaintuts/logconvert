# This file contains a simple script for combining CSV files
#
# Author: Josh McIntyre
#
import argparse
import csv
import os

def main():

    # Define the argparser and execute desired commands
    parser = argparse.ArgumentParser()
    parser.add_argument("--source_dir", help="The source log directory", required=True)
    parser.add_argument("--output_name", help="The output log name", required=True)
    parser.add_argument("--header", help="The header to write", required=True)
    parser.add_argument("--skip_rows", type=int, help="The number of header rows to skip", required=True)

    args = parser.parse_args()

    # Generate a list of all logs in the directory
    all_entries = os.listdir(args.source_dir)
    files = [ os.path.join(args.source_dir, _) for _ in all_entries if os.path.isfile(os.path.join(args.source_dir, _)) ]

    # Combine each into a file
    # Use a specified number of header lines to strip
    with open(args.output_name, "a", newline="") as fo:
        writer = csv.writer(fo)
        writer.writerow([args.header])
        for file in files:
            with open(file, newline="") as fi:
                reader = csv.reader(fi)
                for i in range(0, args.skip_rows):
                    _ = next(reader)
                writer.writerows(reader)
                

if __name__ == "__main__":
    main()
