<?php
 
//  $data = [
//     [1, 2, 3],
//     [1, 2, 3],
//     [1, 2, 3],
// ];
function formatCsvData(){
    $exportOrderList = [
        [1, 2, 3],
        [1, 2, 3],
        [1, 2, 3],
    ];
    // $csvHeaders = [
    //     '订单ID',
    //     '用户ID',
    //     '期权代码',
    // ];
    $csvBody = [];
    foreach ($exportOrderList as $key => $value) {
        $csvBody [] = [
            'order_sys_id' => 1,
            'participant_id' => 2,
            'instrument_id' => 3,
        ];
    }
    //$csvData = array_merge([$csvHeaders], $csvBody);
  
    return $csvBody;
   
    
}

function dynamicExportCsv(){ 

    $data = formatCsvData();
    ExportCsv($data);
    

}

function ExportCsv($data){
     //让程序一直运行
     set_time_limit(0);
     //设置程序运行内存
     ini_set('memory_limit', '64M');
     $tableName = "委托列表";
     $fileName = $tableName.date('YmdHis');
     $csvHeaders = [
         '订单ID',
         '用户ID',
         '期权代码',
     ];
     header('Content-Type: application/download');
     header("Content-type:text/csv;");
     header("Content-Disposition:attachment;filename=".$fileName.".csv");
     header('Cache-Control:must-revalidate,post-check=0,pre-check=0');
     header('Expires:0');
     header('Pragma:public');


     $fp = fopen('php://output', 'a');

     foreach ($csvHeaders as $key => $value) {
         //CSV的Excel支持GBK编码，一定要转换，否则乱码
         $csvHeaders[$key] = iconv('utf-8', 'gbk', $value);
     }
 
     fputcsv($fp, $csvHeaders);
 
     $nums = 1;
 
     $dataCount = 3;
     $step = ceil($dataCount/$nums);//循环次数
 
     for ($i = 0; $i < $step; $i++) {
         //$result = $this->where($condition)->field($field)->order('order_sys_id desc')->offset($i * $nums)->limit($nums);
        $result = $data;
         foreach ($result as $item) {
             fputcsv($fp, $item);
         }
         //每1万条数据就刷新缓冲区
         ob_flush();
         flush();
     }
}


dynamicExportCsv();
