# sysmonitor
 This is a small tool built in go to get the system metrics. For beginners, there is a brief steps on how to install and setup go for development. This tool development will involve utilizing the following Go concepts:
 - Variables and data types
 - Code Organization
 - Access and read system information using  gopsutil library.
 - Go Routines to concurrently read and display system information.
 - Communication between goroutines using Channels.
 - Terminal user interface.
 - Debugging
 - Unit Testing
 - Error management in go.

### Setting up go for development involves 3 steps:
- Install Go and its toolchains:- 
  - Go to the Go projects [home page](https://go.dev/) and click Downloads and than download the binary as per your OS. If you are using MacOS, you can also use homebrew to install and manage go.
  - To validate if it is installed successfully run the following commands in the CMD/terminal: `go version`.You should see the latest go version installed, OS and platform it is running. 
  - When you run the `go` command you can see a lot of options to run the command e.g. `go <command> [argument]`.
- Setup code editors:
  - We will be using Visual Studio Code here.
  - Download the VSCode as per your OS from the [official site](https://code.visualstudio.com/Download).
  - Install the go extensions in the editor required for the development.
- Create your first go programme:
  - Open the Terminal and navigate to your project folder.
  - When you execute `go mod`, you can see a lot of options to run this command.
  - Than execute this command to create a module `go mod init user_define_module_name`.
  - Create the main.go file with main function which is the entry point of the application.
  - Inorder to build and run your application execute this command in the terminal from your project directory `go run .`
