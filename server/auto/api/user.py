from client import HttpClient


class LoginRequest(object):
    def __init__(self, key: str, pwd: str):
        self.key = key
        self.pwd = pwd


class RegisterRequest(object):
    def __init__(self, key: str, pwd: str, name: str):
        self.key = key
        self.pwd = pwd
        self.name = name


cli = HttpClient('http://localhost:8025')

userRouter = cli.group("/user")
res = userRouter.postWithResult(
    '/register', data=RegisterRequest('ansurfen@truffle.com', '123456', 'ansurfen').__dict__)
print(res)
res = userRouter.postWithResult(
    '/login', data=LoginRequest('ansurfen@truffle.com', '123456').__dict__)
print(res)
