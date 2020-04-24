from api.blog.reply.main import get_reply, post_reply


class Reply(object):
    @staticmethod
    def get_reply(data):
        return get_reply(data)

    @staticmethod
    def post_reply(data):
        return post_reply(data)
