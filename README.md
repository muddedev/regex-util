# regex-util

A simple CLI tool to test and highlight regular expression matches in your terminal.

---

## Usage

```bash
go run main.go --pattern='<your-regex>'
```

### Example

```bash
go run main.go --pattern='go[a-z]+'
```

Input:

```
gopher godot go go!
```

Output:

![Regex match demo](example/example.png)
