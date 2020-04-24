from api.blog.article.article import Article
from api.blog.home.home import Home
from api.blog.search.search import Search
from api.blog.reply.reply import Reply


class Blog(object):
    article = Article()
    home = Home()
    search = Search()
    reply = Reply()
