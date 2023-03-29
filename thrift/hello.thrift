namespace go hello

struct HelloReq {
    1: string name;
}

struct HelloRes {
    1: string greeting;
}

service Hello {
    HelloRes greet(1: HelloReq req);
}
