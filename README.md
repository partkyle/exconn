exconn
======

A simple io.WriteCloser implementation of UDP via the syscall library.

This gives the same interface as the `net.Dial("udp", ...)`,
without the need to reset connections when the remote address is no longer available.


### Example

```
conn, err := exconn.Dial("127.0.0.1:8125")
if err != nil {
  log.Fatal(err)
}

defer conn.Close()

_, err := fmt.Fprintf(conn, "test.the.udp.connection:1234|c")
if err != nil {
  log.Println(err)
}
```
