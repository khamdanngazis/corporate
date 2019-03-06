package services

import (
	"context"

	dt "../datastruct"
	lib "../lib"
	"../logging"
)

// CorporateCategoryListServices provides operations for endpoint
type CorporateCategoryListServices interface {
	CorporateCategoryList(context.Context, dt.CorporateCategoryListJSONRequest, lib.DbConnection) dt.CorporateCategoryListJSONResponse
}

// CorporateCategoryListService is a concrete implementation of QueueServices
type CorporateCategoryListService struct{}

//Get Category list
func GetCategoryList(dbConn lib.DbConnection, level int, key1 string, key2 string, key3 string) ([]dt.CategoryList, error) {
	var e error
	var categoryList dt.CategoryList
	var aCategoryList []dt.CategoryList
	var key string
	var value string
	var haschild string
	if level == 1 {
		var retVal, e = dbConn.Query("GetCategoryListLevel1")
		if e != nil {
			logging.Logf("query GetCategoryListLevel1 error : %s", e)
			return nil, e
		}
		defer retVal.Close()
		for retVal.Next() {
			if e = retVal.Scan(&key, &value, &haschild); e != nil {
				logging.Logf("scan GetCategoryListLevel1 error : %s", e)
				return nil, e
			}
			categoryList.Key = key
			categoryList.Value = value
			if haschild != "" {
				categoryList.HasChild = true
			} else {
				categoryList.HasChild = false
			}
			aCategoryList = append(aCategoryList, categoryList)
		}

	}
	if level == 2 {
		var retVal, e = dbConn.Query("GetCategoryListLevel2", key1)
		if e != nil {
			logging.Logf("query GetCategoryListLevel2 error: %s", e)
			return nil, e
		}
		defer retVal.Close()
		for retVal.Next() {
			if e = retVal.Scan(&key, &value, &haschild); e != nil {
				logging.Logf("scan GetCategoryListLevel2 error : %s", e)
				return nil, e
			}
			categoryList.Key = key
			categoryList.Value = value
			if haschild != "" {
				categoryList.HasChild = true
			} else {
				categoryList.HasChild = false
			}
			aCategoryList = append(aCategoryList, categoryList)
		}
	}
	if level == 3 {
		var retVal, e = dbConn.Query("GetCategoryListLevel3", key1, key2)
		if e != nil {
			logging.Logf("query GetCategoryListLevel3 error : %s", e)
			return nil, e
		}
		defer retVal.Close()
		for retVal.Next() {
			if e = retVal.Scan(&key, &value, &haschild); e != nil {
				logging.Logf("scan GetCategoryListLevel3 error : %s", e)
				return nil, e
			}

			categoryList.Key = key
			categoryList.Value = value
			if haschild != "" {
				categoryList.HasChild = true
			} else {
				categoryList.HasChild = false
			}
			aCategoryList = append(aCategoryList, categoryList)
		}
	}
	if level == 4 {
		var retVal, e = dbConn.Query("GetCategoryListLevel4", key1, key2, key3)
		if e != nil {
			logging.Logf("query GetCategoryListLevel4 error : %s", e)
			return nil, e
		}
		defer retVal.Close()
		for retVal.Next() {
			if e = retVal.Scan(&key, &value, &haschild); e != nil {
				logging.Logf("scan GetCategoryListLevel4 error : %s", e)
				return nil, e
			}
			categoryList.Key = key
			categoryList.Value = value
			if haschild != "" {
				categoryList.HasChild = true
			} else {
				categoryList.HasChild = false
			}

			aCategoryList = append(aCategoryList, categoryList)
		}
	}
	return aCategoryList, e
}

// CorporateCategoryList service
func (CorporateCategoryListService) CorporateCategoryList(ctx context.Context, req dt.CorporateCategoryListJSONRequest, dbConn lib.DbConnection) dt.CorporateCategoryListJSONResponse {

	CategoryList, qlErr := GetCategoryList(dbConn, req.Level, req.KeyLevel1, req.KeyLevel2, req.KeyLevel3)
	if qlErr != nil {
		logging.Logf("%s Failed to get Category list", qlErr)
		return dt.CorporateCategoryListJSONResponse{
			ResponseCode: dt.ErrGetCategoryListFailed,
			ResponseDesc: dt.DescGetCategoryListFailed,
		}
	}
	return dt.CorporateCategoryListJSONResponse{
		CategoryList: CategoryList,
		ResponseCode: dt.ErrSuccess,
		ResponseDesc: dt.DescSuccess,
	}

}
