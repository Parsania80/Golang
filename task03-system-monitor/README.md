Task 3: The System Resource Monitor (Daemon)
Folder Name: task03-system-monitor

Scenario:

In a cloud environment, you often need custom agents running on servers to report resource usage or trigger alerts when things go wrong (e.g., “Disk is 90% full”). While tools like Prometheus exist, writing a lightweight monitor is a fundamental skill for understanding how Go interacts with the underlying Operating System.

Goal:

Create a program that runs indefinitely (like a daemon), snapshots the CPU and Memory usage every few seconds, checks against a “threshold,” and logs an alert if the threshold is exceeded.

Technical Concepts to Learn:

OS Interaction: Reading system files (like /proc on Linux) or using a library to fetch OS stats.
Signals: Handling Ctrl+C (SIGINT) gracefully to shut down the server safely.
Time & Tickers: Using time.NewTicker to run a task repeatedly.
Structs & JSON: Formatting the output data cleanly.
External Libraries: Using a popular library like gopsutil (Standard practice in Go is knowing when not to reinvent the wheel).
Requirements:

Library Setup:
You will not write C-style system calls manually. Use the standard Go library for this: github.com/shirou/gopsutil (specifically cpu and mem subpackages).
Run go get github.com/shirou/gopsutil/v3 in your project folder.
The Monitor Loop:
The program must run an infinite loop.
Every 2 seconds, it should fetch:
CPU Usage Percentage (total).
RAM Usage (Total, Used, and Percentage).
Constraint: Do not use time.Sleep(). Use time.NewTicker() for precise interval scheduling.
Threshold Logic:
Define a constant CPU_THRESHOLD = 50.0 (or lower if your PC is powerful, so you can trigger it).
If the current CPU usage > Threshold, print a warning message: ALERT: High CPU Usage detected: 55%!
Graceful Shutdown (Signal Handling):
Typically, if you hit Ctrl+C, a program dies instantly.
Implement a os.Signal channel to listen for syscall.SIGINT or syscall.SIGTERM.
When the user hits Ctrl+C, the program should catch it, print “Monitoring stopping…”, perform any cleanup (if needed), and then exit.
Output:
Clear, formatted console output updating every tick.
Example:
text
        [10:00:02] CPU: 12% | RAM: 45% (8GB/16GB)
        [10:00:04] CPU: 51% | RAM: 45% (8GB/16GB)
        ALERT: High CPU Usage detected!
Success Criteria (for when you code it):

The program compiles and runs.
It accurately reports your system stats.
It doesn’t crash when you press Ctrl+C, but exits with a custom message.
You understand how select statements work with channels (used for the ticker and the signal listener).