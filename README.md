# Zootoma

<p align="center">
  <img src="https://i.ibb.co/7bPjgrx/zootoma-removebg-preview-Convert-Image.jpg">
</p>

Zootoma is the key-value in memory database that aims to implement path like data structures for data clustering, in addition to being a distributed configuration database for another clusters.

## Build

In the project root directory:

```bash
go build -o zootoma cmd/zootoma/main.go
```

## Execution

In the project root directory:

```bash
./zootoma -port=9009
```

It will run zootoma at port 9009 of localhost

## Protocol

Zootoma application protocol is composed by three blocks:
- Main Header: Composed by the fundamental parts of application basic functionality, such as method(get and set), Key to the data and the Data Length(Interger)
- Meta Header: Composed by line separated key value pairs for further implementation of adicional functionality, such as partitioning and indexing
- Data: Data of size Data Length


```txt
<Method> <Key> <Data Length>    #Main Header

<key>::<value>                  #Meta Header
<key>::<value>

<Data>                          #Data
```

### Protocol example

```txt
set /data/foo 7

partition::xpto

footoma
```


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. There are lots of things to do, fill free to join us in this challenge!!
