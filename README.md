exconn
======

A simple io.WriteCloser implementation of UDP via the syscall library.

This gives the same interface as the `net.Dial("udp", ...)`,
without the need to reset connections when the remote address is no longer available.
