# xor-operation

---
tags: golang, 鐵人30天, leetcode
---
# golang 鐵人賽 自我挑戰賽 leetcode 30 天 第12天

## 題目解讀：

### 題目來源:
[xor-operation-in-an-array](https://leetcode.com/problems/xor-operation-in-an-array/)
### 原文:
Given an integer n and an integer start.

Define an array nums where nums[i] = start + 2*i (0-indexed) and n == nums.length.

Return the bitwise XOR of all elements of nums.
### 解讀：
給定一個正整數n,一個正整數 start

定義一個陣列nums, 長度為n

每個元素 nums[i]的值 = start + 2 * i

求出把每個nums[i]做XOR的結果


## 初步解法:
### 初步觀察:

首先是每個陣列數都是start+2 * i

因此只要循序把數值算出即可 不需要使用陣列儲存

另外可以把逐步把每個數值做XOR

### 初步設計:

Given an integer n, an integer start

step 0: let an integer result = 0, an integer idx = 0

step 1: if idx > n go to step 4

step 2: set result ^= start + 2 * idx

step 3: idx = idx + 1 go to step 3

step 4: return result

## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code

```golang
package xor_op

func xorOperation(n int, start int) int {
	result := 0
	idx := 0
	for idx < n {
		result ^= (start + 2*idx)
		idx++
	}
	return result
}

```
## 測資的撰寫
```golang
package xor_op

import "testing"

func Test_xorOperation(t *testing.T) {
	type args struct {
		n     int
		start int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				n:     5,
				start: 0,
			},
			want: 8,
		},
		{
			name: "Example2",
			args: args{
				n:     4,
				start: 3,
			},
			want: 8,
		},
		{
			name: "Example3",
			args: args{
				n:     1,
				start: 7,
			},
			want: 7,
		},
		{
			name: "Example4",
			args: args{
				n:     10,
				start: 5,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorOperation(tt.args.n, tt.args.start); got != tt.want {
				t.Errorf("xorOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

```
## my self record
[golang ironman 30day 12th day](https://hackmd.io/ept1MbTJSAOOPqPuRXQ7xg?view)

## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)