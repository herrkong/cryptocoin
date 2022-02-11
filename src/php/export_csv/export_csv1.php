
<?php 
    $len = 100;
    $orderListAll = [];
    for ($i = 0; $i < ceil($list->select()->count() / $len); $i++) {
        $orderListAll = array_merge($orderListAll, $list->limit($i * $len, $len)->select()->toArray());
    }
    $csvHeaders = [
        '订单ID',
        '用户ID',
        '期权代码',
        '操作方向',
        '期权方向',
        '订单类型',
        '委托数量',
        '委托金额',
        '成交数量',
        '成交额',
        '实际手续费',
        '挂单时间',
        '最后修改时间',
        '撤销时间',
        '订单状态',
    ];
    $csvBody = [];
    foreach ($orderListAll as $key => $value) {
        if (explode('-',$value['instrument_id'])[3]=='C'){
            $upDown = "Call";
        }else{
            $upDown = "Put";
        }
        $csvBody [] = [
            'order_sys_id' => $value['order_sys_id'],
            'participant_id' => $value['participant_id'],
            'instrument_id' => $value['instrument_id'],
            'direction' => $value['direction'],
            'up_down' => $upDown,
            'order_price_type' => $value['order_price_type'],
            'volume_total_original' => $value['volume_total_original'],
            'limit_price' => $value['limit_price'],
            'volume_traded' => $value['volume_traded'],
            'turnover' => $value['turnover'],
            'fee_used' => $value['fee_used'],
            'insert_time' => $value['insert_time'],
            'update_time' => $value['update_time'],
            'cancel_time' => $value['cancel_time'],
            'order_status' => $value['order_status'],
        ];
    }
    //todo 导出csv方法待优化
    exportCsv(array_merge([$csvHeaders], $csvBody), '委托列表');
    unset($csvBody);
    exit();




    public function select($data = null)
    {
        if ($data instanceof Query) {
            return $data->select();
        } elseif ($data instanceof \Closure) {
            $data($this);
            $data = null;
        }

        $this->parseOptions();

        if (false === $data) {
            // 用于子查询 不查询只返回SQL
            $this->options['fetch_sql'] = true;
        } elseif (!is_null($data)) {
            // 主键条件分析
            $this->parsePkWhere($data);
        }

        $this->options['data'] = $data;

        $resultSet = $this->connection->select($this);

        if ($this->options['fetch_sql']) {
            return $resultSet;
        }

        // 返回结果处理
        if (!empty($this->options['fail']) && count($resultSet) == 0) {
            $this->throwNotFound($this->options);
        }

        // 数据列表读取后的处理
        if (!empty($this->model)) {
            // 生成模型对象
            $resultSet = $this->resultSetToModelCollection($resultSet);
        } else {
            $this->resultSet($resultSet);
        }

        return $resultSet;
    }

?>


<!-- // $csv = "";
// foreach ($data as $row) {
//     $csv .= join(",", $row) . "n";
// }
// echo $csv;


function toCSV(array $data, array $colHeaders = array(), $asString = false) {  
    $stream = ($asString)
        ? fopen("php://temp/maxmemory", "w+")
        : fopen("php://output", "w");  
 
    if (!empty($colHeaders)) {
        fputcsv($stream, $colHeaders);  
    }
 
    foreach ($data as $record) {
        fputcsv($stream, $record);
    }
 
    if ($asString) {
        rewind($stream);  
        $returnVal = stream_get_contents($stream);  
        fclose($stream);  
        return $returnVal;  
    }
    else {
        fclose($stream);  
    }  
}

toCSV() -->


<!-- // 导出CSV    一次性写入过多数据将导致数据库压力过大，分段取出
        if (isset($params['submit']) && $params['submit'] == 'export_csv') {
            //一次读取多少数据，可自由设定
            $len = 100;
            $orderListAll = [];
            for ($i = 0; $i < ceil($list->select()->count() / $len); $i++) {
                $orderListAll = array_merge($orderListAll, $list->limit($i * $len, $len)->select()->toArray());
            }
            $csvHeaders = [
                '订单ID',
                '用户ID',
                '期权代码',
                '操作方向',
                '期权方向',
                '订单类型',
                '委托数量',
                '委托金额',
                '成交数量',
                '成交额',
                '实际手续费',
                '挂单时间',
                '最后修改时间',
                '撤销时间',
                '订单状态',
            ];
            $csvBody = [];
            foreach ($orderListAll as $key => $value) {
                if (explode('-',$value['instrument_id'])[3]=='C'){
                    $upDown = "Call";
                }else{
                    $upDown = "Put";
                }
                $csvBody [] = [
                    'order_sys_id' => $value['order_sys_id'],
                    'participant_id' => $value['participant_id'],
                    'instrument_id' => $value['instrument_id'],
                    'direction' => $value['direction'],
                    'up_down' => $upDown,
                    'order_price_type' => $value['order_price_type'],
                    'volume_total_original' => $value['volume_total_original'],
                    'limit_price' => $value['limit_price'],
                    'volume_traded' => $value['volume_traded'],
                    'turnover' => $value['turnover'],
                    'fee_used' => $value['fee_used'],
                    'insert_time' => $value['insert_time'],
                    'update_time' => $value['update_time'],
                    'cancel_time' => $value['cancel_time'],
                    'order_status' => $value['order_status'],
                ];
            }
            //todo 导出csv方法待优化
            exportCsv(array_merge([$csvHeaders], $csvBody), '委托列表');
            unset($csvBody);
            exit();
        } -->