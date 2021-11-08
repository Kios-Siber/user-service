package repositories

import (
	"context"
	"fmt"
	"skeleton/lib/helper"
	"skeleton/pb/drivers"
	"skeleton/pb/generic"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *repo) FindAll(ctx context.Context, in *drivers.DriverListInput) (*drivers.Drivers, error) {
	select {
	case <-ctx.Done():
		return nil, helper.ContextError(ctx)
	default:
	}

	out := &drivers.Drivers{}
	query := `SELECT id, name, phone, licence_number, company_id, company_name FROM drivers`
	where := []string{"is_deleted = false"}
	paramQueries := []interface{}{}

	if len(in.Ids) > 0 {
		orWhere := []string{}
		for _, id := range in.Ids {
			paramQueries = append(paramQueries, id)
			orWhere = append(orWhere, fmt.Sprintf("id = %d", len(paramQueries)))
		}
		if len(orWhere) > 0 {
			where = append(where, "("+strings.Join(orWhere, " OR ")+")")
		}
	}

	if len(in.CompanyIds) > 0 {
		orWhere := []string{}
		for _, id := range in.CompanyIds {
			paramQueries = append(paramQueries, id)
			orWhere = append(orWhere, fmt.Sprintf("company_id = %d", len(paramQueries)))
		}
		if len(orWhere) > 0 {
			where = append(where, "("+strings.Join(orWhere, " OR ")+")")
		}
	}

	if len(in.LicenceNumbers) > 0 {
		orWhere := []string{}
		for _, licenceNumber := range in.LicenceNumbers {
			paramQueries = append(paramQueries, licenceNumber)
			orWhere = append(orWhere, fmt.Sprintf("licence_number = %d", len(paramQueries)))
		}
		if len(orWhere) > 0 {
			where = append(where, "("+strings.Join(orWhere, " OR ")+")")
		}
	}

	if len(in.Names) > 0 {
		orWhere := []string{}
		for _, name := range in.Names {
			paramQueries = append(paramQueries, name)
			orWhere = append(orWhere, fmt.Sprintf("name = %d", len(paramQueries)))
		}
		if len(orWhere) > 0 {
			where = append(where, "("+strings.Join(orWhere, " OR ")+")")
		}
	}

	if len(in.Phones) > 0 {
		orWhere := []string{}
		for _, phone := range in.Phones {
			paramQueries = append(paramQueries, phone)
			orWhere = append(orWhere, fmt.Sprintf("phone = %d", len(paramQueries)))
		}
		if len(orWhere) > 0 {
			where = append(where, "("+strings.Join(orWhere, " OR ")+")")
		}
	}

	if in.Pagination == nil {
		in.Pagination = &generic.Pagination{}
	}

	if len(in.Pagination.Keyword) > 0 {
		orWhere := []string{}

		paramQueries = append(paramQueries, in.Pagination.Keyword)
		orWhere = append(orWhere, fmt.Sprintf("name = %d", len(paramQueries)))

		paramQueries = append(paramQueries, in.Pagination.Keyword)
		orWhere = append(orWhere, fmt.Sprintf("phone = %d", len(paramQueries)))

		paramQueries = append(paramQueries, in.Pagination.Keyword)
		orWhere = append(orWhere, fmt.Sprintf("licence_number = %d", len(paramQueries)))

		paramQueries = append(paramQueries, in.Pagination.Keyword)
		orWhere = append(orWhere, fmt.Sprintf("company_name = %d", len(paramQueries)))

		if len(orWhere) > 0 {
			where = append(where, "("+strings.Join(orWhere, " OR ")+")")
		}
	}

	if len(in.Pagination.Sort) > 0 {
		in.Pagination.Sort = strings.ToLower(in.Pagination.Sort)
		if in.Pagination.Sort != "asc" {
			in.Pagination.Sort = "desc"
		}
	} else {
		in.Pagination.Sort = "desc"
	}

	if len(in.Pagination.Order) > 0 {
		in.Pagination.Order = strings.ToLower(in.Pagination.Order)
		if !(in.Pagination.Order == "id" ||
			in.Pagination.Order == "name" ||
			in.Pagination.Order == "phone" ||
			in.Pagination.Order == "licence_number" ||
			in.Pagination.Order == "company_id" ||
			in.Pagination.Order == "company_name") {
			in.Pagination.Order = "id"
		}
	} else {
		in.Pagination.Order = "id"
	}

	if in.Pagination.Limit <= 0 {
		in.Pagination.Limit = 10
	}

	if in.Pagination.Offset <= 0 {
		in.Pagination.Offset = 0
	}

	if len(where) > 0 {
		query += " WHERE " + strings.Join(where, " AND ")
	}

	query += " ORDER BY " + in.Pagination.Order + " " + in.Pagination.Sort
	query += " LIMIT " + strconv.Itoa(int(in.Pagination.Limit))
	query += " OFFSET " + strconv.Itoa(int(in.Pagination.Offset))

	rows, err := u.db.QueryContext(ctx, query, paramQueries...)
	if err != nil {
		u.log.Println(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var obj drivers.Driver
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Phone, &obj.LicenceNumber, &obj.CompanyId, &obj.CompanyName)
		if err != nil {
			u.log.Println(err.Error())
			return nil, status.Error(codes.Internal, err.Error())
		}

		out.Driver = append(out.Driver, &obj)
	}

	if rows.Err() != nil {
		u.log.Println(rows.Err().Error())
		return nil, status.Error(codes.Internal, rows.Err().Error())
	}

	return out, nil
}
