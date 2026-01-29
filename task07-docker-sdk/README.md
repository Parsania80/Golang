Task 7: Docker Dashboard (Docker SDK)
Folder Name: task07-docker-sdk

Scenario:

You are building a custom monitoring tool or a cleanup script. Instead of shelling out to the OS to run exec("docker ps") (which is fragile and hard to parse), you will speak directly to the Docker Engine API using Go.

Goal:

Create a CLI tool that lists running Docker containers and provides a command to stop a specific container by ID.

Technical Concepts to Learn:

The Docker SDK: github.com/docker/docker/client.
Contexts (context package): Go’s standard way of handling timeouts and cancellations (crucial when making API calls to the Docker daemon).
Struct Formatting: Using text/tabwriter to print pretty, aligned columns (like the real docker ps command).
API Interaction: Understanding ContainerList, ContainerStop.
Prerequisites:

You must have Docker Desktop or Docker Engine installed and running on your machine.
Requirements:

Setup & Dependencies:
Initialize the module.
Get the SDK: go get github.com/docker/docker/client.
Note: You might need to adjust versioning in go.mod if there are conflicts, but usually go get handles it.
Client Initialization:
Create a Docker client instance:
go
        cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
FromEnv allows it to find your local Docker socket automatically.
Command 1: List Containers (ps)
Fetch the list of containers using cli.ContainerList.
Display the following columns cleanly:
Container ID (shortened to 12 chars)
Image Name
Status (e.g., “Up 2 hours”)
State (e.g., “running”)
Constraint: You must use text/tabwriter to align the columns perfectly, regardless of how long the image name is.
Command 2: Stop Container (stop)
Accept a Container ID as an argument.
Call cli.ContainerStop.
You must pass a context to this call.
Print “Stopping container [ID]…” and then “Success” or the error message.
Logging/Logs (Optional Challenge):
If you breeze through this, try to fetch the logs of a specific container (cli.ContainerLogs) and print the last 5 lines.
Success Criteria (for when you code it):

Start an Nginx container manually: docker run -d --name my-nginx -p 8080:80 nginx.
Run your tool: go run main.go ps. You should see my-nginx listed.
Run your tool: go run main.go stop <container_id>.
Run docker ps manually. The container should be gone (or stopped).
Why this matters for DevOps:

This is the foundation of Kubernetes operators. Kubernetes Controllers are essentially just infinite loops that check the status of containers/pods and issue API calls to fix them.