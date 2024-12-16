### README.md

# 🚀 Go Test Tasks Project


Each `task/` directory contains:
- The Go implementation for the task. 📝
- A `task_test.go` file with test cases for the task. 🔍

---

## 🎯 How to Run

### Using `make` (Recommended)
Run the main entry point (`main.go`) using the `make` command:

```bash
make run
```

### Without `make`
If `make` is not available, you can directly execute the project using Go:

```bash
go run main.go
```

---

## 🧪 How to Test

Testing is a core part of this project. Each task has a test suite located in its respective directory. You can run all tests at once or test specific tasks.

### Run All Tests
To run tests across the entire project:

```bash
go test ./...
```

### Test a Specific Task
To test an individual task:

```bash
cd task1
go test
```
