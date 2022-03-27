#!/usr/bin/python3

from PIL import Image
import sys


def main(argv):
    input_file = argv[0]
    if input_file == "" or None:
        print("file not found")
        return

    output_file = ".".join(input_file.split(".")[:-1]) + "-resized" + "." + input_file.split(".")[-1]
    print(output_file)

    with Image.open(input_file) as im:

        # LOB image ratio for a 4 x 6 postcard
        desired_ratio = 0.68

        # keep width the same for simplicity
        target_width = im.width
        
        # adjust height so that image fits proportions
        target_height = round((1 / desired_ratio) * im.width)

        im_resized = im.resize((target_height, target_width))

        im_resized.save(output_file)


if __name__ == "__main__":
    main(sys.argv[1:])