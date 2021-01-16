package hraftdispatcher

import (
	"crypto/tls"
	"errors"
	"github.com/hashicorp/raft"
	"io"
	"net"
	"time"
)

var (
	errNotAdvertisable = errors.New("local bind address is not advertisable")
	errNotTCP          = errors.New("local address is not a TCP address")
)

// TCPStreamLayer implements StreamLayer interface for plain TCP.
type TCPStreamLayer struct {
	advertise       net.Addr
	listener        *net.TCPListener
	serverTLSConfig *tls.Config
	clientTLSConfig *tls.Config
}

// NewTCPTransport returns a NetworkTransport that is built on top of
// a TCP streaming transport layer.
func NewTCPTransport(
	bindAddr string,
	advertise net.Addr,
	serverTLSConfig *tls.Config,
	clientTLSConfig *tls.Config,
	maxPool int,
	timeout time.Duration,
	logOutput io.Writer,
) (*raft.NetworkTransport, error) {
	return newTCPTransport(bindAddr, advertise, serverTLSConfig, clientTLSConfig, func(stream raft.StreamLayer) *raft.NetworkTransport {
		return raft.NewNetworkTransport(stream, maxPool, timeout, logOutput)
	})
}

func newTCPTransport(bindAddr string,
	advertise net.Addr,
	serverTLSConfig *tls.Config,
	clientTLSConfig *tls.Config,
	transportCreator func(stream raft.StreamLayer) *raft.NetworkTransport) (*raft.NetworkTransport, error) {

	if serverTLSConfig == nil {
		return nil, errors.New("no serverTLSConfig found")
	}
	if clientTLSConfig == nil {
		return nil, errors.New("no clientTLSConfig found")
	}

	// Try to bind
	list, err := tls.Listen("tcp", bindAddr, serverTLSConfig)
	if err != nil {
		return nil, err
	}

	// Create stream
	stream := &TCPStreamLayer{
		advertise:       advertise,
		listener:        list.(*net.TCPListener),
		clientTLSConfig: clientTLSConfig,
	}

	// Verify that we have a usable advertise address
	addr, ok := stream.Addr().(*net.TCPAddr)
	if !ok {
		list.Close()
		return nil, errNotTCP
	}
	if addr.IP == nil || addr.IP.IsUnspecified() {
		list.Close()
		return nil, errNotAdvertisable
	}

	// Create the network transport
	trans := transportCreator(stream)
	return trans, nil
}

// Dial implements the StreamLayer interface.
func (t *TCPStreamLayer) Dial(address raft.ServerAddress, timeout time.Duration) (net.Conn, error) {
	dialer := &net.Dialer{Timeout: timeout}
	return tls.DialWithDialer(dialer, "tcp", string(address), t.clientTLSConfig)
}

// Accept implements the net.Listener interface.
func (t *TCPStreamLayer) Accept() (c net.Conn, err error) {
	return t.listener.Accept()
}

// Close implements the net.Listener interface.
func (t *TCPStreamLayer) Close() (err error) {
	return t.listener.Close()
}

// Addr implements the net.Listener interface.
func (t *TCPStreamLayer) Addr() net.Addr {
	// Use an advertise addr if provided
	if t.advertise != nil {
		return t.advertise
	}
	return t.listener.Addr()
}