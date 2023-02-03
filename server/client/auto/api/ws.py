from client import HttpClient, WsClient
import asyncio

class GetChannelRequest(object):
    def __init__(self, path: str):
        self.path = path


class GetCGroupsRequset(object):
    def __init__(self, path: str):
        self.path = path

# class GetChannelResponse(object):
#     def __init__(self, entries: dict = {}):
#         for k, v in entries.items():
#             if isinstance(v, dict):
#                 self.__dict__[k] = GetChannelResponse(v)
#             else:
#                 self.__dict__[k] = v


class GetMsgRequest(object):
    def __init__(self, path: str, num: int):
        self.path = path
        self.num = num


class SendMsgRequest(object):
    def __init__(self, path: str, msg: str):
        self.path = path
        self.msg = {
            "author": '',
            "msg": msg
        }


class NewTopicRequest(object):
    def __init__(self, user: str, name: str):
        self.user = user
        self.name = name


class JoinTopicRequest(object):
    def __init__(self, user: str, passcode: str):
        self.user = user
        self.passcode = passcode


cli = HttpClient('http://localhost:9095')

channelRouter = cli.group('/channel')
res = channelRouter.postWithData(
    '/get', data=GetChannelRequest('aba').__dict__)
print(res)
res = channelRouter.postWithData(
    '/group/get', data=GetCGroupsRequset('aba').__dict__)
print(res)

messageRouter = cli.group('/message')
res = messageRouter.postWithData('/get', data=GetMsgRequest('aba', 0).__dict__)
print(res)
res = messageRouter.postWithData(
    '/send', data=SendMsgRequest('aba', 'Hello world').__dict__)
print(res)

topicRouter = cli.group('/topic')
res = topicRouter.postWithData(
    '/new', data=NewTopicRequest('ansurfen', 'topic1').__dict__)
print(res)
res = topicRouter.postWithData(
    '/join', data=JoinTopicRequest('ansurfen', '000000').__dict__)
print(res)


wsCli = WsClient('ws://localhost:9091/ws')
asyncio.run(wsCli.echo())