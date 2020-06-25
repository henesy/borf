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
	;
	->
	duplicate
	drop
	swap
	ascii
	newline
	."
	pop"
	peek
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
	
## Syntax 

Symbols (words) are whitespace delimited, therefore, even the `;` symbol must be separated from other runes by whitespace. 