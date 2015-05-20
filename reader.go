package main

import (
  "fmt"
  "net"
)

const BUF_SIZE = 1024

type Reader struct {
  buf []byte
  conn *net.UDPConn
}

func (r Reader) Read() []byte {
    n, _, _ := r.conn.ReadFromUDP(r.buf)
    return r.buf[:n]
}

func (r Reader) Close() {
  r.conn.Close()
}

func NewReader(host string, port string) (*Reader, error) {
  addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", host, port))
  if err != nil {
    return nil, err
  }

  conn, err := net.ListenMulticastUDP("udp", nil, addr)
  if err != nil {
    return nil, err
  }

  conn.SetReadBuffer(BUF_SIZE)

  return &Reader{make([]byte, BUF_SIZE), conn}, nil
}
