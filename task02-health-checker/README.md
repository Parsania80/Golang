Task 2: The Concurrent Website Health Checker
Folder Name: task02-health-checker

Scenario:

As a DevOps engineer, you need to verify if a list of microservices (URLs) are up and running. Checking them one by one (sequential) is too slow when you have 50 services. You need to check them all at the same time (concurrently) to save time.

Goal:

Build a CLI tool that takes a list of websites and checks if they are online (returning a 200 OK status). This project introduces Golang’s “Killer Feature”: Goroutines and Channels.

Technical Concepts to Learn:

Goroutines: The go keyword (lightweight threads).
Channels: Passing data between threads safely (chan string).
WaitGroups: Waiting for all threads to finish (sync.WaitGroup).
HTTP Client: Using net/http to make GET requests.
Timeouts: Ensuring a slow server doesn’t freeze your program.
Requirements:

Input Data: Define a slice of strings in your code containing at least 5 URLs (mix of real and fake):
https://google.com
https://github.com
https://stackoverflow.com
https://invalid-url-example.com (This should fail)
http://non-existent-domain.go (This should fail)
The Worker Function:
Create a function checkUrl(url string, c chan string, wg *sync.WaitGroup).
It must use http.Get to ping the URL.
It must handle errors (e.g., DNS failure).
It must format a result string (e.g., “✅ [200] https://google.com” or “❌ [Error] https://invalid.com”).
It must send this result string into the channel.
It must signal the WaitGroup that it is done.
Concurrency Implementation:
In main, loop through the list of URLs.
Launch a Goroutine for each URL (checking 5 sites should take roughly the same time as checking 1).
Do not use a simple for loop that waits; use the go keyword.
Data Collection:
You must read from the channel to print the results.
Challenge: Ensure the program doesn’t exit before all checks are complete (Hint: Use sync.WaitGroup).
Success Criteria (for when you code it):

The program should finish almost instantly (parallel execution), not pause for each site.
Output order will likely be random (because it’s concurrent)—this is expected behavior.
The code must use defer wg.Done() to ensure the counter decrements even if the function crashes.
Extensions (Optional for later):

Add a flag to check the URLs repeatedly every 5 seconds (like a real monitoring daemon).