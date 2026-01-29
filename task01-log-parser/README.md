Task 1: The Log Analyzer CLI
Scenario: You are debugging a production server. You have a server.log file mixed with INFO, WARNING, and ERROR messages. You need a fast tool to extract only the errors and save them to a separate file for the team to review.

Your Mission:

Write a Go program that accepts a filename as a command-line argument, reads the file line-by-line (to be memory efficient), counts the total errors, and writes all error lines to errors.txt.

Technical Concepts to Learn:

CLI Arguments: Reading os.Args.
File I/O: Using os.Open, os.Create.
Buffered I/O: Using bufio.Scanner (Crucial for large files; C++ devs love this because it’s like std::getline).
String Manipulation: Using strings.Contains.
Error Handling: The explicit if err != nil pattern.
Task 1 Requirements:

Create a dummy log file named server.log (manually or via code) with some random lines containing “INFO”, “WARN”, and “ERROR”.
The program must run like this: go run main.go server.log.
If no file is provided, print a usage message and exit.
Read the file using bufio (do not read the whole file into memory at once).
Filter lines containing the string “ERROR” (case-insensitive).
Write those lines to a new file called errors.txt.
Print a summary to the console: “Processing complete. Found X errors.”
Output Format:

Paste your main.go code below. I will review your handling of file pointers (defer/close) and error checking.