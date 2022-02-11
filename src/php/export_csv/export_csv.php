<?php

$fileName = "拼团订单明细";
        $header = [
            '拼团主单号',
            '拼团子单号',
            '子订单号',
            '支付时间',
            '交易流水号',
            '商品SKU',
            '商品名称',
            '规格',
            '数量',
            '金额',
            '下单用户id',
            '下单用户昵称',
            '团长id',
            '团长昵称',
            '拼团订单状态',
            '销售订单状态',

        ];  //表头信息
        $fields = [
            'groupon_no',
            'groupon_child_no',
            'parcel_number',
            'pay_time',
            'pay_transaction_id',
            'sku_id',
            'good_name',
            'specification',
            'buy_goods_num',
            'pay_amt_e2',
            'buyer_id',
            'buyer_name',
            'groupon_author_id',
            'groupon_author_name',
            'buyer_status',
            'sell_order_status',
        ];


 /**
 * 导出excel(csv)
 * @data 导出数据
 * @headlist 第一行,列名
 * @fileName 输出Excel文件名
 */
function csv_export($data = array(), $headlist = array(), $fileName) {
  
    header('Content-Type: application/vnd.ms-excel');
    header('Content-Disposition: attachment;filename="'.$fileName.'.csv"');
    header('Cache-Control: max-age=0');
  
    //打开PHP文件句柄,php://output 表示直接输出到浏览器
    $fp = fopen('php://output', 'a');
    
    //输出Excel列名信息
    foreach ($headlist as $key => $value) {
        //CSV的Excel支持GBK编码，一定要转换，否则乱码
        $headlist[$key] = iconv('utf-8', 'gbk', $value);
    }
  
    //将数据通过fputcsv写到文件句柄
    fputcsv($fp, $headlist);
    
    //计数器
    $num = 0;
    
    //每隔$limit行，刷新一下输出buffer，不要太大，也不要太小
    $limit = 100000;
    
    //逐行取出数据，不浪费内存
    $count = count($data);
    for ($i = 0; $i < $count; $i++) {
    
        $num++;
        
        //刷新一下输出buffer，防止由于数据过多造成问题
        if ($limit == $num) { 
            ob_flush();
            flush();
            $num = 0;
        }
        
        $row = $data[$i];
        foreach ($row as $key => $value) {
            $row[$key] = iconv('utf-8', 'gbk', $value);
        }

        fputcsv($fp, $row);
    }
  }

  csv_export()

?>