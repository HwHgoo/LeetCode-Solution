import string
import random
def gen_str(len: int) -> str:
    letters = string.ascii_lowercase
    s = ''.join(random.choice(letters) for _ in range(len))
    return s

def gen_bytes(len: int, exclude: list[str] = []) -> list[str] :
    if exclude.__len__() == 0 :
        exclude = [''] * (len+1)
    s = [random.choice(string.ascii_lowercase.replace(exclude[i], '')) for i in range(len)]
    return s

def gen_ints(len: int) -> list[int]:
    s = [random.randint(1, 1000000) for _ in range(len)]
    return s

str_len = 100
op_len = 199

print(gen_str(str_len))
print("\n-----------------------------------------------------\n")
print(gen_str(str_len))
print("\n-----------------------------------------------------\n")
origin = gen_bytes(op_len)
print(origin)
print("\n-----------------------------------------------------\n")
print(gen_bytes(op_len, origin))
print("\n-----------------------------------------------------\n")
print(gen_ints(op_len))