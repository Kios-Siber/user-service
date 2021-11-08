package schema

import (
	"database/sql"

	"github.com/GuiaBolso/darwin"
)

var migrations = []darwin.Migration{
	{
		Version:     1,
		Description: "Create drivers Table",
		Script: `
            CREATE TABLE public.drivers (
                id uuid NOT NULL,
                name varchar NOT NULL,
                phone varchar NOT NULL,
                licence_number varchar NOT NULL,
                company_id varchar NOT NULL,
                company_name varchar NOT NULL,
                is_deleted bool NOT NULL DEFAULT false,
                created timestamp(0) NOT NULL,
                created_by varchar NOT NULL,
                updated timestamp(0) NOT NULL,
                updated_by varchar NOT NULL,
                CONSTRAINT drivers_pk PRIMARY KEY (id)
            );
            CREATE UNIQUE INDEX drivers_phone ON public.drivers USING btree (phone);
        `,
	},
	{
		Version:     2,
		Description: "Create users Table",
		Script: `
            CREATE TABLE public.users (
                id uuid NOT NULL,
                name varchar NOT NULL,
                CONSTRAINT users_pk PRIMARY KEY (id)
            );
        `,
	},
}

// Migrate attempts to bring the schema for db up to date with the migrations
// defined in this package.
func Migrate(db *sql.DB) error {
	driver := darwin.NewGenericDriver(db, darwin.PostgresDialect{})

	d := darwin.New(driver, migrations, nil)

	return d.Migrate()
}
