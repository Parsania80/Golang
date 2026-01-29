Task 6: Simple TCP Proxy (Port Forwarder)
Folder Name: task06-tcp-proxy

Scenario:

In Cloud Engineering, you often deal with “Sidecars,” Load Balancers (HAProxy, Nginx), or kubectl port-forward. These tools accept traffic on one port and blindly forward it to another service, often on a different machine.

Goal:

Build a TCP Proxy. Your program will listen on a local port (e.g., local:3000), accept a connection, dial a specific destination (e.g., google.com:80), and pipe the data back and forth between the two connections.

Technical Concepts to Learn:

The net Package: Working with raw TCP sockets (net.Listen, net.Conn).
io.Copy: The most efficient way to stream data from one reader to a writer in Go.
Bi-directional Concurrency: Handling traffic flowing in two directions simultaneously (Client 
→
→
 Remote AND Remote 
→
→
 Client) using Goroutines.
Resource Cleanup: Ensuring sockets are closed properly to prevent file descriptor leaks.
Requirements:

Arguments:
The program should accept two flags (or arguments):
--local: The port to listen on (e.g., :8080).
--remote: The target address to forward to (e.g., google.com:80 or 127.0.0.1:5432).
The Accept Loop:
Start a net.Listener.
In an infinite loop, Accept() new incoming connections.
Crucial: For every accepted connection, spawn a new Goroutine to handle it (go handleConnection(conn)). If you don’t do this, your proxy can only handle one user at a time!
The Proxy Logic (handleConnection):
Step A: Connect to the remote server (net.Dial("tcp", remoteAddr)). If this fails, close the local connection and log an error.
Step B: You now have two connections: localConn and remoteConn.
Step C: Setup the pipes. You need to stream data in both directions at the same time.
Goroutine 1: Copy data from localConn 
→
→
 remoteConn.
Goroutine 2: Copy data from remoteConn 
→
→
 localConn.
Hint: io.Copy(writer, reader) blocks until the connection closes.
Logging:
Log when a connection opens: Proxying connection from 127.0.0.1:56732 to google.com:80
Log when a connection closes.
Success Criteria (for when you code it):

Test 1 (Web): Run the proxy: ./proxy --local :9090 --remote google.com:80. Open your browser to http://localhost:9090. You should see Google (or a 404 from Google because the Host header is wrong, but you will see traffic).
Test 2 (SSH): If you have an SSH server, proxy to port 22. ssh -p 9090 localhost should connect you to the real SSH server.
Behavior: The proxy does not modify data; it just moves bytes. It essentially acts as a “Man in the Middle”.
Why this matters for DevOps:

This teaches you how Service Meshes (like Istio or Linkerd) and Sidecars actually work under the hood. They are just fancy TCP proxies injecting logic into io.Copy.