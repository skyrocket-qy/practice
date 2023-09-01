# UDP client code
import socket

# Create a socket object
sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

# Bind the socket to a port
server_address = ("localhost", 12345)
sock.bind(server_address)

# Receive a message from the server
message, address = sock.recvfrom(1024)

print("Received message:", message)

# Close the socket
sock.close()