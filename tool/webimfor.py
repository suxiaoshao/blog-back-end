import requests
from lxml import etree
import json


def get_proxy():
    proxy_pool_url = 'http://localhost:5555/random'
    try:
        response = requests.get(proxy_pool_url)
        if response.status_code == 200:
            proxies = {
                'http': 'http://' + response.text,
            }
            return proxies
    except ConnectionError:
        return None


class GetWeb(object):
    def __init__(self, url=''):
        self.url = url
        self.no_proxy = False
        self.en = 'utf-8'
        self.headers = {
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.3\
            6 (KHTML, like Gecko) Chrome/78.0.3904.87 Safari/537.36'}

    def setting(self, no_proxy, en, headers):
        self.no_proxy = no_proxy
        self.en = en
        self.headers = headers

    def setting_no_proxy(self, no_proxy):
        self.no_proxy = no_proxy

    def setting_en(self, en):
        self.en = en

    def setting_headers(self, headers):
        self.headers = headers

    def get(self):
        try:
            if self.no_proxy:
                proxies = get_proxy()
                r = requests.get(self.url, headers=self.headers, proxies=proxies)
                r.raise_for_status()
                r.encoding = self.en
                return r
            else:
                r = requests.get(self.url, headers=self.headers)
                r.raise_for_status()
                return r
        except requests.exceptions.HTTPError:
            # print('爬取失败')
            return None

    def get_html_text(self):
        try:
            r = self.get()
            r.encoding = self.en
            return r.text
        except AttributeError:
            return "{}"

    def get_html_etree(self):
        html = self.get_html_text()
        return etree.HTML(html)

    def get_json(self):
        html = self.get_html_text()
        return json.loads(html)

    def get_content(self):
        try:
            r = self.get()
            return r.content
        except AttributeError:
            return ""

    def post(self, data=None):
        try:
            if self.no_proxy:
                proxies = get_proxy()
                r = requests.post(self.url, headers=self.headers, proxies=proxies, data=data)
                r.raise_for_status()
                r.encoding = self.en
                return r
            else:
                r = requests.post(self.url, headers=self.headers, data=data)
                r.raise_for_status()
                r.encoding = self.en
                return r
        except requests.exceptions.HTTPError:
            print("上传错误")
            return ""

    def post_html_text(self, data=None):
        try:
            r = self.post(data)
            return r.text
        except AttributeError:
            return "{}"

    def post_html_etree(self, data=None):
        html = self.post_html_text(data)
        return etree.HTML(html)

    def post_json(self, data=None):
        html = self.post_html_text(data)
        return json.loads(html)

    def post_content(self, data=None):
        try:
            r = self.post(data)
            return r.content
        except AttributeError:
            return ""

    def download(self, file_path, size=512, open_method='wb'):
        with open(file_path, open_method) as f:
            r = requests.get(self.url, headers=self.headers, stream=True)
            for chuck in r.iter_content(chunk_size=size):
                if chuck:
                    f.write(chuck)


class GetSession(object):
    def __init__(self, session, url):
        self.session = session
        self.headers = {
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.87 Safari/537.36'}
        self.en = 'utf-8'
        self.url = url

    def setting(self, en=None, headers=None):
        self.en = en
        self.headers = headers

    def setting_en(self, en):
        self.en = en

    def setting_headers(self, headers):
        self.headers = headers

    def add_cookies(self, cookies):
        self.session.get(self.url, headers=self.headers, cookies=cookies)
        return self.session

    def get(self):
        try:
            r = self.session.get(self.url, headers=self.headers)
            r.raise_for_status()
            r.encoding = self.en
            # print(r.text)
            return r
        except requests.exceptions.HTTPError:
            print("爬取错误", end='')
            return None

    def get_html_text(self):
        try:
            r = self.get()
            return r.text
        except AttributeError:
            return "{}"

    def get_content(self):
        try:
            r = self.get()
            return r.content
        except AttributeError:
            return ""

    def get_html_etree(self):
        html = self.get_html_text()
        return etree.HTML(html)

    def get_json(self):
        html = self.get_html_text()
        return json.loads(html)

    def post(self, data=None):
        try:
            r = self.session.post(self.url, headers=self.headers, data=data)
            r.raise_for_status()
            r.encoding = self.en
            return r

        except requests.exceptions.HTTPError:
            print("爬取错误", end='')
            return ""

    def post_html_text(self, data=None):
        try:
            r = self.post(data)
            return r.text

        except AttributeError:
            return "{}"

    def post_html_etree(self, data=None):
        html = self.post_html_text(data)
        return etree.HTML(html)

    def post_json(self, data=None):
        html = self.post_html_text(data)
        # print(html)
        return json.loads(html)

    def post_content(self, data=None):
        try:
            r = self.post(data)
            return r.content
        except AttributeError:
            return ""

    def download(self, file_path, size=512, open_method='wb'):
        with open(file_path, open_method) as f:
            r = self.session.get(self.url, headers=self.headers, stream=True)
            for chunk in r.iter_content(chunk_size=size):
                if chunk:
                    f.write(chunk)


def get_session():
    return requests.session()


__all__ = ['GetSession', 'GetWeb', 'get_session']
