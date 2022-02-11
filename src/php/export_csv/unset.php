<?php  
    //$test= str_repeat("1",256);  
    $test=str_repeat("1",255);
    $s = memory_get_usage();   
    //改函数用来查看当前所用内存  
    unset($test);  
    $e = memory_get_usage();  
    echo ' 释放内存: '.($s-$e);   
    //输出为272，但如果上面test变量改为
   // $test=str_repeat("1",255)，输出则为0  
?> 