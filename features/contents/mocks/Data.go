// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	contents "antonio/features/contents"

	mock "github.com/stretchr/testify/mock"
)

// Data is an autogenerated mock type for the Data type
type Data struct {
	mock.Mock
}

// CreateArticle provides a mock function with given fields: data, userId, tags
func (_m *Data) CreateArticle(data contents.Core, userId int, tags []contents.ToolCore) error {
	ret := _m.Called(data, userId, tags)

	var r0 error
	if rf, ok := ret.Get(0).(func(contents.Core, int, []contents.ToolCore) error); ok {
		r0 = rf(data, userId, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTags provides a mock function with given fields: tags
func (_m *Data) CreateTags(tags []contents.ToolCore) ([]contents.ToolCore, error) {
	ret := _m.Called(tags)

	var r0 []contents.ToolCore
	if rf, ok := ret.Get(0).(func([]contents.ToolCore) []contents.ToolCore); ok {
		r0 = rf(tags)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]contents.ToolCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]contents.ToolCore) error); ok {
		r1 = rf(tags)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteArticleById provides a mock function with given fields: articleId
func (_m *Data) DeleteArticleById(articleId int) error {
	ret := _m.Called(articleId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(articleId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllArticles provides a mock function with given fields:
func (_m *Data) GetAllArticles() ([]contents.Core, error) {
	ret := _m.Called()

	var r0 []contents.Core
	if rf, ok := ret.Get(0).(func() []contents.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]contents.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUserArticles provides a mock function with given fields: userId
func (_m *Data) GetAllUserArticles(userId int) ([]contents.Core, error) {
	ret := _m.Called(userId)

	var r0 []contents.Core
	if rf, ok := ret.Get(0).(func(int) []contents.Core); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]contents.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArticleById provides a mock function with given fields: articleId
func (_m *Data) GetArticleById(articleId int) (contents.Core, error) {
	ret := _m.Called(articleId)

	var r0 contents.Core
	if rf, ok := ret.Get(0).(func(int) contents.Core); ok {
		r0 = rf(articleId)
	} else {
		r0 = ret.Get(0).(contents.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(articleId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateArticleById provides a mock function with given fields: articleId, data
func (_m *Data) UpdateArticleById(articleId int, data contents.Core) error {
	ret := _m.Called(articleId, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, contents.Core) error); ok {
		r0 = rf(articleId, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyArticleOwner provides a mock function with given fields: articleId, userId
func (_m *Data) VerifyArticleOwner(articleId int, userId int) error {
	ret := _m.Called(articleId, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(articleId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
