def fizzBuzz(self, n):
    return['FizzBuzz'[i%-3&-4:i%-5&8^12]or`i`for i in range(1,n+1)]