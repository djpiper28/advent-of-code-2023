all: part1 part2

part1.tab.c part1.tab.h:	part1.y
	bison -d part1.y

part2.tab.c part2.tab.h:	part2.y
	bison -d part2.y

lex.yy.c: part1.l part1.tab.h
	flex --header-file=lex.yy.h part1.l

part1: lex.yy.c lex.yy.h part1.tab.c part1.tab.h parser.h
	$(CC) -o part1 main.c part1.tab.c lex.yy.c -Wall -Werror

part2: lex.yy.c lex.yy.h part2.tab.c part2.tab.h parser.h
	$(CC) -o part2 main.c part2.tab.c lex.yy.c -Wall -Werror

make clean:
	rm lex.yy.c lex.yy.h part1.tab.c part1.tab.h part2.tab.c part2.tab.h
