# Contribution
**Welcome to add philosophy or wise sayings.**    
If you add the phrase to subaru, just add the text file(*.subaru) in the fortune directory. That's all. Even if you are not good at programming, you can make a pull request.

# You can add phrase (subcommands) in the following ways:

1. Fork the subaru command.
2. Add "XXX.subaru" under the fortune directory that forked one. XXX is the name of the person (or tool, company, etc.) related to phrase. XXX is extracted and used as the subcommand name at runtime.
3. Write any text in the created XXX.subaru.

The subaru command embedd the .subaru files in the fortune directory at build time. When the subaru command is executed, the embedded fortune/.subaru file is read and the subcommand is created.  
If you contribute, I will record your name in AUTHORS.md.