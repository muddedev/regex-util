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

```
Match: [32mgopher[0m [32mgodot[0m [32mgo[0m [32mgo[0m!
```

_\*Matching terms are highlighted in green_
