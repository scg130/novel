import ctypes
import inspect
import sys
import threading
import time
import tty

import pyautogui


# # 双开魂土
def dance():
    while True:
        # 业原火
        pyautogui.click(1627, 729, 2, 0.5, button='left')
        # # 组队魂土 队员
        # pyautogui.click(1610, 721, 2, 0.25, button='left')
        time.sleep(1)
        # # 双开魂土
        # pyautogui.click(1744, 471, 2, 0.25, button='left')
        # pyautogui.click(812, 986, 2, 0.25, button='left')
        # time.sleep(23)
        # pyautogui.click(1744, 471, 4, 1, button='left')
        # pyautogui.click(812, 986, 4, 1, button='left')
        # time.sleep(2)


def _async_raise(tid, exctype):
    tid = ctypes.c_long(tid)

    if not inspect.isclass(exctype):
        exctype = type(exctype)

    res = ctypes.pythonapi.PyThreadState_SetAsyncExc(
        tid, ctypes.py_object(exctype))

    if res == 0:
        raise ValueError("invalid thread id")

    elif res != 1:
        ctypes.pythonapi.PyThreadState_SetAsyncExc(tid, None)
        raise SystemError("PyThreadState_SetAsyncExc failed")


def stop_thread(thread):
    _async_raise(thread.ident, SystemExit)


if __name__ == '__main__':
    tty.setcbreak(sys.stdin)
    # print(pyautogui.position())
    # sys.exit(1)
    p1 = threading.Thread(target=dance)

    p1.start()
    while True:
        char = ord(sys.stdin.read(1))
        print(char)
        if char == 27:
            stop_thread(p1)
            sys.exit()
