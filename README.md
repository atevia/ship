# 🚀 Ship — Secure Deploy & Transfer Tool for VPS

**Ship** is a CLI tool created by **Atevia** to securely transfer, synchronize and deploy projects between Linux servers using SSH.

Ship is designed for **developers, agencies and small teams** that deploy applications to VPS servers and want a **simple, secure and modern workflow** without heavy infrastructure tools.

---

## ✨ What is Ship?

Ship solves three common problems when working with VPS deployments:

### 1️⃣ Secure transfers (`push`)
- Packages your project
- Ignores unnecessary files
- Calculates SHA256 checksum
- Transfers via SSH
- Verifies integrity
- Extracts safely on the destination

### 2️⃣ Fast synchronization (`sync`)
- Uses **rsync** internally
- Sends only file deltas
- Keeps Ship’s clean UX (profiles, flags, ignore)

### 3️⃣ Atomic deployments (`deploy`)
- Release-based deployments
- `current` / `previous` symlinks
- Zero-downtime switch
- Instant rollback

---

## 📦 Installation

### Install using the installer script
```bash
curl -fsSL https://raw.githubusercontent.com/atevia/ship/main/scripts/install.sh | bash
````

Verify:

```bash
ship help
```

---

## ⚙️ Profiles (recommended)

Profiles avoid repeating SSH credentials on every command.

### Create config

```bash
ship init
```

This creates:

```text
~/.ship/config
```

### Example config

```ini
[prod]
user=atevia
host=51.222.15.232
key=~/.ssh/id_ed25519
port=22

[staging]
user=ubuntu
host=3.227.242.240
key=~/.ssh/lightsail_id
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
ship test prod
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

Override manually:

```bash
--shipignore /path/to/.shipignore
```

---

## 🚚 Commands

---

## 🔹 `ship push`

Secure transfer using packaged artifacts.

### Usage

```bash
ship push <source> <destination> [flags]
```

### Example

```bash
ship push ./foco-app prod:/var/www/foco-app --progress
```

### What it does

1. Packages the project
2. Generates SHA256 checksum
3. Transfers via SSH
4. Verifies checksum remotely
5. Extracts safely

---

## 🔹 `ship sync`

Fast incremental sync using rsync.

### Usage

```bash
ship sync <source> <destination> [flags]
```

### Example

```bash
ship sync ./foco-app prod:/var/www/foco-app --progress
```

Best for frequent updates.

---

## 🔹 `ship deploy`

Atomic deployment with releases and rollback.

### Usage

```bash
ship deploy <source> <destination> [flags]
```

### Example (Next.js + PM2)

```bash
ship deploy ./foco-app prod:/var/www/foco-app \
  --after "cd /var/www/foco-app/current && pnpm install --frozen-lockfile && pnpm build && pm2 reload foco" \
  --keep 7 \
  --progress
```

### Remote structure

```text
/var/www/foco-app/
├── releases/
│   ├── 20240128-142233/
│   ├── 20240128-150912/
│   └── ...
├── current -> releases/20240128-150912
└── previous -> releases/20240128-142233
```

---

## 🔁 `ship rollback`

Rollback to the previous release.

### Usage

```bash
ship rollback <destination> [flags]
```

### Example

```bash
ship rollback prod:/var/www/foco-app --after "pm2 reload foco"
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

Example:

```bash
--after "pm2 reload foco"
```

---

### 🚀 Deploy-only flags

| Flag                    | Description                             |
| ----------------------- | --------------------------------------- |
| `--keep <n>`            | Number of releases to keep (default: 5) |
| `--release-name <name>` | Custom release name                     |

---

## ℹ️ Help & version

```bash
ship help
ship version
```

---

## 🧠 Design philosophy

Ship does **not** try to replace:

* Kubernetes
* Ansible
* Terraform

Ship **does replace**:

* fragile rsync scripts
* `tar | ssh | tar` one-liners
* unsafe overwrites
* downtime caused by manual deploys

---

## 🏆 When to use Ship

✅ VPS
✅ Web applications
✅ Node / Next.js / React
✅ Agencies & freelancers
✅ Simple but professional infra

❌ Large clusters
❌ Continuous bidirectional sync
❌ Backup solutions (use borg/restic)

---

## 🔐 Security

* SSH key authentication
* Key permission validation
* End-to-end checksum verification
* Atomic deploys
* Instant rollback

---

## 🛣 Roadmap

* Healthcheck-based deploy validation
* Automatic rollback
* Remote builds
* Multi-arch releases
* CI/CD integration

---
