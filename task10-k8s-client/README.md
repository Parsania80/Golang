Task 10: Kubernetes Client (The client-go Operator)
Folder Name: task10-k8s-client

Scenario:

You want to automate Kubernetes. Maybe you want to automatically restart pods that have been crashing too often, or perhaps you want to perform a custom backup whenever a new “Deployment” is created. To do this, you don’t shell out to kubectl. You write a program that speaks the native Kubernetes protocol.

Goal:

Create a program that connects to your local Kubernetes cluster, lists all Pods in the “default” namespace, and watches for changes in real-time (e.g., when a new Pod is created or deleted).

Technical Concepts to Learn:

The client-go Library: k8s.io/client-go. The official Go client for Kubernetes.
Kubeconfig: How Go finds your cluster credentials (usually ~/.kube/config).
Typed Clients: Working with Clientset to access specific resources (Pods, Deployments).
The Watch API: Kubernetes is event-driven. You don’t just “list” resources; you “watch” streams of events.
Prerequisites:

You need a running Kubernetes cluster. Use Minikube, Kind (Kubernetes in Docker), or Docker Desktop’s built-in Kubernetes.
Ensure kubectl get pods works in your terminal.
Requirements:

Dependencies:
go get k8s.io/client-go@latest
go get k8s.io/apimachinery@latest
Note: client-go versions can be tricky. It is best to match the version of your Kubernetes cluster, but usually, latest works for simple labs.
Authentication (In-Cluster vs Out-of-Cluster):
Since you are running this from your laptop, use Out-of-Cluster configuration.
Use clientcmd.BuildConfigFromFlags to load your kubeconfig file (default location: $HOME/.kube/config).
Part A: The Listing (Static Snapshot):
Create a kubernetes.Clientset.
Use clientset.CoreV1().Pods("default").List(...).
Iterate through the list and print: Found Pod: [Name] (Status: [Phase]).
Part B: The Watcher (Real-time Events):
This is the core of a Controller/Operator.
Use clientset.CoreV1().Pods("default").Watch(...).
This returns a watch.Interface.
Iterate over the ResultChan() in a for/range loop.
Logic:
Switch on event.Type (ADDED, DELETED, MODIFIED).
Cast event.Object to a *v1.Pod.
Print: EVENT: [Type] -> Pod [Name].
Success Criteria (for when you code it):

Run your program. It should list existing pods (if any) and then hang/wait (because it is watching).
Open a second terminal.
Run kubectl run nginx-test --image=nginx.
Your program should immediately print: EVENT: ADDED -> Pod nginx-test.
Run kubectl delete pod nginx-test.
Your program should print: EVENT: MODIFIED (as it terminates) and finally EVENT: DELETED.
Why this matters for DevOps:

This is how Custom Resource Definitions (CRDs) and Operators work. When you install a database operator (like Postgres Operator) in K8s, it is running this exact Watch loop, waiting for you to create a database resource so it can spin up the actual pods.

Conclusion of the Roadmap
You now have the specifications for all 10 tasks.

My recommendation:

Start with Task 1 (Log Parser). Do not look at Task 2 until Task 1 is working and you are happy with the code.

If you get stuck or want a Code Review:

Paste your solution for Task 1 here, and I will critique it specifically from a “Go idioms” and “Performance” perspective (since you are coming from C++, I will look for pointers vs values, allocation efficiency, and proper error handling).