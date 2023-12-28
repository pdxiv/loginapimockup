# Simple Go Login API

## Overview

This project demonstrates a simple mockup of a login API using Go. It serves as a basic example of handling REST API login functionality with hardcoded credentials. The project showcases:

* A REST API endpoint for handling login requests.
* Hardcoded username and password ("admin" and "password") for demonstration purposes.
* An embedded HTML page in the Go binary to interact with the API.
* Basic handling of HTTP requests and responses in Go.

## Features

* **Login API Endpoint**: The server exposes a `/api/login` endpoint that accepts POST requests with JSON-formatted credentials.
* **Embedded Web Interface**: The `index.html` file, containing a simple user interface for login, is embedded within the Go binary, simplifying deployment and execution.
* **Basic Authentication Logic**: The API compares incoming credentials with the predefined username and password and returns appropriate HTTP status codes based on the validity of the credentials.

## How to Run

1. **Build the Project**: Compile the project into a binary using Go. Run the following command in the project directory:

```bash
go build
```

2. **Run the Server**: Start the server by executing the compiled binary. On Unix-like systems, you can use:

```bash
./loginapimockup
```

The server will start on `http://localhost:8080` .

3. **Access the Web Interface**: Open a web browser and navigate to `http://localhost:8080`. Use the web interface to interact with the login API.

4. **Login**: Enter the username and password ("admin" and "password") and click the login button. The API will authenticate the credentials and display a success or failure message.

## Technology Stack

* **Language**: Go (Golang)
* **Web Technologies**: HTML, JavaScript

## Important Notes

* This project is a basic demonstration and should not be used as-is in production environments.
* The login mechanism uses hardcoded credentials, which is not a secure practice for real-world applications.
* Proper error handling, validation, and security measures (like HTTPS, password hashing, and token-based authentication) should be implemented for production use.

## Creating Releases with GitHub Actions

This project uses GitHub Actions to automate the building and releasing of artifacts for different platforms. When a new version of the software is ready to be released, you can trigger the build and release process by creating and pushing a Git tag.

### Prerequisites

* Ensure you have Git installed on your system.
* You should have push access to the repository.

### Creating a New Release

1. **Check the Latest Version**: First, check the latest release version so you can tag the next version accordingly.

2. **Create a Git Tag**: Create a new Git tag corresponding to the new version of your software. Use semantic versioning for your tags (e.g., `v0.0.3`). To create a tag, use the following command:

```bash
git tag v0.0.3
```

Replace `0.0.3` with your desired version number.

3. **Push the Tag to GitHub**: Push the newly created tag to GitHub. This push will trigger the GitHub Actions workflow which builds and releases the software.

```bash
git push origin v0.0.3
```

Again, replace `0.0.3` with the version number of your tag.

### GitHub Actions Workflow

Once the tag is pushed to GitHub, the following will occur automatically:

* **Build**: GitHub Actions will start the build process for the specified platforms (e.g., Apple ARM, Apple Intel, Raspberry Pi).
* **Release**: After successful builds, a new release will be created on GitHub, and the compiled binaries for each platform will be attached to the release.

### Checking the Release

* After the GitHub Actions workflow completes, you can check the `Releases` section of the GitHub repository for the new release and its assets.
