import re
from pymongo import MongoClient

client = MongoClient('mongodb://127.0.0.1:27017')
my_blog = client.my_blog
article = my_blog.article


def get_result(item):
    name = item[1]
    name = name.replace(' ', '-')
    name = name.replace('.', '')
    name = name.replace(' ', '-')
    name = name.lower()
    return name


def parsing(md):
    title_list = re.findall(r'[\n|](#+) +(.*?)\n', '\n' + md)
    title_list = list(map(get_result, title_list))
    return title_list
