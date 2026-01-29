Task 5: Professional CLI with Cobra (Flags & UX)
Folder Name: task05-cobra-cli

Scenario:

So far, you’ve parsed os.Args manually. In the real world (Docker, Kubernetes, Terraform), CLIs are complex. They have subcommands (docker run, docker build), flags (--verbose, -f file), and help menus. The industry standard library for this in Go is Cobra.

Goal:

Refactor or recreate a tool using the Cobra library. You will build a file utility CLI called mytool that mimics basic Linux commands but with a modern Go structure.

Technical Concepts to Learn:

Cobra Library: github.com/spf13/cobra (used by Kubernetes kubectl).
Subcommands: Structuring commands like app start, app stop.
Flags: Implementing persistent flags (global) vs local flags.
Viper (Optional): Often used with Cobra for reading config files, but focus on Cobra first.
Requirements:

Project Setup:
Initialize the module.
Install Cobra: go get -u github.com/spf13/cobra@latest.
Use the Cobra generator (optional but recommended) or structure it manually:
text
        main.go
        cmd/
          root.go
          count.go
          info.go
The Root Command:
The base command mytool should print a help menu explaining available commands.
Subcommand 1: info
Usage: mytool info
Action: Prints system details (OS, Architecture, current user). Re-use your logic from Task 3 (System Monitor) or just simple runtime package calls.
Subcommand 2: count
Usage: mytool count [string]
Action: Counts the number of characters/words in the provided string.
Flags: Add a specific flag --lines or -l.
If the user provides a file path instead of a string, and uses -l, it counts the lines in that file.
Global Flag:
Add a --verbose flag that works on ALL commands.
If --verbose is true, print extra debug logs (e.g., “DEBUG: Parsing file…”).
Success Criteria (for when you code it):

Running go run main.go shows a nice auto-generated help menu.
Running go run main.go count "hello world" works.
Running go run main.go count --lines test.txt works.
You understand how init() functions are used in Cobra to register flags.