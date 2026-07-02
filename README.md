# CSV Streaming Pipeline in Go

A high-performance streaming pipeline built in Go for reading, parsing, validating, and transforming large volumes of data (benchmarked on datasets of 1,000,000+ rows).

This project is designed with a strong focus on memory efficiency (Zero/Low Heap Allocations) and strictly follows the Separation of Concerns architectural pattern.

## 🚀 Key Features

* **Streaming Architecture:** Data is processed row-by-row using `csv.Reader`. Memory consumption remains low and fixed regardless of the input CSV file size.
* **Blazing Fast Performance:** Capable of processing over **1.5 million rows per second**.
* **Memory-Optimized:** Advanced Go optimizations (like `ReuseRecord` and fast-path ASCII checking for UTF-8 strings) cut total heap allocation size by more than half.
* **Clean Architecture:** Processing logic is split into standalone, easily testable components: Technical Parsing (`Parse`), Business Rule Validation (`Validate`), and Data Cleaning/Mutation (`Transform`).

---

## 📂 Project Structure

```text
csvreader/
├── cmd/
│   └── main.go           # Entry point to execute the pipeline and output metrics
├── internal/
│   ├── models.go         # Domain model definition (Company, Stats structs)
|   ├── pipeline_test.go  # Benchmark 
│   ├── pipeline.go       # Core pipeline execution loop 
│   └── steps.go          # Pipeline processing stages: Parse, Validate, and Transform
└── go.mod