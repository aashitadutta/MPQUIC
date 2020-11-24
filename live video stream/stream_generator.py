import io
import time
import picamera
import socket


IP = "localhost"
PORT = 8002


sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
sock.bind((IP, PORT))
sock.listen(1)
print('Listening at', sock.getsockname())

sc, sockname = sock.accept()

frames = 0
start_time = time.time()

MAX_BYTES = 65535

with picamera.PiCamera() as camera:

    camera.resolution = (300, 300)
    stream = io.BytesIO()

    for _ in camera.capture_continuous(stream, format='jpeg', use_video_port=True):

        # Truncate the stream to the current position (in case
        # prior iterations output a longer image)
        stream.truncate()
        stream.seek(0)
        print(len(stream.getbuffer()))

        time.sleep(0.05)

        buffer_size = str(len(stream.getbuffer()))
        buffer_size_data = buffer_size.ljust(20, ':')

        sc.sendall(buffer_size_data.encode())

        sc.sendall(stream.getbuffer())
        print("frame sent")

        frames += 1
        current_time = time.time()
        fps = frames / (current_time - start_time)
        print("fps: ", fps)

        # if (frames == 100):
        #     break
