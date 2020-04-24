from flask import make_response
import json
from pymongo import MongoClient
from werkzeug.utils import secure_filename
from bson.objectid import ObjectId

client = MongoClient('mongodb://127.0.0.1:27017')
blog = client.my_blog
img = blog.img


def upload_img(files):
    file_set = {'jpg', 'png', 'JPG', 'PNG', 'JEPG', 'jepg', 'jpeg', 'JPEG'}  # 允许的图片文件扩展名
    file = files['file']
    file_tag = secure_filename(file.filename).split('.')[-1]  # 文件扩展名
    if file_tag in file_set:
        _id = img.insert_one(
            {'binary': file.read(), 'file_extension': file_tag, 'tag': 'article'}).inserted_id  # 上传图片的_id
        return make_response(
            json.dumps({'url': f"http://www.sushao.top/api/blog/article/img/{_id}.{file_tag}", 'success': True}))
    else:
        return make_response(json.dumps({'success': False}))


def read_img(pid):
    data = img.find_one({'_id': ObjectId(pid.split('.')[0])})
    if data:
        binary = data['binary']
        res = make_response(binary)
        res.headers['Content-Type'] = 'image/png'
        return res
    else:
        return '404'


def main():
    with open('D:\\Downloads\\75924852ee7ca15bdd58b733cb4f70eb.png', 'rb') as f:
        jpg = f.read()
    _id = img.insert_one({'binary': jpg, 'file_extension': 'png', 'tag': 'article'}).inserted_id
    print(_id)


if __name__ == '__main__':
    main()
