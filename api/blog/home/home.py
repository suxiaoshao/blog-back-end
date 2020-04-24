from api.blog.home.main import get_random_jpg, get_base, get_list
from api.blog.home.login import login, judge_login, other_login


class Home(object):
    @staticmethod
    def get_random_jpg():
        return get_random_jpg()

    @staticmethod
    def get_article_num(data):
        return get_base(data)

    @staticmethod
    def get_article_list(data):
        return get_list(data)

    @staticmethod
    def login(data):
        return login(data)

    @staticmethod
    def is_login(cookie):
        return judge_login(cookie)

    @staticmethod
    def other_login(cookie):
        return other_login(cookie)
