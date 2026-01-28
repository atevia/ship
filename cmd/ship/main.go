package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Profile struct {
	User string
	Host string
	Key  string
	Port string
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "help", "--help", "-h":
		printHelp()
	case "push":
		handlePush()
	case "test":
		handleTest()
	default:
		fmt.Println("❌ Unknown command:", os.Args[1])
		printHelp()
	}
}

/* =========================
   HELP
========================= */

func printHelp() {
	fmt.Println(`
ship – Atevia internal transfer tool

USAGE:
  ship push <source> <target> [flags]
  ship test <profile>
  ship help

TARGET:
  user@host:/path
  profile:/path

FLAGS:
  --key <path>        SSH private key
  --port <port>       SSH port (default 22)
  --profile <name>    Use profile from ~/.ship/config
  --progress          Show transfer progress (uses pv)

FILES:
  .shipignore         Ignore patterns (gitignore-like, basic)

EXAMPLES:
  ship push ./app ubuntu@1.2.3.4:/var/www/app --key ~/.ssh/id_ed25519
  ship push ./app prod:/var/www/app --progress
  ship test prod
`)
}

/* =========================
   PUSH
========================= */

func handlePush() {
	if len(os.Args) < 4 {
		printHelp()
		os.Exit(1)
	}

	source := os.Args[2]
	target := os.Args[3]

	key := ""
	port := "22"
	profileName := ""
	progress := false

	for i := 4; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--key":
			key = os.Args[i+1]
			i++
		case "--port":
			port = os.Args[i+1]
			i++
		case "--profile":
			profileName = os.Args[i+1]
			i++
		case "--progress":
			progress = true
		}
	}

	sshUserHost, remotePath := resolveTarget(target, profileName)

	if key == "" && profileName != "" {
		key = loadProfile(profileName).Key
	}

	if key != "" {
		validateKey(key)
	}

	excludeArgs := buildShipIgnore()

	pv := ""
	if progress && commandExists("pv") {
		pv = "pv |"
	}

	sshCmd := fmt.Sprintf("ssh -p %s", port)
	if key != "" {
		sshCmd = fmt.Sprintf("ssh -i %s -p %s", key, port)
	}

	cmd := exec.Command(
		"bash",
		"-c",
		fmt.Sprintf(
			`tar %s -czf - %s | %s %s %s "mkdir -p %s && tar -xzf - -C %s"`,
			excludeArgs,
			source,
			pv,
			sshCmd,
			sshUserHost,
			remotePath,
			remotePath,
		),
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Transfer failed:", err)
		os.Exit(1)
	}

	fmt.Println("✅ Transfer completed successfully")
}

/* =========================
   TEST
========================= */

func handleTest() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ship test <profile>")
		os.Exit(1)
	}

	profile := loadProfile(os.Args[2])

	fmt.Println("✔ Profile loaded")

	validateKey(profile.Key)
	fmt.Println("✔ SSH key valid")

	cmd := exec.Command(
		"ssh",
		"-i", profile.Key,
		"-p", profile.Port,
		fmt.Sprintf("%s@%s", profile.User, profile.Host),
		"exit",
	)

	if err := cmd.Run(); err != nil {
		fmt.Println("❌ SSH test failed:", err)
		os.Exit(1)
	}

	fmt.Println("✅ SSH connection OK")
}

/* =========================
   PROFILES
========================= */

func loadProfile(name string) Profile {
	configPath := filepath.Join(os.Getenv("HOME"), ".ship", "config")

	file, err := os.Open(configPath)
	if err != nil {
		fmt.Println("❌ Cannot open profile config:", configPath)
		os.Exit(1)
	}
	defer file.Close()

	var p Profile
	current := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "[") {
			current = strings.Trim(line, "[]")
			continue
		}

		if current == name {
			parts := strings.SplitN(line, "=", 2)
			switch parts[0] {
			case "user":
				p.User = parts[1]
			case "host":
				p.Host = parts[1]
			case "key":
				p.Key = parts[1]
			case "port":
				p.Port = parts[1]
			}
		}
	}

	if p.Port == "" {
		p.Port = "22"
	}

	if p.User == "" || p.Host == "" {
		fmt.Println("❌ Invalid profile:", name)
		os.Exit(1)
	}

	return p
}

func resolveTarget(target, profile string) (string, string) {
	parts := strings.SplitN(target, ":", 2)
	if len(parts) != 2 {
		fmt.Println("❌ Invalid target format")
		os.Exit(1)
	}

	if profile != "" || !strings.Contains(parts[0], "@") {
		p := loadProfile(parts[0])
		return fmt.Sprintf("%s@%s", p.User, p.Host), parts[1]
	}

	return parts[0], parts[1]
}

/* =========================
   IGNORE
========================= */

func buildShipIgnore() string {
	if _, err := os.Stat(".shipignore"); err == nil {
		return "--exclude-from=.shipignore"
	}
	return ""
}

/* =========================
   UTILS
========================= */

func validateKey(path string) {
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("❌ SSH key not found:", path)
		os.Exit(1)
	}

	if info.Mode().Perm() > 0600 {
		fmt.Println("❌ SSH key permissions too open (use chmod 600)")
		os.Exit(1)
	}
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
