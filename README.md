# Music DataDeck

Welcome to Music DataDeck. 

## Getting Started

First, you have to clone the project from GitHub into your workspace.

### 1. Cloning the project

Run the following command to clone the project. Clone the project in the $GOPATH/src folder of your worksapce.

```
git clone https://github.com/kenquiros64/music-data-deck.git
```

### 2. Installing dependencies

The next step you must do is to install all dependecies.It is necessary to have installed glide. Run the command below to install all dependencies.

```
glide install
```

### 3. Running the project

To run the project you should build it first, to do this you must go to the project folder and execute the command below:

```
go build
```

To execute the project, you must execute the following command

```
./music-data-deck
```

Now the project is running on http://localhost:8000. In the next step, it shows how the API can be consumed.

### 3. Consuming API

The table below shows the different actions that can be consumed via API. The project is running on http://localhost:8000.

| ACTION                 | GET                     | COMMENTS                                                                                                |
| ---------------------- | ----------------------- | ------------------------------------------------------------------------------------------------------- |
| All songs              | /songs                  | Get all songs stored on database.                                                                       |
| Search songs           | /songs/search/:value    | Search all the songs that matches with the given value. You can search it by artist, song name or genre.|
| Search genres          | /genres/search/:value   | Search all the genres that matches with the given value. You can search it by name.                     |
| Genre Song Info        | /genres/info            | Get a list of the genres, and the number of songs and the total length of all the songs by genre.       |
| Search songs by length | /songs/search/:min/:max | Get all songs from database between minimun and maximum length.                                         |

List of examples

| ACTION                 | GET                     | EXAMPLES                                                                                                |
| ---------------------- | ------------------------| ------------------------------------------------------------------------------------------------------- |
| Search songs           | /songs/search/:value    | /songs/search/424                                                                                      |
| Search genres          | /genres/search/:value   | /genres/search/Classic%20Rock                                                                           |
| Search songs by length | /songs/search/:min/:max | /songs/search/200/245                                         |

### 3. Testing application

To execute all the tests, you must run the following command in the terminal. Go to the project folder and execute:

```
go test
```

## Project configuration

### Prerequisites

What things you need to install the software and how to install them. For this project we will use Golang 1.9.7.

```
Golang 1.9.7
goji.io
github.com/mattn/go-sqlite3
github.com/zenazn/goji/web
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


## Additonal questions

1. Are there any suggestions you could make about the structure of our data?

First, create a new table Artist with id and name. This could help us to separate entities with the correct structure and attributes.

The maximum length of the varchar is very large, it could be made smaller.

2. What fields would you index in these tables?

For Songs, I could index id and genre.
For Genres, I could index just the id.

