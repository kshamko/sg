# Description

2 ways of calculations were researched (please check notebooks folder for the details). Option 2 was chosen for the implementation as it's a bit easier to implement and less computational heavy

## Run the util. Examples

```bash
$ go run cmd/predict/main.go --help
```

```bash
$ go run cmd/predict/main.go --model=linex --source=test_data.json --aggregate=country
```

```bash
$ go run cmd/predict/main.go --model=linex --source=test_data.csv --aggregate=country
```

```bash
$ go run cmd/predict/main.go --model=expsmoothing --source=test_data.csv --aggregate=campaign
```

```bash
$ go run cmd/predict/main.go --model=expsmoothing --source=test_data.json --aggregate=campaign
```