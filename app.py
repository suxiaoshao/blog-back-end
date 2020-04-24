from flask import Flask, request, make_response
from flask_cors import *
from api.api import Api

app = Flask(__name__)

CORS(app, supports_credentials=True, origins=[r'.*\.sushao\.top', r".*localhost:.*"])


@app.route('/api/blog/<name>/base', methods=['get', 'post'])
def bilibili_base(name):
    if name == 'article':
        if request.method == 'POST':
            return Api.blog.article.get_content(request.data)
    elif name == 'home':
        if request.method == 'POST':
            return Api.blog.home.get_article_num(request.data)
    elif name == 'search':
        if request.method == 'POST':
            return Api.blog.search.search_num(request.data)
    elif name == 'reply':
        if request.method == "POST":
            return Api.blog.reply.get_reply(request.data)


@app.route('/api/blog/<name>/<method>', methods=['post'])
def bilibili_av_save(name, method):
    if name == 'article':
        # 判断是否cookie是否正确
        if Api().blog.home.other_login(request.cookies):  # 正确的话继续
            if method == 'upload_pic':
                return Api().blog.article.upload_img(request.files)
            elif method == 'upload_content':
                return Api().blog.article.post_base(request.data)
        else:
            res = make_response({'success': False})
            for i in request.cookies:
                res.set_cookie(i, '', expires=0)
    elif name == 'home':
        if method == 'article_list':
            return Api().blog.home.get_article_list(request.data)
        elif method == 'login':
            return Api().blog.home.login(request.data)
        elif method == 'is_login':
            return Api().blog.home.is_login(request.cookies)
    elif name == "search":
        if method == 'search_list':
            return Api().blog.search.search_data(request.data)
    elif name == 'reply':
        if method == "post":
            return Api().blog.reply.post_reply(request.data)


@app.route("/api/blog/<name>/img/<pid>")
def bilibili_img(name, pid=None):
    if name == 'article':
        return Api().blog.article.get_pic(pid)
    elif name == 'home':
        return Api().blog.home.get_random_jpg()


@app.route('/<name>/<tag>/<aid>')
@app.route('/<name>/<tag>')
@app.route('/<name>')
@app.route('/')
def main_html(name=None, tag=None, aid=None):
    return open('static/index.html', 'r', encoding='utf-8').read()


if __name__ == '__main__':
    app.config['DEBUG'] = True
    app.run(host='0.0.0.0')
