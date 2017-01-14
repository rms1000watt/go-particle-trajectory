# Create project path
mkdir go-particle-trajectory
cd go-particle-trajectory

# Install and run govendor
go get -u github.com/kardianos/govendor
govendor fetch github.com/spf13/pflag
govendor fetch github.com/spf13/cobra/cobra

# Run Cobra
go run vendor/github.com/spf13/cobra/cobra/main.go init
go run vendor/github.com/spf13/cobra/cobra/main.go add start

# Create a bin directory
mkdir bin

# Create a src directory
mkdir src
touch src/main.go
# Edit src/main.go
