# part of the makefile
fb3: fb1-1.l basic.y
	rm -rf purecalc.tab.c
	rm -rf purecalc.yy.c
	bison -d purecalc.y
	flex purecalc.l
	gcc -o $@ purecalc.tab.c lex.yy.c  purecalcfuncs.c -ll

fb1:
	flex fb1-1.l
	gcc -o m$@ fb.c lex.yy.c -ll

#fb2:
#	bison -d basic.y
#	gcc -o m$@ basic.tab.c -ll
