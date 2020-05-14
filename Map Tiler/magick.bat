
magick convert -density 5 -trim "L02-Background.pdf" -quality 100 -flatten L02-0.png
magick convert -density 10 -trim "L02-Background.pdf" -quality 100 -flatten L02-1.png
magick convert -density 20 -trim "L02-Background.pdf" -quality 100 -flatten L02-2.png
magick convert -density 40 -trim "L02-Background.pdf" -quality 100 -flatten L02-3.png
magick convert -density 80 -trim "L02-Background.pdf" -quality 100 -flatten L02-4.png
magick convert -density 160 -trim "L02-Background.pdf" -quality 100 -flatten L02-5.png
magick convert -density 320 -trim "L02-Background.pdf" -quality 100 -flatten L02-6.png
magick convert -density 640 -trim "L02-Background.pdf" -quality 100 -flatten L02-7.png
magick convert L02-0.png -modulate 100,50,100 L02-0.png
magick convert L02-1.png -modulate 100,50,100 L02-1.png
magick convert L02-2.png -modulate 100,50,100 L02-2.png
magick convert L02-3.png -modulate 100,50,100 L02-3.png
magick convert L02-4.png -modulate 100,50,100 L02-4.png
magick convert L02-5.png -modulate 100,50,100 L02-5.png
magick convert L02-6.png -modulate 100,50,100 L02-6.png
magick convert L02-7.png -modulate 100,50,100 L02-7.png
