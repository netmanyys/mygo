Error message "go: go.mod file not found in current directory or any parent directory; see 'go help modules'"

Change this:

go env -w GO111MODULE=auto

to this:

go env -w GO111MODULE=off