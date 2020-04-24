from pymongo import MongoClient
import json
from flask import make_response
import random
from tool.webimfor import GetWeb

client = MongoClient('mongodb://127.0.0.1:27017')
blog = client.my_blog
img = blog.img
article = blog.article


def get_random_jpg():
    flag_num = random.randint(1, img.count_documents({'flag_num': {'$exists': True}}))
    data = img.find_one({'flag_num': flag_num})
    try:
        if 'binary' in data:
            # 如果文件已保存，直接发送
            response = make_response(data['binary'])
            response.headers['Content-Type'] = 'image/png'
            return response
        else:
            # 如果文件未保存，保存后发送
            binary = GetWeb(data['url']).get_content()
            response = make_response(binary)
            response.headers['Content-Type'] = 'image/png'
            img.delete_one(data)
            img.insert_one({'file_extension': data['file_extension'], 'tag': 'wallpaper', 'binary': binary,
                            'flag_num': data['flag_num'], 'url': data['url']})
            return response
    except TypeError:
        print(flag_num)


def get_base(data):
    data = json.loads(data)
    article_num = article.find(data['find_condition']).count()  # 获得筛选结果的数值
    return make_response(json.dumps({'success': True, 'article_num': article_num}))


def get_list(data):
    try:
        data = json.loads(data)
        find_condition = data['find_condition']  # 筛选条件
        limit_num = int(data['limit_num'])  # 发送的数量
        article_num = article.find(find_condition).count()  # 符合条件的总数
        offset = int(data['offset'])  # 偏移量
    except ValueError:
        response = json.dumps({'success': False})
        return make_response(response)
    if article_num > offset:
        data_list = list(
            article.find(find_condition, {'_id': 0, 'title': 1, 'aid': 1, 'type': 1, 'time_str': 1}).sort('aid',
                                                                                                          -1).skip(
                offset).limit(limit_num))
        return make_response(json.dumps({'success': True, 'data': data_list}))
    response = make_response(json.dumps({'success': False}))
    return response
