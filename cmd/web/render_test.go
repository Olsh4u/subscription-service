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
	testApp.Session.Put(ctx, "error", "error")
	testApp.Session.Put(ctx, "warning", "warning")

	td := testApp.AddDefaultData(&TemplateData{}, req)

	if td.Flash != "flash" {
		t.Errorf("Flash expected: %s, got: %s", "flash", td.Flash)
	}
	if td.Error != "error" {
		t.Errorf("Error expected: %s, got: %s", "error", td.Error)
	}
	if td.Warning != "warning" {
		t.Errorf("Warning expected: %s, got: %s", "warning", td.Warning)
	}
}

func TestConfig_IsAuthenticated(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	auth := testApp.IsAuthenticated(req)
	if auth {
		t.Error("returns true for authenticated, when it should be false")
	}

	testApp.Session.Put(ctx, "userID", 1)
	auth = testApp.IsAuthenticated(req)
	if !auth {
		t.Error("returns false for authenticated, when it should be true")
	}
}

func TestConfig_render(t *testing.T) {
	pathToTemplates = "./templates"

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	testApp.render(rr, req, "home.page.gohtml", &TemplateData{})

	if rr.Code != 200 {
		t.Errorf("Status code expected: %d, got: %d", 200, rr.Code)
	}
}
