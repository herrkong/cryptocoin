
#### insert sort 插入排序

d []int{5, 4, 3, 6, 2, 1}
 i = -1 j = 1
 i = 0
d[j=1] d[i=0]
453621


j = 2 i = 1
d[1] d[2]
435621


j=3 
j=4  i=2
432651

j=5 i =3 
432156

i = 4 




[0,4) 
4321

[4,5)
56
















a = 0
b = len - 1

i = a + 1 = 1
j = i = 1
j(1) > a(0) && d[j] < d[j-1]
4 5 3 6 2 1

i = 2 
j = 2 > 0 && d[2] < d[1]
4 3 5 6 2 1

j = 1 > 0 && d[1] < d[0]
3 4 5 6 2 1 

i = 3 
j = 3 d[3] > d[2]
j = 2 
j = 1

i = 4 
j = 4 && d[4] < d[3]
3 4 5 2 6 1
...
2 3 4 5 6 1

i = 5
j = 5 
...
1 2 3 4 5 6 


#### 插入排序的特点

挑待未排序元素 往之前排好的序列中找插入位置。

适合本来就基本有序的了，只有少量个元素无序的数组



#### quick sort 快速排序 pivot的选择  最坏O(n) 
平均 O(nlogn)

//sort.Sort()
使用快排  数组元素小于12个的时候 使用shell sort

实现Len() Less() Swap()
https://leetcode-cn.com/problems/sort-an-array/solution/pai-xu-shu-zu-by-leetcode-solution/


d []int{5, 4, 3, 6, 2, 1}

4 5  

doPivot(d,0,5)


 

前言
本题你可以选择直接调用库函数来对序列进行排序，但意义不大。由于排序算法有很多，本文只介绍三种常见的基于比较的复杂度较低的排序。

方法一：快速排序
思路和算法

快速排序的主要思想是通过划分将待排序的序列分成前后两部分，其中前一部分的数据都比后一部分的数据要小，然后再递归调用函数对两部分的序列分别进行快速排序，以此使整个序列达到有序。

我们定义函数 randomized_quicksort(nums, l, r) 为对 nums 数组里 [l,r][l,r] 的部分进行排序，每次先调用 randomized_partition 函数对 nums 数组里 [l,r][l,r] 的部分进行划分，并返回分界值的下标 pos，然后按上述将的递归调用 randomized_quicksort(nums, l, pos - 1) 和 randomized_quicksort(nums, pos + 1, r) 即可。

那么核心就是划分函数的实现了，划分函数一开始需要确定一个分界值（我们称之为主元 pivot)，然后再进行划分。而主元的选取有很多种方式，这里我们采用随机的方式，对当前划分区间 [l,r][l,r] 里的数等概率随机一个作为我们的主元，再将主元放到区间末尾，进行划分。

整个划分函数 partition 主要涉及两个指针 ii 和 jj，一开始 i = l - 1，j = l。我们需要实时维护两个指针使得任意时候，对于任意数组下标 kk，我们有如下条件成立：

l\leq k\leq il≤k≤i 时，\textit{nums}[k]\leq \textit{pivot}nums[k]≤pivot。

i+1\leq k\leq j-1i+1≤k≤j−1 时，\textit{nums}[k]> \textit{pivot}nums[k]>pivot。

k==rk==r 时，\textit{nums}[k]=\textit{pivot}nums[k]=pivot。

我们每次移动指针 jj ，如果 \textit{nums}[j]> pivotnums[j]>pivot，我们只需要继续移动指针 jj ，即能使上述三个条件成立，否则我们需要将指针 ii 加一，然后交换 \textit{nums}[i]nums[i] 和 \textit{nums}[j]nums[j]，再移动指针 jj 才能使得三个条件成立。

当 jj 移动到 r-1r−1 时结束循环，此时我们可以由上述三个条件知道 [l,i][l,i] 的数都小于等于主元 pivot，[i+1,r-1][i+1,r−1] 的数都大于主元 pivot，那么我们只要交换 \textit{nums}[i+1]nums[i+1] 和 \textit{nums}[r]nums[r] ，即能使得 [l,i+1][l,i+1] 区间的数都小于 [i+2,r][i+2,r] 区间的数，完成一次划分，且分界值下标为 i+1i+1，返回即可。

如下的动图展示了一次划分的过程，刚开始随机选了 44 作为主元，与末尾元素交换后开始划分：



![](2022-01-25-10-42-55.png)

复杂度分析

时间复杂度：基于随机选取主元的快速排序时间复杂度为期望 O(n\log n)O(nlogn)，其中 nn 为数组的长度。详细证明过程可以见《算法导论》第七章，这里不再大篇幅赘述。

空间复杂度：O(h)O(h)，其中 hh 为快速排序递归调用的层数。我们需要额外的 O(h)O(h) 的递归调用的栈空间，由于划分的结果不同导致了快速排序递归调用的层数也会不同，最坏情况下需 O(n)O(n) 的空间，最优情况下每次都平衡，此时整个递归树高度为 \log nlogn，空间复杂度为 O(\log n)O(logn)。




#### 稳定排序和不稳定排序




#### heap Sort 

构建大根堆

、


d []int {5, 4, 3, 6, 2, 1}

(6 -1) / 2 =  2

i = 2 hi = 6 fisrt = 0

root = 2 

child = 2 * 2 + 1 = 5 
root的两个子节点 2* root + 1  2 * root + 2
保证root > child 


data[first+root] = data[0+2] = data[2] = 3 
data[first + child] = data[0+5] = data[5] = 1

5 4 3 6 2 1 

root = 5 

child = 2 * 5 + 1 = 11 > hi break 


i = 1  
root = 1 child = 2 * 1 + 1 = 3 

3 + 1 < 6 data[0+3] data[4]

data[1] data[3]

5 6 3 4 2 1


i = 0 

6 5 3 4 2 1 


弹出 根节点

再调整堆 始终维持大根堆的特性














Top K，consumer消费带有不同优先级的任务队列时，用什么数据结构，就用这个


####  quicksort 和 mergeSort的对比

https://zhuanlan.zhihu.com/p/95080265


