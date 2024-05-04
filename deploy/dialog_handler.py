#! /usr/bin/python3
# coding=utf8

import os
import json
import time

from pywinauto import base_wrapper
from pywinauto import findwindows
from pywinauto.controls import hwndwrapper
from pywinauto.controls import win32_controls
from uiautomation import DocumentControl
import clipboard


# 响应体
class Response:
    # 处理成功标识
    result = False
    # 处理失败消息
    message = ''
    # 响应数据
    content = None

    # 转成 JSON格式
    def to_json(self):
        return json.dumps(self, default=lambda o: o.__dict__)


# 获取所有对话框标题
def list_windows():
    print('list windows')
    res = Response()
    try:
        res.result = True
        res.content = []
        # 获取所有对话框
        dialogs = findwindows.find_windows()
        for v in dialogs:
            hw = hwndwrapper.HwndWrapper(v)
            res.content.append(hw.window_text())
    except BaseException as e:
        res.message = repr(e)
    return res


# 根据对话框标题获取桌面对话框内容
def get_window_dialog(title):
    print('get window dialog ' + title)
    res = Response()
    try:
        cur = 0
        res.content = {}
        # 获取所有对话框
        dialogs = findwindows.find_windows()
        for v in dialogs:
            hw = hwndwrapper.HwndWrapper(v)
            # 过滤出所有该标题对话框
            if hw.window_text() == title:
                children = hw.children()
                txt = ''
                for child in children:
                    text = str(base_wrapper.BaseWrapper.window_text(child))
                    txt += text
                res.content[cur] = txt
                res.result = True
                cur += 1
    except BaseException as e:
        res.message = repr(e)
    return res


# 点击对话框的按钮
def click_button(title, button, index):
    print("click button " + title + " " + button)
    res = Response()
    try:
        cur = 0
        # 找出所有对话框
        dialogs = findwindows.find_windows()
        for v in dialogs:
            hw = hwndwrapper.HwndWrapper(v)
            # 过滤出该标题对话框
            if hw.window_text() == title:
                # 获取指定位置的同标题对话框
                if cur != index:
                    cur += 1
                    continue
                children = hw.children()
                for child in children:
                    text = str(base_wrapper.BaseWrapper.window_text(child))
                    # 找出要点击的按钮
                    if text.__contains__(button):
                        win32_controls.ButtonWrapper.click(child, double=True)
                        res.result = True
    except BaseException as e:
        res.message = repr(e)
    return res


# 关闭对话框
def close_window(title, index):
    print('close window ' + title + " " + index)
    res = Response()
    try:
        cur = 0
        # 找出所有对话框
        dialogs = findwindows.find_windows()
        for v in dialogs:
            hw = hwndwrapper.HwndWrapper(v)
            # 过滤出该标题对话框
            if hw.window_text() == title:
                # 获取指定位置的同标题对话框
                if cur != index:
                    cur += 1
                    continue
                hw.close()
                res.result = True
    except BaseException as e:
        res.message = repr(e)
    return res


# 获取 CMD 对话框文本
def get_cmd_dialog(index):
    print('get cmd dialog ' + index)
    res = Response()
    try:
        # 获取 CMD 对话框
        window = DocumentControl(Name="Text Area", searchDepth=3, foundIndex=index)
        if window.Exists():
            # 复制对话框内容
            window.SendKeys('{Ctrl}A')
            window.SendKeys('{Ctrl}C')
            data = clipboard.paste()
            res.content = str(data)
            res.content = res.content.replace('\n', r'\n')
            res.result = True
    except BaseException as e:
        res.message = repr(e)
    return res


# 向 CMD 对话框输入字符
def input_to_cmd(content, index):
    print('input to cmd ' + content + " " + index)
    res = Response()
    try:
        # 获取 CMD 对话框
        window = DocumentControl(Name="Text Area", searchDepth=3, foundIndex=index)
        window.SendKeys(content)
        res.result = True
    except BaseException as e:
        res.message = repr(e)
    return res


# 监听文件，执行操作
def main():
    username = os.getlogin()
    print('run handler on user ' + username)
    # 循环读取文件内容
    while True:
        time.sleep(3)
        try:
            print("reading")
            with open(r'dialog_conversation.txt', 'r+', encoding='utf-8') as file:
                content = file.read()
                lines = content.split('\n')
                # 获取最后一行字符
                line = lines[len(lines) - 1]
                # 是请求 ask 数据
                if len(lines) % 2 == 1 and content != '' and line.startswith('ask: '):
                    rsp = ''
                    # 解析 JSON 请求数据
                    data = json.loads(line[4:])
                    method = data['method']
                    # 获取桌面对话框
                    if method == 'list_windows':
                        rsp = list_windows().to_json()
                    # 获取桌面对话框内容
                    elif method == 'get_window_dialog':
                        rsp = get_window_dialog(data['title']).to_json()
                    # 点击对话框按钮
                    elif method == 'click_button':
                        rsp = click_button(data['title'], data['button'], data['index']).to_json()
                    # 关闭对话框
                    elif method == 'close_window':
                        rsp = close_window(data['title'], data['index']).to_json()
                    # 获取 CMD 对话框
                    elif method == 'get_cmd_dialog':
                        rsp = get_cmd_dialog(data['index']).to_json()
                    # 向 CMD 对话框输入字符
                    elif method == 'input_to_cmd':
                        rsp = input_to_cmd(data['content'], data['index']).to_json()
                    else:
                        print('no method match: ' + method)
                    # 响应数据写到文件尾部
                    if rsp != '':
                        file.seek(0, 2)
                        file.write('\nanswer: ')
                        file.write(rsp)
                else:
                    print('file content mismatch: ' + line)
                file.close()
        except BaseException as e:
            print("error: ", e)


if __name__ == '__main__':
    main()
