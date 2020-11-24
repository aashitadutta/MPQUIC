from matplotlib.animation import FuncAnimation
import matplotlib.pyplot as plt
import matplotlib.image as img

import time
import pathlib

# reading the image
# testImage = img.imread('sample/img0.jpg')

# # displaying the image
# plt.imshow(testImage)
# plt.show()

# time.sleep(5)
# testImage = img.imread('sample/img50.jpg')
# draw()

counter = 0

start_time = time.time()
current_time = time.time()


def grab_frame():
    global counter

    file = pathlib.Path('sample/img' + str(counter) + '.jpg')

    missing_frame_counter = 0
    while (not file.exists() and missing_frame_counter <= 100):
        time.sleep(0.2)
        missing_frame_counter += 1

    image = img.imread('sample/img' + str(counter) + '.jpg')
    counter += 1

    current_time = time.time()
    print("fps: ", counter / (current_time - start_time))

    file.unlink()
    return image


# create axes
ax1 = plt.subplot(111)

# create axes
im1 = ax1.imshow(grab_frame())


def update(i):
    im1.set_data(grab_frame())


ani = FuncAnimation(plt.gcf(), update, interval=25)
plt.show()
