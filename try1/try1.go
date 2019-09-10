package main

import (
	"github.com/pingcap/parser"
	_ "github.com/pingcap/tidb/types/parser_driver"
)

func main() {
	p := parser.New()

	_, err := p.ParseOneStmt("CREATE TABLE `c` ( `id` int(11) NOT NULL, PRIMARY KEY (`id`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_as_cs", "", "")

	panicIfErr(err)
}

func panicIfErr(err interface{}) {
	if err != nil {
		panic(err)
	}
}
