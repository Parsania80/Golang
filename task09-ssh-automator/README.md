Task 9: SSH Automator (Remote Execution)
Folder Name: task09-ssh-automator

Scenario:

You need to update 50 servers. You could open 50 terminal windows, or you could write a Go program to loop through them, connect via SSH, run the update command, and report back the results.

Goal:

Create a program that connects to an SSH server using a username and password (or private key), establishes a session, runs a command (like whoami or df -h), and prints the output to your local console.

Technical Concepts to Learn:

The SSH Package: golang.org/x/crypto/ssh (Note: This is an “extended” standard library, managed by the Go team but not in the main tree).
Client Configuration: Setting up ssh.ClientConfig with Auth methods and Host Key callbacks.
Sessions: Understanding that in SSH, you dial a Client (the connection), but you execute commands inside a Session.
Output Capture: redirecting remote stdout to a local buffer.
Prerequisites (The Target):

You need an SSH server to test against. DO NOT use your main OS if you are uncomfortable. The safest way is to spin up a throwaway Docker container that acts as a server:

bash
# Runs a tiny Alpine Linux with SSH listening on port 2222
# User: root, Password: root
docker run -d -p 2222:22 -e ROOT_PASSWORD=root --name my-ssh-target sickp/alpine-sshd:7.5
Requirements:

Dependencies:
go get golang.org/x/crypto/ssh
Configuration Struct:
Create a struct to hold connection details: IP, Port, User, Password, Command.
The Connection Logic (ssh.Dial):
Configure ssh.ClientConfig.
Auth: Use ssh.Password("root").
Host Key Callback: For this specific learning task, use ssh.InsecureIgnoreHostKey().
Note: In production, this is dangerous (Man-in-the-Middle attacks), but for a local lab, it simplifies the code significantly.
Dial the server: ssh.Dial("tcp", "localhost:2222", config).
The Session Logic:
Once connected, create a session: client.NewSession().
Important: A session can usually only run one command. If you want to run two commands, you often need two sessions (or chain them in the command string).
Execution:
Run a command that returns output, e.g., cat /etc/os-release.
Capture the output using session.Output() or session.CombinedOutput() (which grabs both stdout and stderr).
Cleanup:
Ensure you defer client.Close() and defer session.Close().
Success Criteria (for when you code it):

Start the Docker SSH container (command above).
Run your Go program.
Your program should output the contents of the Alpine Linux release file from inside the container.
If you change the command to a nonsense command (e.g., blarg), your program should print the error returned by the SSH session.
Why this matters for DevOps:

This is the fundamental logic behind configuration management. When you run a Terraform remote-exec provisioner or an Ansible playbook, this is exactly what is happening under the hood: Dial 
→
→
 Session 
→
→
 Exec 
→
→
 Report.