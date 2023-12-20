local pb = require("protobuf")
local protofile = "protofile.proto"

function setup(thread)
    thread:set("protofile", protofile)
    pb.register_file(protofile)
end

function init(args)
    requests = pb.encode("speedTestReq", {t = 12345})
end

function request()
    return wrk.format("POST", nil, nil, requests)
end

function response(status, headers, body)
    local resp = pb.decode("speedTestResp", body)
    -- 处理响应数据
en