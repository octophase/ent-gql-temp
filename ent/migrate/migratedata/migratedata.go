package migratedata

import (
	"ariga.io/atlas/sql/migrate"
	//"ariga.io/atlas/sql/sqltool"
	"context"
	"myapp/ent"
	//"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	//"entgo.io/ent/schema"
	"fmt"
)

// SeedUsers add the initial users to the database.
func SeedUsers(dir *migrate.LocalDir) error {
	w := &schema.DirWriter{Dir: dir}
	client := ent.NewClient(ent.Driver(schema.NewWriteDriver(dialect.Postgres, w)))

	/*
		// The statement that generates the INSERT statement.
		err := client.User.CreateBulk(
			client.User.Create().SetName("foo"),
			client.User.Create().SetName("bar"),
		).Exec(context.Background())
		if err != nil {
			return fmt.Errorf("failed generating statement: %w", err)
		}
	*/
	_, err := client.User.
		Create().
		SetName("foo").
		Save(context.Background())
	if err != nil {
		return fmt.Errorf("failed generating statement: %w", err)
	}

	// Write the content to the migration directory.
	return w.FlushChange(
		"seed_users",
		"Add the initial users to the database.",
	)
}
