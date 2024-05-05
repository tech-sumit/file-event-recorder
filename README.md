
# 🚀 File Event Recorder

## Overview
This project provides a solution for monitoring file events within a specified directory and handling these events with concurrency control. It uses a goroutine-based event executor to efficiently manage file operations like creation, modification, and deletion.

## 🛠 Environment Setup

Before running the project, ensure that your environment is set up correctly. You should configure your settings in an `.env` file based on the example provided below:

```env
# service config
APP_ENV=development

# DB Config
DB_PATH=sql/fileStorage.sqlite

# Concurrency setting
MAX_CONCURRENCY=100

# Directory to watch
DIR_PATH=tmp
```

## Setup Instructions

1. **Clone the repository**: Clone this repository to your local machine or download the source code.

2. **Configure environment variables**: Create an `.env` file in the root directory of the project, following the structure of the `.env.example` provided above. Adjust the values as necessary for your environment.

3. **Install dependencies**: Ensure that Go is installed on your machine and then install the required Go modules by running:
    ```bash
    go mod tidy
    ```

4. **Build the project**:
    ```bash
    go build -o fileEventRecorder
    ```

5. **Run the application**:
    ```bash
    ./fileEventRecorder
    ```

## Event Executor

The event executor is a core component of this application, responsible for handling filesystem events asynchronously. It operates using a pool of goroutines, each listening for file events such as create, update, delete, and rename. This concurrency model allows the application to efficiently process multiple events in parallel without blocking the main execution flow.

Each event triggers a specific goroutine tasked with handling that event type, ensuring that file operations are managed swiftly and effectively. The concurrency level, controlled by the `MAX_CONCURRENCY` setting, limits the number of goroutines running simultaneously, preventing system overload and optimizing resource usage.

## Using the Application

- **Start the application** using the command line by navigating to the project's directory and running the built executable. Ensure your `.env` file is correctly configured as the application depends on these settings to operate.

- **Monitoring**: The application will start monitoring the directory specified by `DIR_PATH` in the `.env` file and will process events based on the configured concurrency level.

## Contributing

Contributions to this project are welcome! Please feel free to fork the repository, make changes, and submit pull requests.

## License

This project is open-sourced under the MIT License.
