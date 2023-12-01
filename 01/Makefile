all: part2

part2.tab.c part2.tab.h:	part2.y
	bison -d part2.y

lex.yy.c: part2.l part2.tab.h
	flex --header-file=lex.yy.h part2.l

part2: lex.yy.c lex.yy.h part2.tab.c part2.tab.h parser.h
	$(CC) -o part2 part2.c part2.tab.c lex.yy.c -Wall -Werror

make clean:
	rm lex.yy.c lex.yy.h part2.tab.c part2.tab.h
