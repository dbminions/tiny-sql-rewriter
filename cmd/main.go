package main

import (
	"fmt"
	"strings"
	advisor "tiny_rewriter/pkg/a_advisor"
	rewrite "tiny_rewriter/pkg/b_rewriter"
	env "tiny_rewriter/pkg/d_catalog"
)

func main() {
	sql := "select * from tbl t1 where id < 1000"

	// 1. suggest optimizations
	heuristicSuggest := suggestOptimizations(sql)
	fmt.Println(heuristicSuggest)
	/*
		map[
			ALI.001: {ALI.001 L0 It is recommended to use the AS keyword ...}
			OK: {OK L0 OK OK OK 0 0x10327f920}
		]
		NOTE:
		1. ALI.001 is added by the advisor rule "RuleImplicitAlias".
	*/

	// 2. rewrite sql
	newSql := rewriteSql(sql)
	fmt.Println(newSql)
	/*
		select id, name from tbl as t1 where id < 1000

		NOTE:
		1. tbl as t1 is done by `sqlparser.String(rw.Stmt)`. It is not part of the rewrite rules.
		2. id, name is added by the rewrite rule "RewriteStar2Columns".
	*/
}

func suggestOptimizations(sql string) map[string]advisor.Rule {
	q, _ := advisor.NewQuery4Audit(sql)
	return q.Advise()
}

func rewriteSql(sql string) string {
	rw := rewrite.NewRewrite(sql)

	// TODO: populate columns. For now, we use mock data.
	vEnv, _ := env.BuildEnv() //Environment initialization, connection check online environment + build test environment
	//rw.Columns = vEnv.GenTableColumns(rewrite.GetMeta(rw.Stmt, nil)) // Get the table structure of the test environment
	rw.Columns = vEnv.GenTableColumnsMock(rewrite.GetMeta(rw.Stmt, nil))

	rw.Rewrite()
	return strings.TrimSpace(rw.NewSQL)
}
