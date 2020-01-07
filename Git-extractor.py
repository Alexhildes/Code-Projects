import os
import zlib

dir = 'C:\\Schneider\\objects'

for root, dirs, files in os.walk(dir, topdown=True):
    for name in files:
        print(os.path.join(root, name))
