# cli-task-manager

The **lightweight issue tracker for terminal** concept aims to provide a practical, fast, and simple tool especially for developers working in terminal environments. This type of tool allows developers to manage tasks, bugs, requests, and other notes in their projects through the terminal. This tool is **CLI (Command Line Interface)** based and can work with small files or structured data files like JSON or YAML without requiring a database or large infrastructure.

## Installation

### Prerequisites

- Go 1.18 or higher

### Installation from Source

```bash
# Clone the repository
git clone https://github.com/mstgnz/cli-task-manager.git
cd cli-task-manager

# Build the application
make build

# Install the application (optional)
make install
```

### Using Docker

You can also run the application using Docker without installing Go:

```bash
# Build the Docker image
make docker-build

# Run the application with Docker (uses a persistent volume for data)
make docker-run CMD="help"

# Examples
make docker-run CMD="list"
make docker-run CMD="add \"Create API documentation\" --label feature"
```

The Docker setup uses a named volume (`cli-task-manager-data`) to persist your tasks between container runs.

## Usage

### Commands

1. **Add Task:**

   ```bash
   issue-tracker add "Create API draft for new feature" --label "feature"
   ```

2. **List Tasks:**

   ```bash
   issue-tracker list
   ```

3. **Update Status:**

   ```bash
   issue-tracker update 1 --status "in-progress"
   ```

4. **Filter by Label:**

   ```bash
   issue-tracker filter --label "bug"
   ```

5. **Remove Task:**
   ```bash
   issue-tracker remove 3
   ```

### Example Outputs

#### Task List:

```bash
$ issue-tracker list
1. [Feature] Create API draft for new feature   [Status: To-Do]
2. [Bug] User login screen error                [Status: In-Progress]
3. [Task] Update README file                    [Status: Done]
```

#### Adding a Task:

```bash
$ issue-tracker add "New user registration form" --label "feature"
Task successfully added: [Feature] New user registration form [Status: To-Do]
```

#### Filtering:

```bash
$ issue-tracker filter --label "bug"
1. [Bug] User login screen error [Status: In-Progress]
```

## Technical Details

- **Data Storage:** Tasks are stored in JSON format in the `.cli-task-manager/tasks.json` file in the user's home directory.
- **Status Types:** Tasks can be in three different states: `to-do`, `in-progress`, `done`.
- **Labels:** Special labels can be assigned to tasks (e.g., `feature`, `bug`, `task`).

## Development

### Project Structure

```
cli-task-manager/
├── cmd/
│   └── main.go      # Main application entry point
├── models/          # Data models
├── storage/         # Data storage operations
|── commands/        # Command handlers
├── Dockerfile       # Docker configuration
├── Makefile         # Build and installation commands
└── README.md        # Project documentation
```

### Tests

You can use the following commands to test the project:

```bash
# Run all tests
make test

# View test coverage
make test-coverage
```

### Contributing

This project is open-source, and contributions are welcome. Feel free to contribute or provide feedback of any kind.

## License

This project is licensed under the [MIT License](LICENSE).
