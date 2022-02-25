# Cat Herder

Cat Herder is a Go tool for preventing configuration drift between .env files

## Installation

You can build the tool using Go

```bash
go build .
```

## Usage

The tool is used by passing an expected .env and a .env you want to test against it. It will 
return exit(0) if both files contain the same keys (it doesn't check values), and exit(1) if
the tested file is missing any keys or has additional keys. It will also print the keys at 
fault.

```bash
./catherder <expected_env_path> <testing_env_path>
```


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)