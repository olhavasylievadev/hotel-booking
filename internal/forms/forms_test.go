package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("Got invalid while has to be valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.Form)

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("form shows valid while required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData

	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows that doesn't have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("whatever")

	if has {
		t.Error("form shows has field when it doesn't")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("shows that form doesn't have the request, while has it")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min-length for non-existent field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("expected an error, but didn't get one")
	}

	postedValue := url.Values{}
	postedValue.Add("some_field", "some value")
	form = New(postedValue)

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows min-length of 100 met when the data is shorter")
	}

	postedValue = url.Values{}
	postedValue.Add("another_field", "abc123")
	form = New(postedValue)
	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("Min-length should be valid, but it's not")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("should not have an error, but got one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedValue := url.Values{}
	form := New(postedValue)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for a non-existent field")
	}

	postedValue = url.Values{}
	postedValue.Add("email", "123@a.com")
	form = New(postedValue)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("form shows that email isn't valid, while it is")
	}

	postedValue = url.Values{}
	postedValue.Add("email", "123")
	form = New(postedValue)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("form shows that email isn valid, while it isn't")
	}
}
