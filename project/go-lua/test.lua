--0返回值0参数
function GetStr()
        print "world"
end

--多参数0返回值
function GetBigger(a,b)
    if a >= b then
        print (a)
    else
        print (b)
    end
end

--0参数1返回值
function GetResult()
    return "hello"
end

--多参数1返回值
function Compare(a,b)
    if a >= b then
        return a
    else
        return b
    end
end

--1参数多返回值
function MoreReturn(a)
    if (a == 10) then
        return "world","hello","golang"
    end
end