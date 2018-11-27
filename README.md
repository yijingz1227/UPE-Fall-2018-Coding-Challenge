# 0. How to build and run my solution 
#### i. First you need to install Go on your PC/Laptop
##### Step 1: 
Use this official link: https://golang.org/doc/install to download according to your platform.
The Go version I used here is "go version go1.10.3 darwin/amd64", which can be checked after installation with command` go version`. 

##### Step 2 (Crucial):
After the installation, make sure you create the workspace directory, which can be `$HOME/go`.
Under the workspace directory, create a` src` folder, all your Go source codes shoud go here.
(This is very crucial due to the way Go manages packages and source codes)

##### Step 3 (Optional but recommended):
You can follow the Test your installation part in the offcial link to test if your installation is successful. 

#### ii. After the installation and to build my solution
Step 1: create a new folder under the `$HOME/go/src` directory

Step 2: Do the `git init`, `git pull` ... using this GitHub repo

Step 3: Once you have main.go, rest.go, maze.go in your folder, `go build *.go` will give you an executable named `main`. 

Step 4: To run the executable, `./main`

Step 5: You should see a prompt on your command-line tool "Mission starts." 

# 1. Module explanation
#### The file contains 3 .go files, main, rest, and maze
**main.go** contains the main logic of the program, which starts a timer, prints a starting prompt, and gives a summary report when the maze is solved.

**rest.go** contains all the functions that make HTTP requests

**maze.go** contains the loop that calls DFS algorithm to solve the maze


