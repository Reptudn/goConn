# Contributing to goConn

## Development Setup

```bash
git clone https://github.com/Reptudn/goConn.git
cd goConn
go mod download
go test ./...
```

## Project Structure

- `bot/` - Main Bot API (public interface)
- `shared/` - Game objects, types, and types (public)
- `internal/` - Internal implementation details (not for external use)
- `actions/` - Action queue system
- `main.go` - Example usage (remove before releasing as a package)

## Code Guidelines

1. Keep the public API minimal - only expose what users need
2. Use interfaces for extensibility
3. Write clear error messages
4. Add godoc comments to public functions
5. Use type-safe patterns with Go's type system

## Testing

```bash
go test -v ./...
go test -cover ./...
```

## Release Process

1. Update version in `go.mod` and tags
2. Update CHANGELOG.md
3. Create a GitHub release
4. Tag the release: `git tag v1.0.0`


