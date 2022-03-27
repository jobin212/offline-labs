#!/usr/bin/python3

import sys
import os
import lob

def main(args):
    lob.api_key = os.getenv("LOB_API_TEST_KEY")
    from_address_id = args[0]
    to_address_id = args[1]

    front_photo = args[2]
    back_photo = args[3]

    f_front = open(front_photo, 'rb')
    f_back = open(back_photo, 'rb')

    postcard = lob.Postcard.create(
        description = "test postcard",
        to_address = to_address_id,
        from_address = from_address_id,
        front = f_front,
        back = f_back
    )

    print(postcard)
    f_front.close()
    f_back.close()

if __name__ == "__main__":
    main(sys.argv[1:])