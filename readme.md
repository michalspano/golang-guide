# Golang Guide

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/512px-Go_Logo_Blue.svg.png" alt="Golang Logo" width="250" height="auto" align="right">

Welcome to the Golang Guide repository! This project aims to provide a comprehensive introduction to the `Go` programming language. While it's not a complete guide, it serves as an excellent starting point for beginners. Here, you'll find essential information about `Go`, useful tips and tricks, and a collection of beginner-friendly projects to help you understand the language better.

## Table of Contents

- [Golang Guide](#golang-guide)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Installation](#installation)
  - [Getting Started](#getting-started)
    - [Running a Go File](#running-a-go-file)
    - [Compiling a Go File into an Executable Binary](#compiling-a-go-file-into-an-executable-binary)
    - [Executing the Binary](#executing-the-binary)
    - [Compiling a Go File into the `bin` Directory](#compiling-a-go-file-into-the-bin-directory)
  - [Project Structure](#project-structure)
  - [Usage](#usage)
  - [Contributing](#contributing)

## Introduction

`Go`, also known as `Golang`, is an open-source programming language designed for efficiency and simplicity. With its strong focus on readability, concurrency, and built-in support for concurrent programming, `Go` is becoming increasingly popular among developers.

This repository serves as a beginner-friendly guide to help you get started with `Go`. It covers the basics of the language, provides useful tips, and offers hands-on projects for practice.

## Installation

To use `Go` on your machine, you need to install it first. Follow the official `Go` installation guide for your operating system:

- [Official `Go` **Installation Guide**](https://golang.org/doc/install)

## Getting Started

Once you have `Go` installed, you can start writing and executing `Go` programs. Here are a few essential commands to get you started:

### Running a Go File

To run a `Go` file, use the following command:

```shell
$ go run src.go
```

### Compiling a Go File into an Executable Binary

To compile a `Go` file into a binary executable, use the following command:

```shell
$ go build src.go
```

### Executing the Binary

After compiling a `Go` file, you can execute the binary directly:

```shell
$ ./src
```

### Compiling a Go File into the `bin` Directory

To compile a `Go` file into the `bin` directory, use the following command:

```shell
$ go install src.go
```

For more detailed information and advanced usage, refer to the official [**`Go` documentation**](https://go.dev/doc/).

## Project Structure

The project repository is structured as follows:

```
.
├── src/
│   ├── src.go
│   └── ...
├── bin/
│   ├── src
│   └── ...
```

- The `src` directory contains the source code for various projects. Feel free to explore and experiment with them.
- The `bin` directory contains the compiled binaries for the projects in the `src` directory. You can execute them directly without compiling them again.

## Usage

Clone the repository to your local machine:

```shell
$ git clone https://github.com/michalspano/golang-guide.git
```

Mentioned above, the [_"Getting Started"_](#getting-started) section contains the essential commands to get you started with `Go`. For more detailed information and advanced usage, refer to the official [`Go` documentation](https://go.dev/doc/).

## Contributing

Contributions are welcome! If you have any suggestions, bug fixes, or improvements, feel free to *open* an **issue** or *submit* a **pull request**.