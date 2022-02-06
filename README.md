[![Build](https://github.com/nao1215/subaru/actions/workflows/build.yml/badge.svg)](https://github.com/nao1215/subaru/actions/workflows/build.yml)
[![UnitTest](https://github.com/nao1215/subaru/actions/workflows/unit_test.yml/badge.svg)](https://github.com/nao1215/subaru/actions/workflows/unit_test.yml)
# subaru - inspired by fortune command
subaru command print philosophy or wise sayings. If you are not specified subcommand, subaru print the random phrase. The subcommand name is a person's name, programming language, or company name, etc. If you execute subaru command with a subcommand, subaru print the phrase related to that name.

# How to install
## Step.1 Install golang
If you don't install golang in your system, please install Golang first. Check the [Go official website](https://go.dev/doc/install) for how to install golang.
## Step2. Install subaru
```
$ go install github.com/nao1215/subaru@latest
```
  
# How to use
## Ranodom output
```
$ subaru
[nietzsche]
Love is more afraid of change than destruction.
```
## Specifying the subcommand
```
$ subaru atari
[atari]
Business is a good game. Lots of competition and a minimum of rules.
You keep score with money.
```
If you want to know the list of subcommands, execute "subaru --help".

# Contribution
**Welcome to add philosophy or wise sayings.**    
If you add the phrase to subaru, just add the text file(*.subaru) in the fortune directory. That's all. You can add phrase (subcommands) in the following ways:

1. Add "XXX.subaru" under the fortune directory. XXX is the name of the person (or tool, company, etc.) related to phrase. XXX is extracted and used as the subcommand name at runtime.
2. Write any text in the created XXX.subaru.

The subaru command embedd the .subaru files in the fortune directory at build time. When the subaru command is executed, the embedded fortune/.subaru file is read and the subcommand is created.  
If you contribute, I will record your name in AUTHORS.md.

# Contact
If you would like to send comments such as "find a bug" or "request for additional features" to the developer, please use one of the following contacts.

- [GitHub Issue](https://github.com/nao1215/subaru/issues)

# LICENSE
The subaru project is licensed under the terms of [the Apache License 2.0](./LICENSE).
