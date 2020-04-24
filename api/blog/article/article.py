from api.blog.article.content import *
from api.blog.article.img import *


class Article(object):
    @staticmethod
    def get_content(data):
        return get_base(data)

    @staticmethod
    def post_base(data):
        return post_base(data)

    @staticmethod
    def upload_img(files):
        return upload_img(files)

    @staticmethod
    def get_pic(pid):
        return read_img(pid)
