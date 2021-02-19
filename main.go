package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Post struct definition
type Post struct{
	ID int
	Title string
}

func main()  {
	db, err := sql.Open("mysql", "root:Nomzano100%@tcp(127.0.0.1:3306)/golang_mysql")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// insert, err := db.Query("INSERT INTO posts(title) VALUES('My Post')")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer insert.Close()

	 posts, err := db.Query("SELECT id, title FROm posts")

	 if err != nil {
		panic(err.Error())
	 }

	 var post Post
	 for posts.Next(){
		
		 err := posts.Scan(&post.ID, &post.Title)
		 if err != nil {
			panic(err.Error())
		 }
	 }

	 fmt.Println(post)
}