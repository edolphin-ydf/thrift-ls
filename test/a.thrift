include "b.thrift"

struct BaseResp {
    1: string StatusMessage,
    2: i32 StatusCode ,
    3: optional map<string,string> Extra 
}

struct BaseResp2 {
    1: BaseResp StatusMessage,
    2: b.BaseResp1 StatusCode ,
    3: optional map<string,string> Extra 
}
