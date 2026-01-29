Task 4: Simple REST API (Configuration Service)
Folder Name: task04-rest-api

Scenario:

Modern infrastructure runs on APIs. Before you can work with Kubernetes (which is entirely API-driven), you must understand how to build a REST API. You will build a microservice that manages a “Configuration Store”—think of it as a tiny database where other apps can retrieve their settings.

Goal:

Build a web server that handles HTTP requests (GET, POST, DELETE) to manage a list of configurations using JSON. This moves you from “CLI tools” to “Network Services”.

Technical Concepts to Learn:

The net/http Package: Go’s standard library for building web servers (very powerful).
JSON Marshalling: Converting Go Structs to JSON and back (encoding/json).
HTTP Verbs: Handling GET vs POST logic.
Routing: Simple path matching (/config, /config/{id}).
State Management: Managing data in memory (Maps/Slices) safely.
Requirements:

Data Structure:
Define a struct Config with fields:
go
        type Config struct {
            Key   string `json:"key"`
            Value string `json:"value"`
            Service string `json:"service"`
        }
Use a global Slice or Map to store these configs in memory (no real database needed yet).
Endpoints:
GET /status: Returns 200 OK and text “Server is running”.
GET /configs: Returns the list of all configurations as a JSON array.
POST /configs: Accepts a JSON body, creates a new config, and adds it to the memory store.
GET /configs/{key}: Returns a specific config by its Key. If not found, return 404.
Implementation Details:
Use http.HandleFunc to route requests.
For the POST request, you must read the r.Body, decode the JSON into your struct, and append it to your slice/map.
Set the correct headers: w.Header().Set("Content-Type", "application/json").
Testing (The DevOps Part):
You must be able to test this using curl or Postman.
Example test command:
bash
        curl -X POST -H "Content-Type: application/json" -d '{"key":"db_host", "value":"localhost", "service":"payment"}' http://localhost:8080/configs
Success Criteria (for when you code it):

The server runs on port 8080.
You can add data via POST and immediately see it via GET.
You handle invalid JSON gracefully (don’t crash the server).
You strictly use the standard library net/http (don’t use frameworks like Gin or Echo yet—you need to know how the raw library works first).
