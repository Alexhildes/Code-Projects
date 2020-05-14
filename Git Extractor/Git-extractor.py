import os
import zlib

dir = 'C:\\Schneider\\objects'
f = open(dir + '\\recover.txt','w')

for root, dirs, files in os.walk(dir, topdown=True):        #Loop through directory
    for name in files:

        filename = os.path.join(root, name)

        compressed_contents = open(filename, 'rb').read()
        compressed_contents_obj = zlib.decompressobj()
        decompressed_contents = compressed_contents_obj.decompress(compressed_contents)

        if 'L02-0.png' in str(decompressed_contents):
            print(filename,file=f)      #Print filename
            print('',file=f)

            print(decompressed_contents, file=f)

