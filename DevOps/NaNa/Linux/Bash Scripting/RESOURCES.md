# Bash Scripting Resources

## Knowledge

- [GNU Bash Reference Manual: Shell Scripts](https://www.gnu.org/software/bash/manual/html_node/Shell-Scripts.html)
  Primary source for script files, invoking Bash, executable permission, and the interpreter line. Use for Lesson 1 and later script execution questions.
- [GNU Bash Reference Manual: Bash Startup Files](https://www.gnu.org/software/bash/manual/html_node/Bash-Startup-Files.html)
  Explains which startup files are read by interactive and non-interactive shells. Use when investigating why a CI job does not see an interactive terminal's configuration.
- [GNU Bash Reference Manual: Quoting](https://www.gnu.org/software/bash/manual/html_node/Quoting.html)
  Primary source for single quotes, double quotes, and escaping. Use when a value contains spaces or shell-special characters.
- [GNU Bash Reference Manual: Shell Parameter Expansion](https://www.gnu.org/software/bash/manual/html_node/Shell-Parameter-Expansion.html)
  Covers variable references, braces, and default or required values. Use for deployment configuration and script inputs.
- [GNU Bash Reference Manual: Shell Expansions](https://www.gnu.org/software/bash/manual/html_node/Shell-Expansions.html)
  Describes expansion order, word splitting, and filename expansion. Use when command arguments behave differently than expected.
- [ShellCheck SC2086: Double quote to prevent globbing and word splitting](https://www.shellcheck.net/wiki/SC2086)
  Practical explanation of a common unquoted-expansion hazard. Use when reviewing scripts for arguments that may contain spaces or wildcard characters.
- [GNU Bash Reference Manual: Positional Parameters](https://www.gnu.org/software/bash/manual/html_node/Positional-Parameters.html)
  Defines arguments such as `$1` and their relationship to a script invocation. Use for Lesson 3.
- [GNU Bash Reference Manual: Special Parameters](https://www.gnu.org/software/bash/manual/html_node/Special-Parameters.html)
  Documents `$#`, `"$@"`, and `$?`. Use when reading script inputs and command status.
- [GNU Bash Reference Manual: Exit Status](https://www.gnu.org/software/bash/manual/html_node/Exit-Status.html)
  Explains that zero means success and non-zero means failure in Bash's command model. Use for Lesson 4.
- [GNU Bash Reference Manual: Conditional Constructs](https://www.gnu.org/software/bash/manual/html_node/Conditional-Constructs.html)
  Primary reference for `if`, `case`, and their evaluation rules. Use for Lesson 5.
- [GNU Bash Reference Manual: Conditional Expressions](https://www.gnu.org/software/bash/manual/html_node/Bash-Conditional-Expressions.html)
  Lists file, string, and numeric tests used by `[[ ... ]]`. Use for Lesson 4 preflight checks.
- [GNU Bash Reference Manual: Looping Constructs](https://www.gnu.org/software/bash/manual/html_node/Looping-Constructs.html)
  Primary reference for `for`, `while`, and loop control. Use for Lesson 6 and future batch-processing lessons.
- *The Software Developer’s Guide to Linux*, Chapter 12, “Automating Tasks with Shell Scripts,” pp. 151–166 (local book PDF)
  Main book sequence for Bash scripting and developer workflows. Use as the course's primary book reference.
- *Linux Bible*, 11th edition, Chapter 7, “Writing Simple Shell Scripts,” pp. 145–161 (local book PDF)
  Adds administrator examples and broader shell practice.
- *CompTIA Linux+ Guide to Linux Certification*, Chapter 7, “Working with the BASH Shell,” pp. 320–334 (local book PDF; variable coverage pp. 312–323)
  Reinforces Bash shell fundamentals and certification-oriented practice.

## Wisdom

- [Unix & Linux Stack Exchange](https://unix.stackexchange.com/tour)
  A moderated Q&A community that accepts practical shell scripting and Linux administration questions. Use it to compare real operational approaches after checking the Bash manual and showing a minimal example.

## Gaps

- No trusted resource has been selected yet for later Bash security, portability, and testing lessons. Add focused primary references when those topics enter the sequence.
