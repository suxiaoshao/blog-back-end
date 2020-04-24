import json
import re
from flask import make_response
from pymongo import MongoClient
import pytz
import time
import datetime

client = MongoClient('mongodb://127.0.0.1:27017')
my_blog = client.my_blog
article = my_blog.article
reply = my_blog.reply


def validate_email(email):
    if len(email) > 7:
        if re.match("[^@]+@[^@]+[.][^@]+", email):
            return True
    return False


def post_reply(data):
    data = json.loads(data)
    try:
        aid = int(data['aid'])  # 文章aid
    except ValueError:
        response = json.dumps({'success': False})
        return make_response(response)
    content = data['content']  # 评论内容
    name = data['name']  # 评论者名字
    email = data['email']  # 评论者邮箱
    if 'url' in data:
        url = data['url']  # 评论者网址
    else:
        url = None
    time_stamp = time.time()  # 获取时间戳
    time_str = datetime.datetime.fromtimestamp(int(time_stamp), pytz.timezone('Asia/Shanghai')).strftime(
        '%Y-%m-%d %H:%M:%S')  # 获取时间字符串
    rid = int(reply.find().count() + 1)  # 评论id
    article_data = article.find_one({'aid': aid}, {'content': 0})  # 文章数据
    if validate_email(email) and article_data:
        article.update_one(article_data, {'$set': {'reply_num': int(article_data['reply_num']) + 1}})  # 文章评论数据更新
        insert_data = {"rid": rid, 'aid': aid, 'content': content, "name": name, "email": email, "url": url,
                       "time_stamp": time_stamp, "time_str": time_str}
        reply.insert_one(insert_data)
        insert_data.pop("_id")
        response = json.dumps({"success": True, "data": insert_data})
        return make_response(response)

    response = json.dumps({'success': False})
    return make_response(response)


def get_reply(data):
    data = json.loads(data)
    try:
        aid = int(data['aid'])  # 文章aid
        offset = int(data['offset'])  # 偏移量
        limit_num = int(data['limit_num'])  # 限制数
    except ValueError:
        response = json.dumps({'success': False})
        return make_response(response)
    reply_num = reply.find({'aid': aid}).count()  # 文章的评论总数
    if reply_num > offset:
        all_data = list(reply.find({'aid': aid}, {'_id': 0}).sort("rid", 1).skip(offset).limit(limit_num))
        response = json.dumps({'success': True, 'data': all_data, "reply_num": reply_num})
        return make_response(response)
    response = json.dumps({'success': False, 'reply_num': reply_num})
    return make_response(response)


def main():
    # article.update_many({}, {'$set': {'reply_num': 0}})
    pass


if __name__ == '__main__':
    main()
