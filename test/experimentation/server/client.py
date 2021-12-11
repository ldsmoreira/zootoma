#!/usr/bin/env python3

import socket

HOST = '127.0.0.1'  # The server's hostname or IP address
PORT = 9000        # The port used by the server

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect((HOST, PORT))
    s.sendall(b'get /home/xap 12\n\nlucas::toma\nleo::toma\narthur::toma\n\nmoreiralucas')
    data = s.recv(1024)

print('Received', repr(data))