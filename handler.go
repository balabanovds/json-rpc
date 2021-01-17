package jsonrpc

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/balabanovds/jsonrpc/internal/parsers"
	"github.com/balabanovds/jsonrpc/internal/parsers/iter"
	"github.com/balabanovds/jsonrpc/internal/parsers/std"
	"github.com/balabanovds/jsonrpc/pkg/rpcerror"
	jsoniter "github.com/json-iterator/go"
)

type JSONParserType int

const (
	JSONStd JSONParserType = iota
	JSONIter
)

type Method func(ctx context.Context, message *json.RawMessage) (interface{}, *rpcerror.Error)

type Handler struct {
	*processor
	json parsers.JSONParser
}

func NewHandler(methods map[string]Method, jsonType JSONParserType) *Handler {
	var json parsers.JSONParser

	switch jsonType {
	case JSONStd:
		json = std.New()
	case JSONIter:
		json = iter.New()
	}

	return &Handler{
		processor: newProcessor(methods, json),
		json:      json,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.serverError(w, rpcerror.NewServerError("failed to read body"))
		return
	}

	if len(body) == 0 {
		h.clientError(w, http.StatusNotAcceptable, "body length zero")
		return
	}

	resp := h.Call(r.Context(), body)
	h.respond(w, resp)
}

func (h *Handler) respond(w http.ResponseWriter, data interface{}) {
	if data != nil {
		if err := jsoniter.NewEncoder(w).Encode(data); err != nil {
			h.serverError(w, rpcerror.NewServerError("failed to encode data"))
		}
		return
	}

	h.serverError(w, rpcerror.NewServerError("incoming data is empty in respond"))
}

func (h *Handler) serverError(w http.ResponseWriter, err *rpcerror.Error) {
	_ = jsoniter.NewEncoder(w).Encode(err)
}

func (h *Handler) clientError(w http.ResponseWriter, code int, message string) {
	h.respond(w, rpcerror.NewError(code, message))
}
