package exconn

import (
	"fmt"
	"io"
	"net"
	"syscall"
)

// Exconn is an implementation of UDP using syscalls. It does not suffer
// from the disconnect issue of the standard library net.Dial method.
type exconn struct {
	// file descriptor for the socket
	fd int

	remoteAddr syscall.Sockaddr
}

func Dial(addr string) (io.WriteCloser, error) {
	fd, err := syscall.Socket(
		syscall.AF_INET,
		syscall.SOCK_DGRAM,
		syscall.PROT_NONE,
	)

	if err != nil {
		return nil, err
	}

	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}

	ip := udpAddr.IP.To4()
	if ip == nil {
		return nil, fmt.Errorf("exconn: cannot accept ip %q as an IPv4 Address", udpAddr.IP)
	}

	remoteAddr := &syscall.SockaddrInet4{Addr: [4]byte{ip[0], ip[1], ip[2], ip[3]}, Port: udpAddr.Port}

	wc := &exconn{
		fd:         fd,
		remoteAddr: remoteAddr,
	}

	return wc, nil
}

// Close closes the underlying file descriptor for the socket
func (e *exconn) Close() error {
	return syscall.Close(e.fd)
}

// Write writes the byte to the socket.
// It returns the number of bytes written and any error that occurs.
func (e *exconn) Write(b []byte) (int, error) {
	err := syscall.Sendmsg(e.fd, b, nil, e.remoteAddr, syscall.MSG_DONTWAIT)
	if err != nil {
		return 0, err
	}

	return len(b), nil
}
