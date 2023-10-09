package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFromJsonReq(t *testing.T) {
	type args struct {
		endpoint string
		method   string
		payload  string
		headers  []Header
		field    string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr error
	}{
		{
			name: "basic test",
			args: args{
				endpoint: "/mirror",
				method:   "POST",
				payload:  `{"test": "test"}`,
				headers: []Header{
					{
						Key:   "Content-Type",
						Value: "application/json",
					},
				},
				field: "test",
			},
			want:    "test",
			wantErr: nil,
		},
	}

	mockApi := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/mirror":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"test": "test"}`))

		case "whole-struct":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"foo": "bar", "hello": "world"}`))
		}
	}))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFromJsonReq[string](mockApi.URL+tt.args.endpoint, Method(tt.args.method), tt.args.payload, tt.args.headers, tt.args.field)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("GetFromJsonReq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFromJsonReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFromJsonReqEmptyField(t *testing.T) {
	type testStruct struct {
		Foo string `json:"foo"`
		Baz int    `json:"baz"`
	}

	type args struct {
		endpoint string
		method   string
		payload  string
		headers  []Header
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr error
	}{
		{
			name: "whole struct test",
			args: args{
				endpoint: "/whole/struct",
				method:   "GET",
				payload:  "",
				headers: []Header{
					{
						Key:   "Content-Type",
						Value: "application/json",
					},
				},
			},
			want: testStruct{
				Foo: "bar",
				Baz: 1,
			},
			wantErr: nil,
		},
	}

	mockApi := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/mirror":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"test": "test"}`))

		case "/whole/struct":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`[{"foo": "bar", "baz": 1}]`))
		}
	}))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFromJsonReq[[]testStruct](mockApi.URL+tt.args.endpoint, Method(tt.args.method), tt.args.payload, tt.args.headers, "")
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("GetFromJsonReq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got[0] != tt.want {
				t.Errorf("GetFromJsonReq() = %v, want %v", got, tt.want)
			}
		})
	}
}
