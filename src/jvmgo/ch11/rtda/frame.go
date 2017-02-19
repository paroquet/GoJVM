package rtda

import "jvmgo/ch11/rtda/heap"

type Frame struct{
	lower			*Frame
	localVars		LocalVars
	operandStack	*OperandStack
	thread			*Thread
	method       	*heap.Method
	//下一条指令的地址
	nextPC			int
}

func newFrame(thread *Thread,method *heap.Method) *Frame{
	return &Frame{
		thread:			thread,
		method:			method,
		localVars:		newLocalVars(method.MaxLocals()),
		operandStack:	newOperandStack(method.MaxStack()),
	}
}
// getters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) Thread() *Thread{
	return self.thread
}
func (self *Frame) Method() *heap.Method {
	return self.method
}
func (self *Frame) NextPC() int{
	return self.nextPC
}
//setter
func (self *Frame) SetNextPC(nextPC int){
	self.nextPC = nextPC
}
//重置指令
func (self *Frame) RevertNextPC(){
	self.nextPC = self.thread.pc
}