package main

import (
	"context"
	"database/sql"
	"fmt"
)

func run(ctx context.Context, db *sql.DB) error {
	q := &Queries{db: db} // defined in the next section

	insertedAuthor, err := q.CreateAuthor(ctx, CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio: sql.NullString{
			String: "Co-author of The C Programming Language",
			Valid:  true,
		},
	})
	if err != nil {
		return err
	}

	authors, err := q.ListAuthors(ctx)
	if err != nil {
		return err
	}
	fmt.Println(authors)

	err = q.DeleteAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// TODO: Open connection to your PostgreSQL database
	//assuming you are going name your database instance db like so: db, err := sql.Open("postgres", pgConString)
	run(context.Background(), db)
}
