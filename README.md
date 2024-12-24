# tufin-assignment

## Installation

```bash
go install github.com/DimRev/tufin-assignment@latest
```

## Usage

```bash
tufin-assignment <command> [flags]
```

where `<command>` is one of the following:

- `cluster` - will deploy a k3s cluster
- `deploy` - will deploy two pods: MySQL and WordPress
- `status` - will print the status table of pods in the default namespace
- `remove` - will remove the k3s cluster

### Flags

- `--help, -h` - will print details about the command
- `--version, -v` - will print the version of the program

## Example

```bash
tufin-assignment cluster --help
```
