package stack

import "fmt"

type Stack struct{
    data []int
    top int
    max_size int
}

func main() {
    stack  := new(Stack)
    fmt.Println("Tell me the size of the stack")
    fmt.Scanf("%d", &stack.max_size)
    iniStack(stack)
    
    push(stack, 100)
    push(stack, 200)
    push(stack, 300)
    
    pop(stack)
    pop(stack)
    pop(stack)
}

func iniStack(s *Stack){
    s.top = -1
}

func isEmpty(s *Stack) bool{
    if (s.top == -1){
        return true
    }
    return false
}

func isFull(s *Stack, max_size int) bool{
    if (s.top == max_size-1){
        return true
    }
    return false
}

func pop(s *Stack) (int, error) {
    if (isEmpty(s)){
		err := fmt.Errorf("the stack is underflow")
		return -1, err
	}
    x := s.top
    s.top--
    return x, nil

}

func push(s *Stack, value int) error{
    if (isFull(s, s.max_size)){
		err := fmt.Errorf("the stack is underflow")
		return err
	}
    s.top++
    s.data[s.top] = value
    return nil
}
