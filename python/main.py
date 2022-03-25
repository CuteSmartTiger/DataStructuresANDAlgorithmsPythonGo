

def isValidParentheses(s: str) -> bool:
    str_map = {')': '(', '}': '{', ']': '['}
    stack = []
    for element in s:
        if element in str_map:
            if stack and stack[-1] == str_map[element]:
                stack.pop()
            else:
                return False
        else:
            stack.append(element)

    return stack == []


if __name__ == '__main__':
    s = '()'
    print(isValidParentheses(s))

