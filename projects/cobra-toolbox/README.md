
# install cobra lib
```BASH
go get -u github.com/spf13/cobra@latest
```


# install cobra cli
```BASH
go install github.com/spf13/cobra-cli@latest
```


# initilize project
```BASH
cobra-cli init cobra-toolbox
```

# Verify Your Project Directory
```BASH
# your project is located at
~/dev/mygo/projects/cobra-toolbox
```

# Check Your $GOPATH
```BASH
echo $GOPATH
# If the output is empty, you are likely using Go modules (which is the default for modern Go versions). If GOPATH is set, ensure your project is not dependent on the older GOPATH-based structure.
```

#

# Initialize the Module
```BASH
go mod init cobra-toolbox
# This creates a go.mod file, which tells Go to use the module system. You can name the module cobra-toolbox or use a path like github.com/<username>/cobra-toolbox.
```

# Fix the import statement
```GO
import "cobra-toolbox/cmd"
```

```BASH
$ tree cobra-toolbox/
cobra-toolbox/
├── cmd
│   └── root.go
├── go.mod
├── go.sum
├── LICENSE
├── main.go
└── README.md
```


# Download dependencies and cleansup the go.dmo file
```
go mod tidy
```


# Run the program
```BASH
go run main.go
# A longer description that spans multiple lines and likely contains
# examples and usage of using your application. For example:
# 
# Cobra is a CLI library for Go that empowers applications.
# This application is a tool to generate the needed files
# to quickly create a Cobra application.


```

# url https://cobra.dev/