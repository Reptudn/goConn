# Publishing goConn as a Go Package - Checklist

## Pre-Publication Checklist

- [ ] Remove or move `main.go` from root to `examples/basic_bot/main.go`
- [ ] Add godoc comments to all public functions in `bot/` package
- [ ] Add godoc comments to all public types in `shared/` package
- [ ] Ensure `internal/` package is truly private (not imported by public API)
- [ ] Update `go.mod` with proper module path: `module github.com/Reptudn/goConn`
- [ ] Create `README.md` with quick start guide
- [ ] Create `LICENSE` file (MIT recommended)
- [ ] Create `CONTRIBUTING.md` for developers

## Before First Release

- [ ] Test imports work correctly
  ```bash
  # In a new directory
  go get github.com/Reptudn/goConn
  ```
- [ ] Run all tests: `go test ./...`
- [ ] Check code coverage: `go test -cover ./...`
- [ ] Run linter: `go vet ./...`
- [ ] Format code: `go fmt ./...`

## Publication Steps

1. **Push to GitHub**
   ```bash
   git add .
   git commit -m "Ready for public release"
   git push -u origin main
   ```

2. **Create a Release Tag**
   ```bash
   git tag -a v1.0.0 -m "Initial public release"
   git push origin v1.0.0
   ```

3. **Create GitHub Release**
   - Go to: https://github.com/Reptudn/goConn/releases
   - Click "Create a new release"
   - Tag version: `v1.0.0`
   - Title: "goConn v1.0.0 - Initial Release"
   - Description: Copy from README.md
   - Publish release

4. **Wait for Indexing**
   - Go pkg.go.dev checks every ~10 minutes
   - View at: https://pkg.go.dev/github.com/Reptudn/goConn
   - May take up to 30 minutes for initial indexing

## After Publication

- [ ] Verify package is visible on pkg.go.dev
- [ ] Verify `go get` works in a fresh environment
- [ ] Add badge to README.md:
  ```markdown
  [![Go Reference](https://pkg.go.dev/badge/github.com/Reptudn/goConn.svg)](https://pkg.go.dev/github.com/Reptudn/goConn)
  ```

## Version Management (Future Releases)

For each new release:

```bash
# Create and tag the release
git tag -a v1.1.0 -m "Add new features"
git push origin v1.1.0

# Users can then use:
go get github.com/Reptudn/goConn@v1.1.0
go get github.com/Reptudn/goConn@latest
```

## Public API Best Practices

✅ **Do This:**
- Expose high-level functions: `CreateUnit()`, `Move()`
- Hide implementation: Don't expose socket or protocol details
- Use clear naming: `GetTeamUnits()` not `FetchTeamUnits()`
- Document with examples

❌ **Don't Do This:**
- Don't export internal socket functions
- Don't change public API frequently
- Don't break backwards compatibility without major version

## Troubleshooting

### Package not appearing on pkg.go.dev
- Ensure you have a valid GitHub repository
- Make sure you created a GitHub release (not just a tag)
- Wait 30 minutes for indexing

### Import errors
- Check module path in `go.mod` matches your GitHub path
- Ensure all public functions have godoc comments
- Run `go mod tidy` locally

### Users can't find your package
- Make sure GitHub repository is public
- Create a proper GitHub release
- Add meaningful description to releases


