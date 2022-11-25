package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConfig_AddDefaultData(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)

	ctx := getCtx(req)

	req = req.WithContext(ctx)

	testApp.Session.Put(ctx, "flash", "flash")
	testApp.Session.Put(ctx, "warning", "warning")
	testApp.Session.Put(ctx, "error", "error")

	td := testApp.AddDefaultData(&TemplateData{}, req)

	if td.Flash != "flash" {
		t.Errorf("failed to get flash data")
	}
	if td.Warning != "warning" {
		t.Errorf("failed to get warning data")
	}
	if td.Error != "error" {
		t.Errorf("failed to get error data")
	}
}

func TestConfig_IsAuthenticated(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)

	ctx := getCtx(req)

	req = req.WithContext(ctx)

	auth := testApp.IsAuthenticated(req)

	if auth {
		t.Errorf("return true for authenticated, when it should be false")
	}

	testApp.Session.Put(ctx, "userID", 1)

	auth = testApp.IsAuthenticated(req)

	if !auth {
		t.Errorf("return false for authenticated, when it should be true")
	}
}

func TestConfig_render(t *testing.T) {
	pathToTemplate = "./templates"

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	testApp.render(rr, req, "home.page.gohtml", &TemplateData{})

	if rr.Code != 200 {
		t.Errorf("failed to rnder page")
	}
}
