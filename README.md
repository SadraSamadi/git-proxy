# git-proxy

A tiny CLI for managing and switching git proxies.

Save proxies by name, switch between them instantly, and keep the git config clean, all from one command.

---

## Build

**Requirements:** Go 1.18+

```bash
git clone https://github.com/SadraSamadi/git-proxy.git
```

```bash
cd git-proxy
```

Unix:
```bash
go build -o git-proxy ./cmd/main
```

Windows:
```bash
go build -o git-proxy.exe ./cmd/main
```

---

## Usage

```txt
Git Proxy Manager

Usage:
  git-proxy save <key> <value>   Save a proxy  (e.g. save test socks5://127.0.0.1:1080)
  git-proxy list                 List all saved proxies
  git-proxy remove <key>         Remove a saved proxy
  git-proxy use <key>            Configure git to use a saved proxy
  git-proxy current              Show the current git proxy
  git-proxy unset                Unset the current git proxy
```

Proxies are saved in a `data.json` file in the current working directory.

Accepted proxy formats: `http://`, `socks5://`, `socks5h://`
