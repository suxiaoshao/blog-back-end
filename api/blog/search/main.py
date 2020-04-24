from pymongo import MongoClient
import json
from flask import make_response

client = MongoClient('mongodb://127.0.0.1:27017')
my_blog = client.my_blog
article = my_blog.article


def search_num(data):
    data = json.loads(data)
    search_name = data['search_name']  # 搜索的关键字
    find_condition = {'title': {'$regex': f".*{search_name}.*", "$options": "$i"}}  # 筛选条件
    article_num = article.find(find_condition).count()  # 符合条件的总数
    response = make_response(json.dumps({'success': True, 'article_num': article_num}))
    return response


def search_list(data):
    data = json.loads(data)
    search_name = data['search_name']  # 搜索的关键字
    try:
        limit_num = int(data['limit_num'])  # 发送的数量
        find_condition = {'title': {'$regex': f".*{search_name}.*", "$options": "$i"}}  # 筛选条件
        article_num = article.find(find_condition).count()  # 符合条件的总数
        offset = int(data['offset'])  # 偏移量
    except ValueError:
        response = json.dumps({'success': False})
        return make_response(response)
    if article_num > offset:
        all_data = list(
            article.find(find_condition, {'_id': 0, 'title': 1, 'aid': 1, 'type': 1, 'time_str': 1}).sort('aid',
                                                                                                          -1).skip(
                offset).limit(limit_num))
        response = make_response(json.dumps({'success': True, 'data': all_data}))
        return response
    else:
        response = make_response(json.dumps({'success': False}))
        return response


def main():
    # search_list(json.dumps({'search_name': "python", "limit_num": 10, "offset": 0}))
    all_data = list(article.find())
    for data in all_data:
        article.update_one(data, {'$unset': {'text': 1}})


if __name__ == '__main__':
    main()
