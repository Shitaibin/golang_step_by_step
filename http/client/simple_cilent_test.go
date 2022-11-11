package client

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

type Language struct {
	name string
}

func TestSimple_Quest(t *testing.T) {
	port := "10234"
	endpoint := fmt.Sprintf("127.0.0.1:%v", port)
	java := Language{
		name: "java",
	}

	svr := newServer()
	svr.inject(java.name, java)

	s, err := NewClient(endpoint, http.DefaultTransport)
	if err != nil {
		t.Fatal(err)
	}

	req := &Request{
		LeftUrl:    java.name,
		HttpMethod: HTTP_GET,
	}
	got := &Language{}
	err = s.Quest(context.TODO(), req, got)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(java, got) {
		t.Fatalf("want: %v, got: %v", java, got)
	}
}
