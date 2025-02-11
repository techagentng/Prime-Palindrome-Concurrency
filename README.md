# Concurrent Log Processor

A Go program to efficiently process large log files using concurrency. It counts occurrences of specific keywords and outputs their frequencies in descending order.

---

## Features

- **Concurrency**: Uses Goroutines to process log files in chunks.
- **Keyword Counting**: Tracks occurrences of specified keywords.
- **Sorting**: Outputs keyword frequencies in descending order.
- **Error Handling**: Handles file not found and empty file errors.

---

## Setup

1. **Clone the Repository**:
   ```bash
   git clone <repository-url>
   cd concurrent-log-processor
