# Contributing to Ship

Thank you for your interest in contributing to **Ship** 🚀

We welcome contributions that improve stability, security, usability and documentation.

---

## 📋 Guidelines

### 1️⃣ Philosophy
Ship aims to be:
- Simple
- Secure
- Predictable
- Production-ready

Please avoid adding features that:
- Increase complexity unnecessarily
- Break backward compatibility without strong reasons
- Depend on heavy external services

---

## 🧱 Project Structure
````

.
├── cmd/ship        # CLI entrypoint
├── scripts         # Installer and helper scripts
├── README.md
├── CONTRIBUTING.md
└── go.mod

````

---

## 🛠 Development Setup

Requirements:
- Go 1.22+
- Linux environment

Clone the repo:
```bash
git clone https://github.com/atevia/ship.git
cd ship
````

Build locally:

```bash
go build -o ship ./cmd/ship
```

Run:

```bash
./ship help
```

---

## 🧪 Testing

Currently, Ship relies on:

* manual testing
* real SSH environments

If you add logic-heavy features:

* prefer pure functions
* add unit tests where applicable

---

## 🧹 Code Style

* Follow standard Go formatting (`gofmt`)
* Keep functions small and readable
* Prefer clarity over cleverness

---

## 🚀 Submitting Changes

1. Fork the repository
2. Create a feature branch
3. Commit your changes with clear messages
4. Open a Pull Request

Example commit message:

```
feat: add atomic rollback command
```

---

## 🐛 Bug Reports

Please include:

* OS and architecture
* Ship version
* Command used
* Full error output

---

## 📄 License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

Thanks for helping improve Ship ❤️
