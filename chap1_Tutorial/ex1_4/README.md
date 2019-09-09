go run dup2_revised.go  
s b s  
s c s  
s b b  
s c s  
s b s  
s c s	 occurs inside file: 	keyboard input  
s b s	 occurs inside file: 	keyboard input  
  
[note Ctrl+D is EOF for stdin]  
  
go run dup2_revised.go a.txt b.txt  
s c b	 occurs inside file: 	a.txt  
a b c	 occurs inside file: 	a.txt	b.txt  


