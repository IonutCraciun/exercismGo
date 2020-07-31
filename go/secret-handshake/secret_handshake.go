package secret

/*
1 = wink
10 = double blink
100 = close your eyes
1000 = jump
10000 = Reverse the order of the operations in the secret handshake.

*/

type move struct{
	i int
	mv string
}

var moves = []move{
	{1, "wink"},
	{10, "double blink"},
	{100, "close your eyes"},
	{1000, "jump"},
}


// Handshake func
func Handshake(input int) (result []string) {
	
	for key, val := range moves {
		
	}
	return result
}