# Hello World GoLang

This repository contains a simple Go application that prints "hello world",
along with a randomly generated word continuously every 5 seconds. It serves as
a basic demonstration of Go's capabilities, including concurrency, external
libraries, and logging.

Additionally, this application incorporates a basic cryptographic operation by
leveraging Go's standard crypto/sha256 library to generate a SHA-256 hash of the
randomly generated word in each iteration. Please note that while this
demonstrates cryptographic hashing, Go's standard cryptographic libraries are
not FIPS 140-2 validated.

## Why add
While the application utilizes Go's crypto/sha256 for hashing, please note that
Go's standard library, including its cryptographic packages, is not
FIPS 140-2 validated.

## Prerequisites

Before you begin, ensure you have Go installed on your system (version 1.15 or
later recommended).

## Cloning the repository

Clone the repository to your local machine to get started with this project:

```bash
git clone https://github.com/marmccor/hello-world-golang.git
cd hello-world-golang
```

## Building the application

o compile the program into a binary, run the following command in the project
directory:

```bash
go build -o hello-world-golang
```

This command compiles the source code into a binary named hello-world-golang.

## Running the application

After building the application, you can run it directly from the command line:

```bash
./hello-world-golang
```

Example output:

```bash
% ./hello-world-golang
INFO[0000] hello world - iteration 1 at 2024-02-27 11:17:13 - random word: Beetleberyl - hash: 5dc4927eacf941b994d5b98eb5451887394e8f2ce2050e05eaa2481faf81661c
INFO[0005] hello world - iteration 2 at 2024-02-27 11:17:18 - random word: Catchertrail - hash: ed35b39871fa62b9f91e9dd11618ff851917a336c771b9157ec108836017202a
INFO[0010] hello world - iteration 3 at 2024-02-27 11:17:23 - random word: Biterbrindle - hash: 0f98eb5e6b046363c42d5bd36a8e6d99dd1e0d08a1d7ea629541ae7de9987cfd
INFO[0015] hello world - iteration 4 at 2024-02-27 11:17:28 - random word: Thumbhail - hash: 5cb5b9dd0327e1eee59305c343459f1cf9cb3dc116ccd0743771a64650690e90
INFO[0020] hello world - iteration 5 at 2024-02-27 11:17:33 - random word: Kittenvivid - hash: 55d50efe20193f12f2e1eb26ba2069fd4422a294364d74489537c8e72ed25234
INFO[0025] hello world - iteration 6 at 2024-02-27 11:17:38 - random word: Duckpeppermint - hash: 19299196b763ffce9afd13280c268aab0a9eae237a2e3c06bac625a21f824f47
INFO[0030] hello world - iteration 7 at 2024-02-27 11:17:43 - random word: Footwool - hash: a3a3fd79f76761091199925cbc19d548fb238897a9eca2f3e9d66d743e658110
```

## Updating dependencies

To update a specific dependency to a new version, you can use the go get command
followed by the dependency path and the desired version. For example, to update
github.com/sirupsen/logrus to version v1.9.0, you would run:

```bash
go get github.com/sirupsen/logrus@v1.9.0
```

Which will update the pinned version in go.mod.
