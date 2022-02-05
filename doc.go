// subaru print philosophy or wise sayings.
//
// If no subcommand is specified, the random philosophy will be output. The subcommand
// name is a person's name, programming language, or company name. If you execute subaru
// command with a subcommand, the phrase related to that name is output.
//
// **Welcome to add philosophy or wise sayings.**  You can add subcommands in the following ways:
// 1. Add "XXX.subaru"(XXX is any sub command name) under the fortune directory.
// 2. Write any text in the created XXX.subaru.
// The subaru command includes the \*.subaru files in the fortune directory at build time. When
// the subaru command is executed, the embedded fortune/\*.subaru file is read and the subcommand
// is created.
package main
