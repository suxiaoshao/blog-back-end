from flask import make_response
import json
from pymongo import MongoClient
import hashlib

client = MongoClient('mongodb://127.0.0.1:27017')
blog = client.my_blog
info = blog.info


# 登陆函数
def login(data):
    data = json.loads(data)
    if 'user' in data and 'password' in data:  # 判断用户发送数据是否有这两个值
        real_data = info.find_one()  # 获得管理者信息
        result = judge_data(data, real_data)  # 判断账号密码是否正确,正确返回编码后的数据，不正确返回None
        if result:
            response = make_response(json.dumps({'success': True}))
            for i in result:
                response.set_cookie(i, result[i])  # 设置cookie
            return response
        else:
            return make_response(json.dumps({'success': False}))
    return make_response(json.dumps({'success': False}))


# 判断是否登陆成功
def judge_login(cookie):
    if other_login(cookie):
        return make_response({'success': True})
    else:
        # cookie 错误，删除所有cookie
        res = make_response({'success': False})
        for i in cookie:
            res.set_cookie(i, '', expires=0)
        return res


def other_login(cookie):
    real_data = info.find_one()  # 获得管理者信息
    # cookie 正确
    if judge_data({'user': cookie.get('uid'), 'password': cookie.get('pid')}, real_data):
        return True
    else:
        return False


# 判断密码账号是否正确，正确返回密码账号的加密dict,错误返回None
def judge_data(judged_data, standard_data):
    standard_user = string_to_md5_3(standard_data['user'])
    standard_password = string_to_md5_3(standard_data['password'])
    if judged_data['user'] == standard_user and judged_data['password'] == standard_password:
        return {'uid': standard_user, 'pid': standard_password}
    else:
        return None


# 加密函数
def string_to_md5_3(string):
    string = hashlib.md5(string.encode('ascii')).hexdigest()[:16]
    string = hashlib.md5(string.encode('ascii')).hexdigest()[16:]
    string = hashlib.md5(string.encode('ascii')).hexdigest()
    return string
