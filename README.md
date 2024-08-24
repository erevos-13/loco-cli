# Loco CLI for translation of the project.

- If you have a project in the [LOCO](https://localise.biz/) you can use this CLI to download the translations and save them in the project.

## Installation

- You need to download the binary from the root the `loco-cli-go` is for mac and the `loco-cli-go.exe` is for windows.

- You can also build the project with the following command:

```bash
▶  go build
```

## Usage

- To execute the binary we need to pass the params in the following way:

```bash
loco-cli-go.exe <TOKEN of the LOCO> <Path that you have the translation> <LOCALE> <path you want to save the translation>
```

**You can also run the cli and filed the params that you need to pass.**

- Example:

```bash
loco-cli-go.exe sdfsfd4449dfdsAAA /translation/translate.json  en /src/assets/i18n
```

# License

- MIT [LICENSE](https://opensource.org/license/MIT)

