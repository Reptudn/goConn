# Quick Reference: Publishing goConn

## TL;DR - 5 Steps to Publish

### Step 1: Move main.go to examples
```bash
mkdir -p examples/basic_bot
mv main.go examples/basic_bot/
```

### Step 2: Verify go.mod
```bash
cat go.mod
# Should show: module github.com/Reptudn/goConn
```

### Step 3: Push to GitHub
```bash
git add .
git commit -m "Prepare for public release"
git push origin main
```

### Step 4: Create Release Tag
```bash
git tag v1.0.0
git push origin v1.0.0
```

### Step 5: Create GitHub Release
1. Go to: https://github.com/Reptudn/goConn/releases/new
2. Select tag: v1.0.0
3. Title: "goConn v1.0.0"
4. Click "Publish release"

---

## Done! 🎉

Users can now install your package:
```bash
go get github.com/Reptudn/goConn
```

Package will appear on pkg.go.dev within 30 minutes.

---

## Public API Visibility

### Users can access:
```go
import goconn "github.com/Reptudn/goConn"
import "github.com/Reptudn/goConn/shared"

goconn.NewCoreGameBot()
goconn.(*CoreGameBot).CreateUnit()
goconn.(*CoreGameBot).Move()
shared.Game
shared.Object
shared.Position
```

### Users CANNOT access (good!):
```go
// These are hidden because they're in internal/
internal.Connection
internal.Socket
```

---

## Files Created for You:

- ✅ `README.md` - Package overview
- ✅ `LICENSE` - MIT license
- ✅ `CONTRIBUTING.md` - Development guide
- ✅ `PACKAGE_GUIDE.md` - Detailed setup guide
- ✅ `CHECKLIST.md` - Full publication checklist
- ✅ `examples/basic_bot/main.go` - Example usage
- ✅ Updated `bot.go` with godoc comments

---

## Next: Add Godoc Comments

Add to top of each public package:

```go
// Package goConn provides a clean API for game server bot control.
package goConn

// Exported types need comments:
// Type Name does something useful.
type TypeName struct {
    // fields...
}

// Function names need comments:
// FunctionName does something useful.
func FunctionName() error {
    // ...
}
```

Example in `shared/objects.go`:
```go
// Package shared contains public game objects and types.
package shared
```

---

## Verification

Test locally:
```bash
# Create a test directory
mkdir test-import
cd test-import
go mod init test-app

# Add your package
go get github.com/Reptudn/goConn

# Try importing
cat > main.go << 'EOF'
package main

import goconn "github.com/Reptudn/goConn"

func main() {
    _ = goconn.CoreGameBot{}
}
EOF

go build  # Should work!
```

---

## Common Issues

**"go get not working?"**
- Make sure you pushed to GitHub
- Make sure repository is public
- Wait 5 minutes for caching

**"Can't import internal package?"**
- Good! It's supposed to be hidden
- Only use `goConn` and `shared` packages

**"Package not on pkg.go.dev?"**
- Created a GitHub release? (Important!)
- Just a tag won't work
- Wait 30 minutes for indexing

