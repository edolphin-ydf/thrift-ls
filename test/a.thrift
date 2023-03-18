include "B.c.d.thrift"

struct BaseResp {
    1: string StatusMessage,
    2: i32 StatusCode ,
    3: optional map<string,string> Extra 
}

struct BaseResp2 {
    1: BaseResp StatusMessage,
    2: B.c.d.BaseResp1 StatusCode ,
	2: B.c.d.BaseResp1 
}

service TestService {
	void TestFunc(1: B.c.d.BaseResp1 req)
}
