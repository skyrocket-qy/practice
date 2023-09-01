import socket

def main():
  # Create a socket.
  sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

  # Connect to the server.
  sock.connect(('localhost', 8080))

  # Send data to the server.
  sock.sendall("Hello, world!".encode())

  # Receive data from the server.
  data = sock.recv(1024)
  print(data.decode())

  # Close the socket.
  sock.close()

if __name__ == "__main__":
  main()
