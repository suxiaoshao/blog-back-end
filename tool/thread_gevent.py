from gevent import monkey

monkey.patch_all()
import gevent
from threading import Thread


class CreateThreads(object):
    def __init__(self, func, information_list, *parameter):
        self.thread_num = 5
        self.g_num = 20
        self.func = func
        self.informationList = information_list
        self.parameter = parameter
        if type(self.informationList) == int:
            self.repetitiveTask = True
        else:
            self.repetitiveTask = False

    def setting(self, thread_num, g_num):
        self.thread_num = thread_num
        self.g_num = g_num

    def create(self):
        if not self.repetitiveTask:
            threads = []
            len_information = len(self.informationList) // self.thread_num + 1
            for i in range(self.thread_num):
                information = self.informationList[i * len_information:(i + 1) * len_information]
                p1 = RealThread(information, self.func, self.g_num, *self.parameter)
                threads.append(p1)
                p1.start()
            for thread in threads:
                thread.join()
        else:
            threads = []
            len_information = self.informationList // self.thread_num + 1
            for i in range(self.thread_num):
                all_information = len_information
                p1 = RealThread(all_information, self.func, self.g_num, *self.parameter)
                threads.append(p1)
                p1.start()
            for thread in threads:
                thread.join()


class RealThread(Thread):
    def __init__(self, information_list, func, g_num, *parameter):
        super(RealThread, self).__init__()
        self.information_list = information_list
        self.parameter = parameter
        self.func = func
        self.g_Num = g_num
        if type(self.information_list) == int:
            self.repetitiveTask = True
        else:
            self.repetitiveTask = False

    def run(self):
        if not self.repetitiveTask:
            g_time = len(self.information_list) // self.g_Num + 1
            for i in range(g_time):
                g_information = self.information_list[i * self.g_Num:(i + 1) * self.g_Num]
                self.creat_g(g_information)
        else:
            g_time = self.information_list // self.g_Num + 1
            for i in range(g_time):
                g_information = self.g_Num
                self.creat_g(g_information)

    def creat_g(self, g_information):
        if not self.repetitiveTask:
            func = []
            for information in g_information:
                func.append(gevent.spawn(self.func, information, *self.parameter))
            gevent.joinall(func)
        else:
            func = []
            for information in range(g_information):
                func.append(gevent.spawn(self.func, *self.parameter))
            gevent.joinall(func)


__all__ = ['CreateThreads']
