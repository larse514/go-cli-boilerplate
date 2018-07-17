# hello-cli
CLI for saying hello! And also learning go

## commands

| command  | parameters   | description                                  |
|----------|--------------|----------------------------------------------|
| hello    | name         | command to say hello to someone              |


## usage

### build
```bash
$ make
```
The make command will run the clean, dependencies, test, and build commands from the [Makefile](https://github.com/larse514/go-cli-boilerplate/blob/master/Makefile)

### run 
```bash
$ ./hello-cli hello larse514
Hello "larse514"
```

### cleanup
```bash
$ make clean
``` 
