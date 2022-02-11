sync_pool 

get 

put 


项目中大量重复地创建许多对象，造成 GC 的工作量巨大，CPU 频繁掉底。准备使用 sync.Pool 来缓存对象，减轻 GC 的消耗


https://www.cnblogs.com/qcrao-2018/p/12736031.html