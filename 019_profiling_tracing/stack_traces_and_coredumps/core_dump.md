## Core Dumps

To generate a core dump from a running / hung program -

1. On Linux use `gcore` utility.
   - Make sure `ulimit -c unlimited` is set.
   - Run 
     `# gcore -a -o core <pid | list of pids>`
   - The program continues to run after the core is dumped.
2. On MacOS send `SIGQUIT` (`kill -SIGQUIT <pid>`) or hit `CTRL-\`

*NOTE* - If `GOTRACEBACK` is set to `crash` then a larger core file is dumped for the go program. 
*NOTE* - `dlv` recognizes `gcore` core file.

Use the generated core dump and use `dlv` to analyse the stack trace, variables, goroutines etc.

`dlv core ./your/program/binary ./path/to/your/gcore-core-file`
