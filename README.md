# Music DataDeck

Welcome to Music DataDeck. 

## Getting Started

First, you have to clone the project from GitHub into your workspace.

### 1. Cloning the project

Run the following command to clone the project. Clone the project in the src folder of your worksapce.

```
git clone https://github.com/kenquiros64/music-data-deck.git
```

### 2. Installing dependencies

The next step you must do is to install all dependecies.It is necessary to have installed glide. Run the command below to install all dependencies.

```
glide install
```

## Project configuration

### Prerequisites

What things you need to install the software and how to install them. For this project we will use Golang 1.9.7.

```
Golang 1.9.7
goji.io
```

### Installing Go on Linux, macOS

A step by step series of how to install Golang 1.9.7 on Linux and macOS.

Download the archive from https://golang.org/doc/install?download=go1.9.7.linux-amd64.tar.gz

Then extract it into /usr/local, creating a Go tree in /usr/local/go.

```
sudo tar -C /usr/local -xzf go1.9.7.linux-amd64.tar.gz
```

Add /usr/local/go/bin to the PATH environment variable. You can do this by adding this line to your /etc/profile (for a system-wide installation) or $HOME/.profile:

```
export PATH=$PATH:/usr/local/go/bin
```

### Installing Go on Microsoft

A step by step series of how to install Golang 1.9.7 on Microsoft.

Download the archive from https://golang.org/doc/install?download=go1.9.7.windows-386.msi

Open the MSI file and follow the prompts to install the Go tools.

### Installing Goji

Install Goji by running the following in your terminal:

```
go get goji.io
```

### Add More Dependencies 

It's neccesary to install glide for add more dependecies.

```
glide get github.com/foo/bar
```

Or

```
glide get github.com/foo/bar#^1.2.3
```

