import socket

def main():
  # Create a socket.
  sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

  # Bind the socket to a port.
  sock.bind(('', 8080))

  # Listen for incoming connections.
  sock.listen(1)

  # Accept an incoming connection.
  conn, addr = sock.accept()
  print("Incoming address is " + str(addr))

  # Send and receive data.
  data = conn.recv(1024)
  print(data)

  conn.sendall("Hello, world!".encode())

  # Close the socket.
  conn.close()

if __name__ == "__main__":
  main()
