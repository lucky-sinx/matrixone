package main

import (
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	log "github.com/sirupsen/logrus"
	_ "matrixone/pkg/types/parser_driver"
)

func parse(sql string) ([]ast.StmtNode, error) {
	p := parser.New()

	stmts, _, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}

	return stmts, nil
}

func doVisit(pos int, stmt ast.Node) {
	c := &checker{Pos: pos}
	stmt.Accept(c)
}

func main() {
	stmts, err := parse("INSERT INTO t VALUES(1, 2); SELECT a, b FROM t;")
	if err != nil {
		log.Errorf("parse error: %v\n", err.Error())
		return
	}

	for i, stmt := range stmts {
		log.Infof("[%d] Raw Statement: %v", i, stmt.Text())
		doVisit(i, stmt)
	}
}