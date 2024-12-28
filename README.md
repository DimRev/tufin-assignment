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

### Info

As the tufin-assignment binary will try to read and write files to the OS temp directory, it is recommended to run it as root via `sudo`.

Use a the -h flag on each one of the commands to see the available flags for the given command.

```bash
tufin-assignment status -pv -n example
```

- The -pv flag will show the status of all PersistentVolumeClaims and Pods in the example namespace.

## CHANGELOG

### v1.0.0 - Finalize the assignment

- App deploys a k3s cluster
- App deploys two pods: MySQL and WordPress
- App shows status

### v1.1.0 - Add Helm support

- Add ability to parse flags and execute different commands based on them
- Status now can show pods/services/PVCs in a given namespace
- Add Helm support
- Deploy has a helm flag to deploy via a embedded Helm chart
