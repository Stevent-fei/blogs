package main

/*
def segmentation(s):
    msum, out, stmp, n = 0, [], '', len(s)
    for i in range(n):
        if s[i] == ',' and msum == 0:  # 第一种情况结束
            out.append(stmp)
            stmp = ''
        elif s[i] == '"':
            if stmp == '':  # 第二种情况的开头
                stmp = '@'
                msum += 1
            else:
                if stmp[-1] == '#':  # 说明已经进入一个'"'符号了，现在让它变成真正的双引号
                    stmp = stmp[:-1] + '"'
                    msum -= 1
                else:  # 说明还没有出现过
                    stmp = stmp + '#'  # 记录半个
                    msum += 1
        elif s[i] == ',' and msum == 2 and stmp[-1] == '#':  # 第二种情况结束
            out.append(stmp[1:-1])
            stmp, msum = '', 0
        else:
            stmp = stmp + s[i]  # 进栈
    if stmp[0] == '@':  # 最后一个单词要特殊处理
        out.append(stmp[1:-1])
    else:
        out.append(stmp)
    return '\t'.join(out)


if __name__ == '__main__':
    slist = ['Jack,"写""段子""",58,1963-10-28',
             'Rommy,"读书,编程",63,1983-7-17',
             'Kate,购物,49.6,1979-12-56',
             'Jerry,"羽毛球,爬山",55.6,1980-5-26',
             'Jane,"下棋,""飞""",56.2,1976-8-23',
             'Ndldl,"""北京,""飞""",86.2,",""1976-,8-23,"""']
    for s in slist:
        print(segmentation(s))
*/

//func main() {
//	name := `'Jack,"写""段子""",58,1963-10-28',
//             'Rommy,"读书,编程",63,1983-7-17',
//             'Kate,购物,49.6,1979-12-56',
//             'Jerry,"羽毛球,爬山",55.6,1980-5-26',
//             'Jane,"下棋,""飞""",56.2,1976-8-23',
//             'Ndldl,"""北京,""飞""",86.2,",""1976-,8-23,"""'`
//
//	exec(name)
//}
//
//func exec(name string) {
//
//	i := len(name)
//
//	for a := range len(name){
//
//	}
//	fmt.Println(i)
//
//}
