# About
ncbi-fetch is a command line tool to retrieve biological information you want to get from NCBI's database.<br>
### What is NCBI?
The **NCBI** (National Center for Biotechnology Information) is an organization in the United States <br>that provides a vast collection of online tools and databases for researchers studying biology and medicine. 
# Installation

0. Prepare go in your environment.

* If you want to use virtual envs insread, you can use the Dockerfile to acquire a container in which the cli work.
```bash
docker build -t $IMAGE .
docker run --rm $IMAGE
```

# Usage
```
Usage:
  ncbi-fetch [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  connect     connect to the ftp server
  get         get the data from the NCBI database
  help        Help about any command
  version     Print the version number

Flags:
  -h, --help   help for ncbi-fetch

Use "ncbi-fetch [command] --help" for more information about a command.
```

```
go run main.go get taxonid -s "Escherichia coli","Faecalibacterium prausnitzii"
```
# License

"ncbi-fetch" is under [MIT license](https://en.wikipedia.org/wiki/MIT_License).
