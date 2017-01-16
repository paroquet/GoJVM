package heap
//非接口方法符号引用
func LookupMethodInClass(class *Class,name, descriptor string) *Method {
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

func lookupMehodInInterfaces(iface []*Class, name, descriptor string) *Method{
	for _, iface := range ifaces {
		for _, method := range iface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
		method := LookupMethodInInterfaces(iface.interfaces, name, descriptor)
		if method != nil{
			return method
		}
	}
	return nil 
}