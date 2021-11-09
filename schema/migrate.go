package schema

import (
	"database/sql"

	"github.com/GuiaBolso/darwin"
)

var migrations = []darwin.Migration{
	{
		Version:     1,
		Description: "Install uuid-ossp",
		Script: `
            CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
		`,
	},
	{
		Version:     2,
		Description: "Create package_features Table",
		Script: `
			CREATE TABLE package_features (
				id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
				name varchar(45) NOT NULL UNIQUE,
				created_at timestamp NOT NULL DEFAULT NOW(),
				updated_at timestamp NOT NULL DEFAULT NOW()
			);
		`,
	},
	{
		Version:     3,
		Description: "Create companies Table",
		Script: `
			CREATE TABLE companies (
				id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
				name varchar(45) NOT NULL,
				code char(4) NOT NULL UNIQUE,
				address varchar NOT NULL,
                city_id uuid NOT NULL,
				city varchar(45) NOT NULL,
                province_id uuid NOT NULL,
				province varchar(100) NOT NULL,
				npwp char(15) NULL,
				phone char(13) NOT NULL,
				pic varchar(45) NOT NULL,
				pic_phone varchar(20) NOT NULL,
				logo varchar NULL,
				package_feature_id uuid NOT NULL,
				created_at timestamp NOT NULL DEFAULT NOW(),
				updated_at timestamp NOT NULL DEFAULT NOW(),
				updated_by uuid NOT NULL,
                CONSTRAINT fk_companies_to_package_features FOREIGN KEY(package_feature_id) REFERENCES package_features(id)
			);
		`,
	},
	{
		Version:     4,
		Description: "Create regions Table",
		Script: `
			CREATE TABLE regions (
				id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
				company_id uuid NOT NULL,
				name varchar(45) NOT NULL,
				code char(4) NOT NULL,
				created_at timestamp NOT NULL DEFAULT NOW(),
				created_by uuid NOT NULL,
				updated_at timestamp NOT NULL DEFAULT NOW(),
				updated_by uuid NOT NULL,
				UNIQUE(company_id, code),
				CONSTRAINT fk_regions_to_companies FOREIGN KEY(company_id) REFERENCES companies(id)
			);
		`,
	},
	{
		Version:     5,
		Description: "Create branches Table",
		Script: `
			CREATE TABLE branches (
				id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
				company_id uuid NOT NULL,
				name varchar(45) NOT NULL,
				code char(4) NOT NULL,
				address varchar NOT NULL,
                city_id uuid NOT NULL,
				city varchar(45) NOT NULL,
                province_id uuid NOT NULL,
				province varchar(45) NOT NULL,
				npwp char(15) NULL,
				phone varchar(13) NOT NULL,
				pic varchar(45) NOT NULL,
				pic_phone varchar(13) NOT NULL,
				created_at timestamp NOT NULL DEFAULT NOW(),
				created_by uuid NOT NULL,
				updated_at timestamp NOT NULL DEFAULT NOW(),
				updated_by uuid NOT NULL,
				UNIQUE(company_id, code),
				CONSTRAINT fk_branches_to_companies FOREIGN KEY(company_id) REFERENCES companies(id)
			);
		`,
	},
	{
		Version:     6,
		Description: "Create users Table",
		Script: `
			CREATE TABLE users (
				id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
				company_id uuid NOT NULL,
				region_id uuid NULL,
				branch_id uuid NULL,
				username varchar(20) NOT NULL UNIQUE,
				password varchar NOT NULL,
				name varchar(45) NOT NULL,
				email varchar(100) NOT NULL UNIQUE,
				created_at timestamp NOT NULL DEFAULT NOW(),
				updated_at timestamp NOT NULL DEFAULT NOW(),
				created_by uuid NOT NULL,
                updated_by uuid NOT NULL,
				CONSTRAINT fk_users_to_companies FOREIGN KEY(company_id) REFERENCES companies(id),
				CONSTRAINT fk_users_to_regions FOREIGN KEY(region_id) REFERENCES regions(id),
				CONSTRAINT fk_users_to_branches FOREIGN KEY(branch_id) REFERENCES branches(id),
                CONSTRAINT fk_created_by_users FOREIGN KEY(created_by) REFERENCES users(id),
                CONSTRAINT fk_updated_by_users FOREIGN KEY(updated_by) REFERENCES users(id)
			);
		`,
	},
	{
		Version:     7,
		Description: "Alter companies Table on updated_by",
		Script: `
			ALTER TABLE companies 
            ADD CONSTRAINT fk_updated_by_companies_to_users FOREIGN KEY(updated_by) REFERENCES users(id);
		`,
	},
	{
		Version:     8,
		Description: "Alter regions Table",
		Script: `
			ALTER TABLE regions 
            ADD CONSTRAINT fk_created_by_regions_to_users FOREIGN KEY(created_by) REFERENCES users(id);
		`,
	},
	{
		Version:     9,
		Description: "Alter regions Table",
		Script: `
			ALTER TABLE regions 
            ADD CONSTRAINT fk_updated_by_regions_to_users FOREIGN KEY(updated_by) REFERENCES users(id);
		`,
	},
	{
		Version:     10,
		Description: "Alter branches Table",
		Script: `
			ALTER TABLE branches 
            ADD CONSTRAINT fk_created_by_branches_to_users FOREIGN KEY(created_by) REFERENCES users(id);
		`,
	},
	{
		Version:     11,
		Description: "Alter branches Table",
		Script: `
			ALTER TABLE branches 
            ADD CONSTRAINT fk_updated_by_branches_to_users FOREIGN KEY(updated_by) REFERENCES users(id);
		`,
	},
	{
		Version:     12,
		Description: "Create branches_regions Table",
		Script: `
			CREATE TABLE branches_regions (
				id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
				region_id uuid NOT NULL,
				branch_id uuid NOT NULL,
				created_at timestamp NOT NULL DEFAULT NOW(),
				created_by uuid NOT NULL,
				updated_at timestamp NOT NULL DEFAULT NOW(),
				updated_by uuid NOT NULL,
				UNIQUE(region_id, branch_id),
				CONSTRAINT fk_branches_regions_to_regions FOREIGN KEY(region_id) REFERENCES regions(id) ON DELETE CASCADE,
				CONSTRAINT fk_branches_regions_to_branches FOREIGN KEY(branch_id) REFERENCES branches(id) ON DELETE CASCADE,
                CONSTRAINT fk_created_by_branches_regions_to_users FOREIGN KEY(created_by) REFERENCES users(id),
                CONSTRAINT fk_updated_by_branches_regions_to_users FOREIGN KEY(updated_by) REFERENCES users(id)
			);
		`,
	},
	{
		Version:     13,
		Description: "Create roles Table",
		Script: `
			CREATE TABLE roles (
				id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
				company_id uuid NOT NULL,
				name varchar(45) NOT NULL,
				is_mutable bool NOT NULL DEFAULT false,
				created_at timestamp NOT NULL DEFAULT NOW(),
				created_by uuid NOT NULL,
				updated_at timestamp NOT NULL DEFAULT NOW(),
				updated_by uuid NOT NULL,
				CONSTRAINT fk_roles_to_companies FOREIGN KEY(company_id) REFERENCES companies(id),
                CONSTRAINT fk_created_by_roles_to_users FOREIGN KEY(created_by) REFERENCES users(id),
                CONSTRAINT fk_updated_by_roles_to_users FOREIGN KEY(updated_by) REFERENCES users(id)
			);
		`,
	},
	{
		Version:     14,
		Description: "Create access Table",
		Script: `
			CREATE TABLE access (
				id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
				parent_id uuid NULL,
				name varchar(45) NOT NULL UNIQUE,
				created_at timestamp NOT NULL DEFAULT NOW(),
				created_by uuid NOT NULL,
				updated_at timestamp NOT NULL DEFAULT NOW(),
				updated_by uuid NOT NULL,
				CONSTRAINT fk_access_to_parents FOREIGN KEY(parent_id) REFERENCES access(id),
                CONSTRAINT fk_created_by_access_to_users FOREIGN KEY(created_by) REFERENCES users(id),
                CONSTRAINT fk_updated_by_access_to_users FOREIGN KEY(updated_by) REFERENCES users(id) 
			);
		`,
	},
	{
		Version:     15,
		Description: "Create access_roles Table",
		Script: `
			CREATE TABLE access_roles (
				id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
				role_id uuid NOT NULL,
				access_id uuid NOT NULL,
				created_at timestamp NOT NULL DEFAULT NOW(),
				created_by uuid NOT NULL,
				updated_at timestamp NOT NULL DEFAULT NOW(),
				updated_by uuid NOT NULL,
				UNIQUE(role_id, access_id),
				CONSTRAINT fk_access_roles_to_roles FOREIGN KEY(role_id) REFERENCES roles(id) ON DELETE CASCADE,
				CONSTRAINT fk_access_roles_to_access FOREIGN KEY(access_id) REFERENCES access(id) ON DELETE CASCADE,
                CONSTRAINT fk_created_by_access_roles_to_users FOREIGN KEY(created_by) REFERENCES users(id),
                CONSTRAINT fk_updated_by_access_roles_to_users FOREIGN KEY(updated_by) REFERENCES users(id)
			);
		`,
	},
	{
		Version:     16,
		Description: "Create employees Table",
		Script: `
			CREATE TABLE employees (
				id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
				user_id uuid NOT NULL UNIQUE,
				name varchar(45) NOT NULL,
				code char(20) NOT NULL UNIQUE,
				address varchar NOT NULL,
                city_id uuid NOT NULL,
				city varchar(45) NOT NULL,
                province_id uuid NOT NULL,
				province varchar(45) NOT NULL,
				jabatan varchar(45) NOT NULL,
				created_at timestamp NOT NULL DEFAULT NOW(),
				created_by uuid NOT NULL,
				updated_at timestamp NOT NULL DEFAULT NOW(),
				updated_by uuid NOT NULL,
				CONSTRAINT fk_employees_to_users FOREIGN KEY(user_id) REFERENCES users(id),
                CONSTRAINT fk_created_by_employees_to_users FOREIGN KEY(created_by) REFERENCES users(id),
                CONSTRAINT fk_updated_by_employees_to_users FOREIGN KEY(updated_by) REFERENCES users(id)
			);
		`,
	},
	{
		Version:     17,
		Description: "Create features Table",
		Script: `
			CREATE TABLE features (
				id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
				name varchar(45) NOT NULL UNIQUE,
				created_at timestamp NOT NULL DEFAULT NOW(),
				updated_at timestamp NOT NULL DEFAULT NOW()
			);
		`,
	},
	{
		Version:     18,
		Description: "Create features_package_features Table",
		Script: `
			CREATE TABLE features_package_features (
				id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
				package_feature_id uuid NOT NULL,
				feature_id uuid NOT NULL,
				created_at timestamp NOT NULL DEFAULT NOW(),
				created_by uuid NOT NULL,
				updated_at timestamp NOT NULL DEFAULT NOW(),
				updated_by uuid NOT NULL,
				UNIQUE(package_feature_id, feature_id),
				CONSTRAINT fk_features_package_features_to_package_features FOREIGN KEY(package_feature_id) REFERENCES package_features(id) ON DELETE CASCADE,
				CONSTRAINT fk_features_package_features_to_features FOREIGN KEY(feature_id) REFERENCES features(id) ON DELETE CASCADE,
                CONSTRAINT fk_created_by_features_package_features_to_users FOREIGN KEY(created_by) REFERENCES users(id),
                CONSTRAINT fk_updated_by_features_package_features_to_users FOREIGN KEY(updated_by) REFERENCES users(id)
			);
		`,
	},
	{
		Version:     19,
		Description: "Create companies_features Table",
		Script: `
			CREATE TABLE companies_features (
				id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
				company_id uuid NOT NULL,
				feature_id uuid NOT NULL,
				created_at timestamp NOT NULL DEFAULT NOW(),
				created_by uuid NOT NULL,
				updated_at timestamp NOT NULL DEFAULT NOW(),
				updated_by uuid NOT NULL,
				UNIQUE(company_id, feature_id),
				CONSTRAINT fk_companies_features_to_companies FOREIGN KEY(company_id) REFERENCES companies(id) ON DELETE CASCADE,
				CONSTRAINT fk_companies_features_to_features FOREIGN KEY(feature_id) REFERENCES features(id) ON DELETE CASCADE,
                CONSTRAINT fk_created_by_companies_features_to_users FOREIGN KEY(created_by) REFERENCES users(id),
                CONSTRAINT fk_updated_by_companies_features_to_users FOREIGN KEY(updated_by) REFERENCES users(id)
			);
		`,
	},
	{
		Version:     20,
		Description: "Create request_passwords Table",
		Script: `
			CREATE TABLE request_passwords (
				id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
				user_id uuid NOT NULL,
				is_used boolean NOT NULL DEFAULT false,
				created_at timestamp NOT NULL DEFAULT NOW(),
				CONSTRAINT fk_request_passwords_to_users FOREIGN KEY(user_id) REFERENCES users(id)
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
