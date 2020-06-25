# Borf language specification

This is heavily based off of the summary at <https://skilldrick.github.io/easyforth/>.

## Reserved Words

The following symbols are reserved:

	+
	-
	*
	/
	.
	pop
	:				Maybe?
	;
	->				Function declaration?
	duplicate
	drop
	swap
	ascii
	newline
	."
	pop"
	=
	<
	>
	and
	or
	not
	if
	else
	then
	variable
	constant
	allocate

Tentatively reserved:

	do
	loop
	begin
	until

Special REPL commands:

	size
	peek
	show
	stack
	bye

## Syntax

Symbols (words) are whitespace delimited, therefore, even the `;` symbol must be separated from other runes by whitespace.