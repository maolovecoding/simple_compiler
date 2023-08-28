# 借助go动手实现一个语言
- 基础语法
- 变量绑定
- 整型和布尔型
- 算术表达式
- 内置函数
- 头等函数和高阶函数
- 闭包
- 字符串数据结构
- 数组数据结构


## parse方法
普拉斯解析的工作方式
- 递归下降算法


## 使用方式
### 定义变量
```js
let name = "zs"
```
### 定义一个布尔
```js
let flag = true
```
### 定义一个数组
```js
let names = ['zs','ls']
```
### 定义一个数子
```js
let num = 10
```
### 定义一个hash
```js
let name = 'ls'
let map = {
    "zs": "ls",
    name: 'zs'
}
```
### 定义函数
```txt
let getName = fn() {
    return "zs";
}
```
### 内置函数调用
```js
let names = ['zs','ls'];
first(names)
```