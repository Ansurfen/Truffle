# mutil thread is unsafe, if Cross-compiling
import threading
import os
import time
cc = bool(False)


def copyFile(src: str, dst: str) -> None:
    with open(src, 'r', encoding='utf-8') as srcFile:
        data = srcFile.read()
        with open(dst, 'w+', encoding='utf-8') as dstFile:
            dstFile.write(data)
            dstFile.close()
        srcFile.close()


if __name__ == '__main__':
    dirs = ['ws', 'topic']
    path = os.getcwd()
    if (path.rfind("auto") + len("auto") == len(path)):
        os.chdir(os.path.abspath('..'))
    if (not cc):
        # shutdown cross-compiling
        parent = os.getcwd()
        for dir in dirs:
            path: str = '{parent}\\{child}'.format(
                parent=parent, child=dir)
            copyFile('.\\auto\\build.bat', path + '/build.bat')
            os.chdir(path)
            os.system('build.bat')
            os.remove('build.bat')
            os.chdir(parent)
    else:
        pass
