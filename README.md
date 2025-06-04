# 🤖 JSON TOGO

**Simple JSON to Go struct converter.**

Parses JSON input and generates Go struct definitions with JSON tags.

---

## ✨ Features

- Parses valid JSON and generates idiomatic Go structs.
- Adds appropriate `json` tags to struct fields.
- Supports optional flags:
  - 📦 Custom package name
  - 🧱 Custom struct name
  - 💾 Output to `.go` file

---

## 📦 Installation

```bash
go install github.com/chaewonkong/json-togo@v1.0.0
```

## 🛠 Usage

```bash
jtg [flags] < input.json
```

## 🔸 With Custom Flags

| Flag        | Shorthand | Description                        | Default              |
| ----------- | --------- | ---------------------------------- | -------------------- |
| `--package` | `-p`      | Package name in the generated file | `main`               |
| `--struct`  | `-s`      | Name of the root struct            | `Data`               |
| `--output`  | `-o`      | Path to output `.go` file          | *(prints to stdout)* |

### Example Usage

```bash
jtg -p user -s UserProfile -o user_profile.go < user.json
```

Generates a model.go file with contents like:

(in `user_profile.go`)

```go
package user

type UserProfile struct {
    // ...
}
```
