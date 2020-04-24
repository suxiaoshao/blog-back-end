import gevent.monkey

gevent.monkey.patch_all()
workers = 5  # 定义同时开启的处理请求的进程数量，根据网站流量适当调整
worker_class = "gunicorn.workers.ggevent.GeventWorker"  # 采用gevent库，支持异步处理请求，提高吞吐量
bind = "127.0.0.1:8081"
