package statsd

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"testing"
	"time"
)

type mockClosedUDPConn struct {
	i int
	net.Conn
}

func (c *mockClosedUDPConn) Write(p []byte) (int, error) {
	c.i++
	if c.i == 2 {
		return 0, errors.New("test error")
	}
	return 0, nil
}

func (c *mockClosedUDPConn) Close() error {
	return nil
}

func mockUDPClosed(string, string, time.Duration) (net.Conn, error) {
	return &mockClosedUDPConn{}, nil
}

func testClient(t *testing.T, f func(*Client), options ...Option) {
	dialTimeout = mockDial
	defer func() { dialTimeout = net.DialTimeout }()

	options = append([]Option{
		FlushPeriod(0),
		ErrorHandler(expectNoError(t)),
	}, options...)
	c, err := New(options...)
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	f(c)
}

func testOutput(t *testing.T, want string, f func(*Client), options ...Option) {
	testClient(t, func(c *Client) {
		f(c)
		c.Close()

		got := getOutput(c)
		if got != want {
			t.Errorf("Invalid output, got:\n%q\nwant:\n%q", got, want)
		}
	}, options...)
}

func expectNoError(t *testing.T) func(error) {
	return func(err error) {
		t.Errorf("ErrorHandler should not receive an error: %v", err)
	}
}

type testBuffer struct {
	buf bytes.Buffer
	err error
	net.Conn
}

func (c *testBuffer) Write(p []byte) (int, error) {
	if c.err != nil {
		return 0, c.err
	}
	return c.buf.Write(p)
}

func (c *testBuffer) Close() error {
	return c.err
}

func getBuffer(c *Client) *testBuffer {
	if mock, ok := c.conn.w.(*testBuffer); ok {
		return mock
	}
	return nil
}

func getOutput(c *Client) string {
	if c.conn.w == nil {
		return ""
	}
	return getBuffer(c).buf.String()
}

func mockDial(string, string, time.Duration) (net.Conn, error) {
	return &testBuffer{}, nil
}

func testNetwork(t *testing.T, network string) {
	received := make(chan bool)
	server := newServer(t, network, testAddr, func(p []byte) {
		s := string(p)
		if s != "test_key:1|c" {
			t.Errorf("invalid output: %q", s)
		}
		received <- true
	})
	defer server.Close()

	c, err := New(
		Address(server.addr),
		Network(network),
		ErrorHandler(expectNoError(t)),
		LazyConnect(),
		FlushesBetweenReconnect(1),
	)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	c.Increment(testKey)
	c.Close()
	select {
	case <-time.After(100 * time.Millisecond):
		t.Error("server received nothing after 100ms")
	case <-received:
	}
}

type server struct {
	t      testing.TB
	addr   string
	closer io.Closer
	closed chan bool
}

func newServer(t testing.TB, network, addr string, f func([]byte)) *server {
	s := &server{t: t, closed: make(chan bool)}
	switch network {
	case "udp":
		laddr, err := net.ResolveUDPAddr("udp", addr)
		if err != nil {
			t.Fatal(err)
		}
		conn, err := net.ListenUDP("udp", laddr)
		if err != nil {
			t.Fatal(err)
		}
		s.closer = conn
		s.addr = conn.LocalAddr().String()
		go func() {
			buf := make([]byte, 1024)
			for {
				n, err := conn.Read(buf)
				if err != nil {
					s.closed <- true
					return
				}
				if n > 0 {
					f(buf[:n])
				}
			}
		}()
	case "tcp":
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}
		s.closer = ln
		s.addr = ln.Addr().String()
		go func() {
			for {
				conn, err := ln.Accept()
				if err != nil {
					s.closed <- true
					return
				}
				p, err := ioutil.ReadAll(conn)
				if err != nil {
					t.Fatal(err)
				}
				if err := conn.Close(); err != nil {
					t.Fatal(err)
				}
				f(p)
			}
		}()
	default:
		t.Fatalf("Invalid network: %q", network)
	}

	return s
}

func (s *server) Close() {
	if err := s.closer.Close(); err != nil {
		s.t.Error(err)
	}
	<-s.closed
}
