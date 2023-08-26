package vm

import "monkey/object"

// True 虚拟机上的布尔类型常量
var True = &object.Boolean{Value: true}

// False 虚拟机上的布尔类型常量
var False = &object.Boolean{Value: false}

// Null 无求值表达式 返回null
var Null = &object.Null{}
