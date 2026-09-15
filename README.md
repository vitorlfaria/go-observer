# Go Observer CLI

A high-performance, concurrent command-line interface for multi-service observability and health checking, written entirely in Go using standard libraries.

## 🚀 Features

* **Extreme Concurrency:** Utilizes Go Routines and Channels to perform simultaneous HTTP health checks across dozens of services.
* **Safe Error Handling:** Robust DNS and timeout error catching without panicking or halting the execution of other routines.
* **Zero Dependencies:** Built exclusively with the Go standard library (`net/http`, `sync`, `encoding/json`, `flag`), ensuring a minimal attack surface and lightweight binary.
* **Custom HTTP Client:** Prevents socket hanging and connection pool exhaustion with strictly configured timeouts.

## 🛠️ Prerequisites

* [Go 1.22+](https://go.dev/dl/) installed.

## 📦 Installation

Clone the repository and build the binary:

```bash
git clone [https://github.com/your-username/go-observer.git](https://github.com/your-username/go-observer.git)
cd go-observer
go build -o observer main.go
```

## ⚙️ Configuration

Create a config.json file containing the URLs you want to monitor:
```json
{
  "urls": [
    "[https://google.com](https://google.com)",
    "[https://github.com](https://github.com)",
    "[https://invalid-url-for-testing.local](https://invalid-url-for-testing.local)"
  ]
}
```

## 💻 Usage

Run the compiled binary, passing the configuration file via the -config flag:

```bash
./observer -config=config.json
```

### Example Output
```
--- Observability Report ---
Url: [https://invalid-url-for-testing.local](https://invalid-url-for-testing.local)
Status: 
Duration: 15.2ms
Error: Get "[https://invalid-url-for-testing.local](https://invalid-url-for-testing.local)": dial tcp: lookup invalid-url-for-testing.local: no such host
----- // -----
Url: [https://github.com](https://github.com)
Status: 200 OK
Duration: 120.5ms
Error: <nil>
----- // -----
Url: [https://google.com](https://google.com)
Status: 200 OK
Duration: 340.1ms
Error: <nil>
----- // -----
```

## 🧠 Architecture Highlights

- WaitGroups (sync.WaitGroup): Ensures the main thread waits for all background HTTP requests to complete.

- Channels (chan): Provides thread-safe communication to aggregate results from multiple goroutines without relying on memory-locking mechanisms (Mutexes).