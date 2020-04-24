from flask import make_response
import json
from pymongo import MongoClient
from api.blog.article.markdown_tool import parsing
import pytz
import time
import datetime

tz = pytz.timezone('Asia/Shanghai')
client = MongoClient('mongodb://127.0.0.1:27017')
blog = client.my_blog
article = blog.article
info = blog.info


def get_base(data):
    try:
        aid = int(json.loads(data)['aid'])
    except ValueError:
        return make_response(json.dumps({'success': False}))
    data = article.find_one({'aid': aid}, {'_id': 0})
    if data:
        # 如果有这个帖子
        update_data = {'read_num': data['read_num'] + 1}
        article.update_one({'aid': aid}, {'$set': update_data})  # 更新阅读量
        data.update(update_data)
        return make_response(
            json.dumps({'success': True, 'article_data': data}))
    else:
        return make_response(json.dumps({'success': False}))


def post_base(data):
    data = json.loads(data)
    if 'content' in data and 'title' in data and 'type' in data:
        if 'aid' not in data:
            article_num = article.find().count() + 1  # 文章总数
            time_stamp = time.time()  # 获取时间戳
            time_str = datetime.datetime.fromtimestamp(int(time_stamp), pytz.timezone('Asia/Shanghai')).strftime(
                '%Y-%m-%d %H:%M:%S')  # 获取时间字符串
            try:
                insert_data = {'content': data['content'], 'title': data['title'], 'aid': article_num,
                               'directory': parsing(data['content']), 'time_stamp': time_stamp, 'time_str': time_str,
                               'type': int(data['type']), 'read_num': 0, 'reply_num': 0}  # 插入的文章数据
            except ValueError:
                response = json.dumps({'success': False})
                return make_response(response)
            article.insert_one(insert_data)
            info.update_one({}, {'$set': {'article': article_num}})  # 文章总数更新
            return make_response(json.dumps({'success': True, 'aid': article_num}))
        else:
            time_stamp = time.time()  # 获取时间戳
            time_str = datetime.datetime.fromtimestamp(int(time_stamp), pytz.timezone('Asia/Shanghai')).strftime(
                '%Y-%m-%d %H:%M:%S')  # 获取时间字符串
            try:
                article.update_one({'aid': int(data['aid'])}, {
                    '$set': {'content': data['content'], 'title': data['title'], 'directory': parsing(data['content']),
                             'time_stamp': time_stamp, 'time_str': time_str,
                             'type': int(data['type'])}})
            except ValueError:
                response = json.dumps({'success': False})
                return make_response(response)
            return make_response(json.dumps({'success': True}))
    else:
        return make_response(json.dumps({'success': False}))
