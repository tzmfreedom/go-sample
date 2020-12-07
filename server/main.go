package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	headerContentLength  = "Content-Length"
	headerGRPCMessage    = "Grpc-Message"
	headerGRPCStatusCode = "Grpc-Status"

	contentTypeGRPCJSON = "application/grpc+json"
	grpcNoCompression byte = 0x00
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		req.URL.Scheme = "https"
		req.URL.Host = req.Header.Get("endpoint")
		r := modifyRequestToJSONgRPC(req)
		buf, err := request(r)
		if err != nil {
			handleError(w, err)
			return
		}
		w.Write(buf)
		w.WriteHeader(http.StatusOK)
	})
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}

func request(req *http.Request) ([]byte, error) {
	req, err := http.NewRequest("GET", req.URL.String(), nil)
	if err != nil {
		return nil, err
	}
	for k, v := range req.Header {
		req.Header.Set(k, v[0])
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	r, err := handleGRPCResponse(res)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(r.Body)
}

func handleError(w http.ResponseWriter, err error) {
	buf, _ := json.Marshal(map[string]string{
		"error": err.Error(),
	})
	w.Write(buf)
	w.WriteHeader(http.StatusBadRequest)
}

func modifyRequestToJSONgRPC(r *http.Request) *http.Request {
	// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md

	var body []byte
	// read body so we can add the grpc prefix
	if r.Body != nil {
		body, _ = ioutil.ReadAll(r.Body)
	}

	b := make([]byte, 0, len(body)+5)
	buff := bytes.NewBuffer(b)

	// grpc prefix is
	// 1 byte: compression indicator
	// 4 bytes: content length (excluding prefix)
	_ = buff.WriteByte(grpcNoCompression) // 0 or 1, indicates compressed payload

	lenBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lenBytes, uint32(len(body)))

	_, _ = buff.Write(lenBytes)
	_, _ = buff.Write(body)

	// create new request
	req, _ := http.NewRequest(r.Method, r.URL.String(), buff)
	req.Header = r.Header

	// remove content length header
	req.Header.Del(headerContentLength)

	return req

}

func handleGRPCResponse(resp *http.Response) (*http.Response, error) {

	code := resp.Header.Get(headerGRPCStatusCode)
	if code != "0" && code != "" {
		buff := bytes.NewBuffer(nil)
		grpcMessage := resp.Header.Get(headerGRPCMessage)
		j, _ := json.Marshal(grpcMessage)
		buff.WriteString(`{"error":` + string(j) + ` ,"code":` + code + `}`)

		resp.Body = ioutil.NopCloser(buff)
		resp.StatusCode = 500

		return resp, nil
	}

	prefix := make([]byte, 5)
	_, _ = resp.Body.Read(prefix)

	resp.Header.Del(headerContentLength)

	return resp, nil

}