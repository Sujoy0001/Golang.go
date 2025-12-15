def p():
    for i in range(100):
        print(i)
    p()
p()

# This code will result in a RecursionError due to infinite recursion.

function_call_stack_limit = 1000  # Example limit for function call stack depth
def safe_p(depth=0):
    if depth < function_call_stack_limit:
        for i in range(100):
            print(i)
        safe_p(depth + 1)
    else:
        print("Reached maximum recursion depth.")
safe_p()