from pymongo import MongoClient

client = MongoClient('mongodb://127.0.0.1:27017')
blog = client.my_blog
img = blog.img


def main():
    print(img.count_documents({"tag": "wallpaper", 'flag_num': {'$exists': False}, "done": True}))


if __name__ == '__main__':
    main()
