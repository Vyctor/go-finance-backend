// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateBill(ctx context.Context, arg CreateBillParams) (int32, error)
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (int32, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (int32, error)
	DeleteBill(ctx context.Context, id int32) error
	DeleteCategory(ctx context.Context, id int32) error
	GetBill(ctx context.Context, id int32) (Bill, error)
	GetBills(ctx context.Context, arg GetBillsParams) ([]GetBillsRow, error)
	GetBillsGraph(ctx context.Context, arg GetBillsGraphParams) (int64, error)
	GetBillsReports(ctx context.Context, arg GetBillsReportsParams) (int64, error)
	GetCategories(ctx context.Context, arg GetCategoriesParams) ([]Category, error)
	GetCategoryById(ctx context.Context, id int32) (Category, error)
	GetCategoryByUserId(ctx context.Context, userID int32) (Category, error)
	GetUserById(ctx context.Context, id int32) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	UpdateBill(ctx context.Context, arg UpdateBillParams) (int32, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (int32, error)
}

var _ Querier = (*Queries)(nil)
