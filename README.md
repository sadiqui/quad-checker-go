# Quadchecker

## Description

Quadchecker is a command-line program written in Go that helps identify and analyze quad functions. It takes a string as an argument and determines whether it matches any quad function. If a match is found, Quadchecker displays the name of the matching quad and its dimensions. If no match is found, it prints "Not a quad function".

This program is based on the quad functions, the same ones in my rectangle-golang repository. I recoded them here in a simpler manner.

This program takes a string as an argument and displays the name of the matching quad and its dimensions.

If the argument is not a quad, the program should print "Not a quad function".

All answers end with a newline (`'\n'`).

If there is more than one quad match, the program must display them all alphabetically ordered and separated by a `||`.

## Usage

To use Quadchecker, follow these steps:

1. Ensure you have Go installed on your system.

2. Clone the repository to your local machine:

   ```bash
   git clone <repository_url>
   ```

**NOTE:** The provided built file are compatible with Linux (Ubuntu) environment, if you are on a different environment, remove the built files, and run `go build` command, to have suitable executables.

## Examples:

```bash
$ ls -l
-rw-r--r-- 1 asadiqui student   28 Mar  5 00:37 go.mod
drwxr-xr-x 2 asadiqui student 4096 Mar  5 00:14 main
drwxr-xr-x 2 asadiqui student 4096 Mar  5 00:27 quadA
drwxr-xr-x 2 asadiqui student 4096 Mar  5 00:26 quadB
drwxr-xr-x 2 asadiqui student 4096 Mar  5 00:26 quadC
drwxr-xr-x 2 asadiqui student 4096 Mar  5 00:26 quadD
drwxr-xr-x 2 asadiqui student 4096 Mar  5 00:26 quadE
-rw-r--r-- 1 asadiqui student  898 Mar  5 09:44 README.md
$ ./quadA 3 3 | go run .
[quadA] [3] [3]
$ ./quadC 1 1
A
$ ./quadD 1 1
A
$ ./quadE 1 1
A
$ ./quadC 1 1 | go run .
[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]
$ ./quadE 1 2
A
C
$ ./quadC 1 2
A
C
$ ./quadE 1 2 | go run .
[quadC] [1] [2] || [quadE] [1] [2]
$ echo 0 0 | go run .
Not a quad function
$ echo -n "o--o"$'\n'"|"$'\n'"o"
o--o
|
o$ echo -n "o--o"$'\n'"|"$'\n'"o" | go run .
Not a quad function
```