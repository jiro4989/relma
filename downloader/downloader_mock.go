package downloader

import (
	"net/http"
	"strings"
)

type MockDownloader struct {
	Body        string
	Bodies      []string
	BodyCounter int
	Err         error
	ReadErr     error
}

type MockReadCloser struct {
	r   *strings.Reader
	err error
}

func (m *MockReadCloser) Read(p []byte) (int, error) {
	if m.err != nil {
		return -1, m.err
	}
	return m.r.Read(p)
}

func (m *MockReadCloser) Close() error {
	return nil
}

func (d *MockDownloader) Download(url string) (*http.Response, error) {
	if d.Err != nil {
		return nil, d.Err
	}

	if 0 < len(d.Bodies) {
		mock := &MockReadCloser{
			r:   strings.NewReader(d.Bodies[d.BodyCounter]),
			err: d.ReadErr,
		}
		resp := &http.Response{
			Body: mock,
		}
		d.BodyCounter++
		return resp, nil
	}

	mock := &MockReadCloser{
		r:   strings.NewReader(d.Body),
		err: d.ReadErr,
	}
	resp := &http.Response{
		Body: mock,
	}
	return resp, nil
}
