# UDP server code
import socket

# Create a socket object
sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

# Bind the socket to a port
server_address = ("localhost", 12345)
sock.bind(server_address)

while True:
    # Receive a message from the client
    message, address = sock.recvfrom(1024)

    # Send a message back to the client
    sock.sendto(message, address)