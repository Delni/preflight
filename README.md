# Preflight 🛫

> Automate your tooling checklist with Preflight

![Go version](https://img.shields.io/github/go-mod/go-version/delni/preflight?style=flat&color=00ADD8)
![Last Commit](https://img.shields.io/github/last-commit/delni/preflight?style=flat)
![Contributors](https://img.shields.io/github/contributors/delni/preflight?style=flat)
[![License](https://img.shields.io/github/license/delni/preflight?style=flat)](./LICENSE)
![Build](https://github.com/Delni/preflight/actions/workflows/ci.yml/badge.svg)
[![Coverage](https://delni.github.io/preflight/coverage-badge.svg)](https://delni.github.io/preflight/index.html)

Preflight will run a list of commands for you to make sure that you are ready to go on your journey. This can be usefull when you are reinstalling your computer and are used to some configuration, or to make sure that the onboarding in a new team is complete.  

![demo](./docs/demo.local.gif)

## Install ⬇️

For now, simply download [preflight](https://github.com/Delni/preflight/releases/latest) and run it. More options will come 💪

## Usage 🚀

```
preflight path/to/yaml/descriptor
```

The `path/to/yaml/descriptor` is a `YAML` file with a list of *System Checks*. You can find an exemple in [docs/demo.yaml](./docs/demo.yaml). Feel free to download and edit it to fit your needs, it showcases every possibilities.

### Don't have the file locally ?

Preflight got you! If you have an URL, Preflight will fetch the file for you and run as usual after that. Just pass the `--remote` (`-r`) flag:

![demo.remote](./docs/demo.remote.gif)

### How to write a checklist 👨‍✈

Firstly, remember that this is check-**list**. The top level object is a **list**. Each entry is called a `SystemCheck`:

| Key         |      Type      | Mandatory | Description |
| :---------- | :------------: | :-------: | :---------- |
| name        |    `string`    |    yes    |  Name of the category of commands | 
| description |    `string`    |    no     |  Explain why this check is being made. (Ex: you should be able to manage different node versions). |
| optional    |   `boolean`    |    no     | Default to `false`. If set to true, can fail. It will still display the verose message |  
| checkpoints     | `[]Checkpoint` |    yes    | The list of options. At least *one* should pass to make the SystemCheck green |

And each `Checkpoint` is defined as follow:
| Key           |      Type      | Mandatory | Description |
| :------------ | :------------: | :-------: | :---------- |
| name          |    `string`    |    yes    |  Name of the command / label to display | 
| command       |    `string`    |    yes    |  the actual command passed to the runner. By default, preflight will only assert if the command exist on your machine. |
| use_interactive      |   `boolean`    |    no     | Default to `false`. Usefull when you need to execute a command in interactive shell (it runs .bashrc file before, usefull for nvm) |  
| documentation |    `string`    |    yes    |  In case of failure, display some info to the user about what to do next. This should at least be a link to the documentation on how to install said command. |

## Roadmap 🚦

- [x] File based decriptor
- [x] fetch file from public URL
- [ ] Make it installable seamlessly
- [ ] Add flag-based descriptor for major use-cases. Go fileless  
    Currently supported Systems:
    - Missing some config ? Feel free to [open an issue](https://github.com/Delni/preflight/issues/new) to discuss it, and read [related docs](./presets/README.md)!

## How to contribute 📝

➡️ See [CONTRIBUTING](./CONTRIBUTING.md)  
➡️ Want to add a new preset ? See [presets docs](./presets/README.md)
