<pre>
                                      ███████╗██╗  ██╗██╗██████╗ 
                                      ██╔════╝██║  ██║██║██╔══██╗
                                      ███████╗███████║██║██████╔╝
                                      ╚════██║██╔══██║██║██╔═══╝ 
                                      ███████║██║  ██║██║██║     
                                      ╚══════╝╚═╝  ╚═╝╚═╝╚═╝     
                                      
                                      Secure Deploy &amp; Transfer Tool
</pre>

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green)
![Platform](https://img.shields.io/badge/Platform-Linux-blue)
![Status](https://img.shields.io/badge/Status-Stable-success)

**Ship** is a modern, secure and opinionated CLI tool to **transfer, synchronize and deploy projects to Linux servers using SSH**.

It is designed for developers, freelancers and small teams that deploy applications to VPS servers and want a **simple, reliable and production-ready workflow** without heavy infrastructure tools.

---

## ✨ Why Ship?

Ship focuses on **Developer Experience (DX)** and **production safety**:

- Secure SSH-based transfers
- Atomic deployments with instant rollback
- Profiles for multiple servers
- Gitignore-style file exclusions
- Progress indicators
- Zero-downtime deploys

---

## 🧠 What Ship Solves

### 1️⃣ Secure transfers (`push`)
- Packages your project
- Excludes unnecessary files
- Generates SHA256 checksum
- Transfers via SSH
- Verifies integrity
- Extracts safely on the destination

### 2️⃣ Fast synchronization (`sync`)
- Uses **rsync** internally
- Transfers only file deltas
- Ideal for frequent updates

### 3️⃣ Atomic deployments (`deploy`)
- Release-based directory structure
- `current` / `previous` symlinks
- Instant rollback if needed
- Zero-downtime switches

---

## 📦 Installation

### Install using the installer script
```bash
curl -fsSL https://raw.githubusercontent.com/atevia/ship/main/scripts/install.sh | bash
````

Verify installation:

```bash
ship help
```

---

## ⚙️ Profiles (recommended)

Profiles allow you to reuse SSH configuration without repeating flags.

### Initialize configuration

```bash
ship init
```

This creates:

```text
~/.ship/config
```

### Example configuration

```ini
[production]
user=deploy
host=203.0.113.10
key=~/.ssh/id_ed25519
port=22

[staging]
user=ubuntu
host=198.51.100.25
key=~/.ssh/id_ed25519
port=22
```

### List profiles

```bash
ship profiles
```

---

## 🔐 Validate connectivity

Before transferring or deploying:

```bash
ship test production
```

This checks:

* profile validity
* SSH key existence
* key permissions
* SSH connectivity

---

## 📂 Ignoring files (`.shipignore`)

Ship supports a **gitignore-style** `.shipignore` file.

### Example

```gitignore
node_modules
.next
.env*
!.env.example
```

Ship automatically detects:

* `<project>/.shipignore`
* `./.shipignore`

Manual override:

```bash
--shipignore /path/to/.shipignore
```

---

## 🚚 Commands

---

## 🔹 `ship push`

Secure artifact-based transfer.

### Usage

```bash
ship push <source> <destination> [flags]
```

### Example

```bash
ship push ./my-app production:/var/www/my-app --progress
```

---

## 🔹 `ship sync`

Fast incremental synchronization.

### Usage

```bash
ship sync <source> <destination> [flags]
```

### Example

```bash
ship sync ./my-app production:/var/www/my-app --progress
```

Best for frequent updates.

---

## 🔹 `ship deploy`

Atomic deployment with releases and rollback.

### Usage

```bash
ship deploy <source> <destination> [flags]
```

### Example

```bash
ship deploy ./my-app production:/var/www/my-app \
  --after "cd /var/www/my-app/current && npm install --production && pm2 reload my-app" \
  --keep 5 \
  --progress
```

### Remote structure

```text
/var/www/my-app/
├── releases/
│   ├── 20240128-142233/
│   ├── 20240128-150912/
│   └── ...
├── current -> releases/20240128-150912
└── previous -> releases/20240128-142233
```

---

## 🔁 `ship rollback`

Rollback to the previous release instantly.

```bash
ship rollback production:/var/www/my-app --after "pm2 reload my-app"
```

---

## 📘 Flags Reference

### 🔧 Global flags

| Flag                  | Description                                          |
| --------------------- | ---------------------------------------------------- |
| `--profile <name>`    | Use a profile from `~/.ship/config`                  |
| `--key <path>`        | SSH private key (overrides profile)                  |
| `--port <port>`       | SSH port (default: 22)                               |
| `--sudo`              | Run remote commands using `sudo`                     |
| `--progress`          | Show real transfer progress (uses `pv` if available) |
| `--shipignore <path>` | Explicit path to `.shipignore`                       |

---

### 🪝 Hooks

| Flag               | Description                                 |
| ------------------ | ------------------------------------------- |
| `--before "<cmd>"` | Run remote command before extract or switch |
| `--after "<cmd>"`  | Run remote command after extract or switch  |

---

### 🚀 Deploy-only flags

| Flag                    | Description                             |
| ----------------------- | --------------------------------------- |
| `--keep <n>`            | Number of releases to keep (default: 5) |
| `--release-name <name>` | Custom release name                     |

---

## 🏆 When to use Ship

✅ VPS deployments
✅ Web applications
✅ Node / Next.js / React
✅ Small teams & agencies
✅ Simple but professional infrastructure

❌ Large clusters
❌ Continuous bidirectional sync
❌ Backup solutions (use borg/restic)

---

## 🔐 Security

* SSH key authentication
* Key permission validation
* End-to-end checksum verification
* Atomic deployments
* Instant rollback

---

## 🛣 Roadmap

* Healthcheck-based deploy validation
* Automatic rollback
* Remote builds
* Multi-arch releases
* CI/CD integration

---

## 📄 License

MIT








