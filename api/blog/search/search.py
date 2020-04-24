from api.blog.search.main import search_list, search_num


class Search(object):
    @staticmethod
    def search_data(data):
        return search_list(data)

    @staticmethod
    def search_num(data):
        return search_num(data)
