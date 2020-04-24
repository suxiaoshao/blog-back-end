class View(object):
    def __init__(self, number, sums=1, new_line=False):
        self.sum = number
        self.view_num = [0] * sums
        if new_line:
            self.string = '{}//////' * sums + '{}\n'
        else:
            self.string = '\r' + '{}//////' * sums + '{}'

    def display(self):
        print(self.string.format(*self.view_num, self.sum), end='')

    def get(self, index=1, done=1):
        self.view_num[index - 1] += done
        self.display()


__all__ = ['View']
if __name__ == '__main__':
    view = View(100, 299)
    view.get(3, 4)
