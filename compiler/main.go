package main

import (
	"fmt"
	"vnh1/ast"
	"vnh1/lexer"
)

func main() {
	_ = `
#auth prammer {
	rblockcall ("server uri", {}) &(userPub:=userPub, host:=store.host, agent:=store.agent) {

	}
	catch(error) {

	}
}

#streamrauth "ssigner://localhost:7200" {
	#sigverify prammer;
}
`
	input := `
rblockcall ("server uri", {}) &(userPub:=userPub, host:=store.host, agent:=store.agent) {

}
catch(error) {

}
`

	// Initialisiere den Lexer mit dem Eingabetext
	lexer := lexer.NewLexer(input)

	// Das Script wird Gelext
	tokenlist := lexer.LexTokenList()

	// Das Parsen des Tokens wird vorbereitet
	parser := ast.NewParser(tokenlist)

	// Das Script wird geparst
	fmt.Println(parser.ParseProgram())
}