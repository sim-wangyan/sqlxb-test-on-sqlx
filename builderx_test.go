package main

import (
	"fmt"
	. "github.com/x-ream/sqlxb"
)

func testBulderX() {

	pet := Pet{}
	cat := Cat{}
	dog := Dog{}

	//"COUNT(DISTINCT d.id) AS `d.id_count`"

	builder := NewBuilderX(&cat,"c")
	builder.ResultKeys( "distinct c.color","COUNT(DISTINCT d.id) AS `d.id_count`")//"COUNT(DISTINCT d.id) AS `d.id_count`"
	builder.Eq("p.id", 1)

	subP := Sub()
	subP.ResultKeys("id").Source(&pet)
	builder.SourceBuilder().Sub(subP).Alia("p").JoinOn(LEFT_JOIN,ON("id","c","pet_id"))

	arr := []interface{}{3000,4000,5000,6000}
	sub := Sub()
	sub.ResultKeys("pet_id").Source(&dog).Eq("age",2).In("weight",arr...)
	builder.SourceBuilder().Sub(sub).Alia("d").JoinOn(LEFT_JOIN,ON("id","c","pet_id"))

	builder.SourceBuilder().Source(&cat).Alia("cat").JoinOn(INNER_JOIN,ON("pet_id","p","id"))
	builder.
		GroupBy("c.color").
		Having(Gt,"d.id_count",1000).
		Sort("p.id",DESC).
		Paged().Rows(10).Last(101)

	vs, dataSql, countSql, kmp:= builder.WithoutOptimization().Build().Sql()
	fmt.Println(dataSql)
	fmt.Println(vs)
	fmt.Println(kmp)
	fmt.Println(countSql)

	// start sqlx
	InitSqlxDB()

	catList := []Cat{}
	err := Db.Select(&catList, dataSql,vs...)
	if err != nil {
		fmt.Println(err)
	}
	s := fmt.Sprintf("price : %v", *(catList[0].Price))
	fmt.Println(s)
}

