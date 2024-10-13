package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/vyctor/go-finance-backend/utils"
	"testing"
)

func createRandomCategory(t *testing.T) Category {
	user := createRandomUser(t)

	arg := CreateCategoryParams{
		UserID:      user.ID,
		Title:       utils.RandomString(10),
		Type:        "debit",
		Description: utils.RandomString(50),
	}

	categoryId, err := testQueries.CreateCategory(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, categoryId)

	createdCategory, err := testQueries.GetCategoryById(context.Background(), categoryId)

	require.NoError(t, err)
	require.NotEmpty(t, createdCategory)
	require.Equal(t, arg.UserID, createdCategory.UserID)
	require.Equal(t, arg.Title, createdCategory.Title)
	require.Equal(t, arg.Type, createdCategory.Type)
	require.Equal(t, arg.Description, createdCategory.Description)
	require.NotEmpty(t, createdCategory.CreatedAt)
	require.NotEmpty(t, createdCategory.UpdatedAt)

	return createdCategory
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	category := createRandomCategory(t)
	require.NotEmpty(t, category)
	createdCategory, err := testQueries.GetCategoryById(context.Background(), category.ID)
	require.NoError(t, err)
	require.NotEmpty(t, createdCategory)
	require.Equal(t, category.UserID, createdCategory.UserID)
	require.Equal(t, category.Title, createdCategory.Title)
	require.Equal(t, category.Type, createdCategory.Type)
	require.Equal(t, category.Description, createdCategory.Description)
	require.NotEmpty(t, createdCategory.CreatedAt)
	require.NotEmpty(t, createdCategory.UpdatedAt)
}

func TestDeleteCategory(t *testing.T) {
	category := createRandomCategory(t)
	require.NotEmpty(t, category)
	err := testQueries.DeleteCategory(context.Background(), category.ID)
	require.NoError(t, err)
	deletedCategory, err := testQueries.GetCategoryById(context.Background(), category.ID)
	require.Error(t, err)
	require.Empty(t, deletedCategory)
}

func TestUpdateCategory(t *testing.T) {
	category := createRandomCategory(t)
	require.NotEmpty(t, category)
	updateCategoryArgs := UpdateCategoryParams{
		ID:          category.ID,
		Title:       utils.RandomString(10),
		Type:        "credit",
		Description: utils.RandomString(50),
	}
	updatedCategoryId, err := testQueries.UpdateCategory(context.Background(), updateCategoryArgs)
	require.NoError(t, err)
	require.NotEmpty(t, updatedCategoryId)
	require.Equal(t, category.ID, updatedCategoryId)
	updatedCategory, err := testQueries.GetCategoryById(context.Background(), category.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedCategoryId)
	require.Equal(t, category.UserID, updatedCategory.UserID)
	require.Equal(t, updateCategoryArgs.Title, updatedCategory.Title)
	require.Equal(t, updateCategoryArgs.Type, updatedCategory.Type)
	require.Equal(t, updateCategoryArgs.Description, updatedCategory.Description)
	require.NotEqual(t, category.Title, updatedCategory.Title)
	require.NotEqual(t, category.Type, updatedCategory.Type)
	require.NotEqual(t, category.Description, updatedCategory.Description)
	require.NotEqual(t, category.UpdatedAt, updatedCategory.UpdatedAt)
}

func TestListCategories(t *testing.T) {
	var lastCategory Category
	for i := 0; i < 10; i++ {
		lastCategory = createRandomCategory(t)
	}

	arg := GetCategoriesParams{
		UserID:      lastCategory.UserID,
		Type:        lastCategory.Type,
		Title:       lastCategory.Title,
		Description: lastCategory.Description,
	}

	categories, err := testQueries.GetCategories(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, categories)

	for _, category := range categories {
		require.Equal(t, lastCategory.ID, category.ID)
		require.Equal(t, lastCategory.UserID, category.UserID)
		require.Equal(t, lastCategory.Title, category.Title)
		require.Equal(t, lastCategory.Type, category.Type)
		require.Equal(t, lastCategory.Description, category.Description)
		require.NotEmpty(t, category.CreatedAt)
		require.NotEmpty(t, category.UpdatedAt)
	}
}
