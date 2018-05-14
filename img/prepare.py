from PIL import Image

img = Image.open('./lena.png')
gray = img.convert("L")
bin_img = gray.point(lambda x:0 if x < 128 else 255)
bin_img.save('./lena_bin.png')
