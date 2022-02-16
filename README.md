[![Build](https://github.com/nao1215/subaru/actions/workflows/build.yml/badge.svg)](https://github.com/nao1215/subaru/actions/workflows/build.yml)
[![UnitTest](https://github.com/nao1215/subaru/actions/workflows/unit_test.yml/badge.svg)](https://github.com/nao1215/subaru/actions/workflows/unit_test.yml)
# subaru - inspired by fortune command
subaru command print philosophy or wise sayings. If you are not specified subcommand, subaru print the random phrase. The subcommand name is a person's name, programming language, or company name, etc. If you execute subaru command with a subcommand, subaru print the phrase related to that name.

**You can use this project to make your first Pull Requests.**  We are waiting for PR from you. Your name record in AUTHORS.md!! If you are interested in, see [CONTRIBUTING.md](./CONTRIBUTING.md)

# How to install
## Step.1 Install golang
If you don't install golang in your system, please install Golang first. Check the [Go official website](https://go.dev/doc/install) for how to install golang.
## Step2. Install subaru
```
$ go install github.com/nao1215/subaru@latest
```
  
# How to use
## Random output
```
$ subaru 
[BENJAMIN FRANKLIN]
Drive thy business; let it not drive thee.

$ subaru 
[MRS GUMP]
Life was like a box of chocolates. You never know what you’re gonna get.

$ subaru 
[CHRISTOPHER CHAMBERS]
Kids lose everything, unless there’s someone there to look out for them. And id your parents are too fucked-up to do it, then maybe I should. 

$ subaru 
[GOETHE]
Chess is the touchstone of intellect.

$ subaru 
[SOCRATES]
Thou shouldst eat to live; not live to eat.

$ subaru 
[EDOGAWA KONAN]
One truth prevails.

$ subaru 
[FYODOR DOSTOEVSKY]
...whether I am a trembling creature or whether I have the right...

$ subaru 
[MARK TWAIN]
To succeed in life, you need two things; ignorance and confidence.

$ subaru 
[CHARLES DARWIN]
A man who dares to waste one hour of time has not discovered the value of life.

$ subaru 
[HELEN KELLER]
Keep your face to the sunshine and you cannot see the shadow.
```
## Specifying the subcommand
```
$ subaru atari
[atari]
Business is a good game. Lots of competition and a minimum of rules.
You keep score with money.
```
If you want to know the list of subcommands, execute "subaru --help".

# Contact
If you would like to send comments such as "find a bug" or "request for additional features" to the developer, please use one of the following contacts.

- [GitHub Issue](https://github.com/nao1215/subaru/issues)

# LICENSE
The subaru project is licensed under the terms of [the Apache License 2.0](./LICENSE).
